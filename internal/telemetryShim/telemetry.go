/*
Copyright (C) 2021-2023, Kubefirst

This program is licensed under MIT.
See the LICENSE file for more details.
*/
package telemetryShim

import (
	"os"
	"time"

	"github.com/denisbrodbeck/machineid"
	pkgtypes "github.com/kubefirst/kubefirst-api/pkg/types"
	"github.com/kubefirst/runtime/pkg"
	"github.com/kubefirst/runtime/pkg/segment"
	"github.com/segmentio/analytics-go"
	log "github.com/sirupsen/logrus"
)

// Heartbeat
func Heartbeat(segmentClient *segment.SegmentClient) {
	// sent one heartbeat for the mgmt cluster
	Transmit(segmentClient, segment.MetricKubefirstHeartbeat, "")
	// workload
	HeartbeatWorkloadClusters()
	//TODO! DIETZ - NO WAY
	for range time.Tick(time.Minute * 2) {
		// sent one heartbeat for the mgmt cluster
		Transmit(segmentClient, segment.MetricKubefirstHeartbeat, "")
		// workload
		HeartbeatWorkloadClusters()
	}
}

// SetupTelemetry
func SetupTelemetry(cl pkgtypes.Cluster) (*segment.SegmentClient, error) {
	kubefirstVersion := os.Getenv("KUBEFIRST_VERSION")
	if kubefirstVersion == "" {
		kubefirstVersion = "development"
	}

	strippedDomainName, err := pkg.RemoveSubdomainV2(cl.DomainName)
	if err != nil {
		return &segment.SegmentClient{}, nil
	}
	machineID, _ := machineid.ID()

	// Segment Client
	segmentClient := &segment.SegmentClient{
		Client:            analytics.New(segment.SegmentIOWriteKey),
		CliVersion:        kubefirstVersion,
		CloudProvider:     cl.CloudProvider,
		ClusterID:         cl.ClusterID,
		ClusterType:       cl.ClusterType,
		DomainName:        strippedDomainName,
		GitProvider:       cl.GitProvider,
		KubefirstClient:   "api",
		KubefirstTeam:     cl.KubefirstTeam,
		KubefirstTeamInfo: os.Getenv("KUBEFIRST_TEAM_INFO"),
		MachineID:         machineID,
	}
	segmentClient.SetupClient()

	return segmentClient, nil
}

// Transmit sends a metric via Segment
func Transmit(segmentClient *segment.SegmentClient, metricName string, errorMessage string) {
	segmentMsg := segmentClient.SendCountMetric(metricName, errorMessage)
	if segmentMsg != "" {
		log.Info(segmentMsg)
	}
}
