/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package civo

import (
	"context"
	"os"
	"strings"

	"github.com/kubefirst/kubefirst-api/internal/constants"
	"github.com/kubefirst/kubefirst-api/internal/controller"
	"github.com/kubefirst/kubefirst-api/internal/db"
	"github.com/kubefirst/kubefirst-api/internal/services"
	"github.com/kubefirst/kubefirst-api/internal/telemetryShim"
	"github.com/kubefirst/kubefirst-api/internal/types"
	"github.com/kubefirst/runtime/pkg/civo"
	"github.com/kubefirst/runtime/pkg/k8s"
	"github.com/kubefirst/runtime/pkg/segment"
	"github.com/kubefirst/runtime/pkg/ssl"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateCivoCluster(definition *types.ClusterDefinition) error {
	ctrl := controller.ClusterController{}
	err := ctrl.InitController(definition)
	if err != nil {
		return err
	}

	err = ctrl.MdbCl.UpdateCluster(ctrl.ClusterName, "in_progress", true)
	if err != nil {
		return err
	}

	err = ctrl.DownloadTools(ctrl.ProviderConfig.(*civo.CivoConfig).ToolsDir)
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.DomainLivenessTest()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.StateStoreCredentials()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.StateStoreCreate()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.GitInit()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.InitializeBot()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.RepositoryPrep()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.RunGitTerraform()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.RepositoryPush()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.CreateCluster()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	// Needs wait after cluster create

	err = ctrl.ClusterSecretsBootstrap()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	//* check for ssl restore
	log.Info("checking for tls secrets to restore")
	secretsFilesToRestore, err := os.ReadDir(ctrl.ProviderConfig.(*civo.CivoConfig).SSLBackupDir + "/secrets")
	if err != nil {
		log.Infof("%s", err)
	}
	if len(secretsFilesToRestore) != 0 {
		// todo would like these but requires CRD's and is not currently supported
		// add crds ( use execShellReturnErrors? )
		// https://raw.githubusercontent.com/cert-manager/cert-manager/v1.11.0/deploy/crds/crd-clusterissuers.yaml
		// https://raw.githubusercontent.com/cert-manager/cert-manager/v1.11.0/deploy/crds/crd-certificates.yaml
		// add certificates, and clusterissuers
		log.Infof("found %d tls secrets to restore", len(secretsFilesToRestore))
		ssl.Restore(ctrl.ProviderConfig.(*civo.CivoConfig).SSLBackupDir, ctrl.DomainName, ctrl.ProviderConfig.(*civo.CivoConfig).Kubeconfig)
	} else {
		log.Info("no files found in secrets directory, continuing")
	}

	err = ctrl.InstallArgoCD()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.InitializeArgoCD()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.DeployRegistryApplication()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.WaitForVault()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.InitializeVault()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	// Create kubeconfig client
	kcfg := k8s.CreateKubeConfig(false, ctrl.ProviderConfig.(*civo.CivoConfig).Kubeconfig)

	// Create Secret for initial cluster import
	cl, _ := db.Client.GetCluster(ctrl.ClusterName)

	_, err = kcfg.Clientset.CoreV1().Secrets(constants.KubefirstNamespace).Get(context.Background(), constants.KubefirstImportSecretName, metav1.GetOptions{})
	if err == nil {
		log.Infof("kubernetes secret %s/%s already created - skipping", constants.KubefirstNamespace, constants.KubefirstImportSecretName)
	} else if strings.Contains(err.Error(), "not found") {
		_, err = kcfg.Clientset.CoreV1().Secrets(constants.KubefirstNamespace).Create(context.Background(), &v1.Secret{
			Type: "Opaque",
			ObjectMeta: metav1.ObjectMeta{
				Name:      constants.KubefirstImportSecretName,
				Namespace: constants.KubefirstNamespace,
			},
			Data: map[string][]byte{
				"CLOUD_PROVIDER":    []byte(cl.CloudProvider),
				"CLOUD_REGION":      []byte(cl.CloudRegion),
				"CLUSTER_NAME":      []byte(cl.ClusterName),
				"ACCESS_KEY_ID":     []byte(cl.StateStoreCredentials.AccessKeyID),
				"SECRET_ACCESS_KEY": []byte(cl.StateStoreCredentials.SecretAccessKey),
				"STATE_STORE_NAME":  []byte(cl.StateStoreDetails.Name),
			},
		}, metav1.CreateOptions{})
		if err != nil {
			log.Fatalf("error creating kubernetes secret for initial import: %s", err)
		}
		log.Info("Created Secret for initial cluster import")
	}

	// SetupMinioStorage(kcfg, ctrl.ProviderConfig.K1Dir, ctrl.GitProvider)

	//* configure vault with terraform
	//* vault port-forward
	vaultStopChannel := make(chan struct{}, 1)
	defer func() {
		close(vaultStopChannel)
	}()
	k8s.OpenPortForwardPodWrapper(
		kcfg.Clientset,
		kcfg.RestConfig,
		"vault-0",
		"vault",
		8200,
		8200,
		vaultStopChannel,
	)

	err = ctrl.RunVaultTerraform()
	if err != nil {
		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.RunUsersTerraform()
	if err != nil {
		return err
	}

	// Wait for console Deployment Pods to transition to Running
	log.Info("deploying kubefirst console and verifying cluster installation is complete")
	consoleDeployment, err := k8s.ReturnDeploymentObject(
		kcfg.Clientset,
		"app.kubernetes.io/instance",
		"kubefirst-console",
		"kubefirst",
		1200,
	)
	if err != nil {
		log.Errorf("Error finding console Deployment: %s", err)
		ctrl.HandleError(err.Error())
		return err
	}
	_, err = k8s.WaitForDeploymentReady(kcfg.Clientset, consoleDeployment, 120)
	if err != nil {
		log.Errorf("Error waiting for console Deployment ready state: %s", err)

		ctrl.HandleError(err.Error())
		return err
	}

	err = ctrl.MdbCl.UpdateCluster(ctrl.ClusterName, "status", "provisioned")
	if err != nil {
		return err
	}

	err = ctrl.MdbCl.UpdateCluster(ctrl.ClusterName, "in_progress", false)
	if err != nil {
		return err
	}

	log.Info("cluster creation complete")

	// Telemetry handler
	rec, err := ctrl.GetCurrentClusterRecord()
	if err != nil {
		return err
	}

	segmentClient, err := telemetryShim.SetupTelemetry(rec)
	if err != nil {
		return err
	}
	defer segmentClient.Client.Close()

	telemetryShim.Transmit(rec.UseTelemetry, segmentClient, segment.MetricClusterInstallCompleted, "")

	// Create default service entries
	cl, _ = db.Client.GetCluster(ctrl.ClusterName)
	err = services.AddDefaultServices(&cl)
	if err != nil {
		log.Errorf("error adding default service entries for cluster %s: %s", cl.ClusterName, err)
	}

	return nil
}
