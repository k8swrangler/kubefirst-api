// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Kubefirst",
            "email": "help@kubefirst.io"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/aws/domain/validate/:domain": {
            "get": {
                "description": "Returns status of whether or not an AWS hosted zone is validated for use with Kubefirst",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aws"
                ],
                "summary": "Returns status of whether or not an AWS hosted zone is validated for use with Kubefirst",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Domain name, no trailing dot",
                        "name": "domain",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.AWSDomainValidateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.JSONFailureResponse"
                        }
                    }
                }
            }
        },
        "/aws/profiles": {
            "get": {
                "description": "Returns a list of configured AWS profiles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aws"
                ],
                "summary": "Returns a list of configured AWS profiles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.AWSProfilesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.JSONFailureResponse"
                        }
                    }
                }
            }
        },
        "/civo/domain/validate/:domain": {
            "get": {
                "description": "Returns status of whether or not a Civo hosted zone is validated for use with Kubefirst",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "civo"
                ],
                "summary": "Returns status of whether or not a Civo hosted zone is validated for use with Kubefirst",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Domain name, no trailing dot",
                        "name": "domain",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Domain validation request in JSON format",
                        "name": "settings",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.CivoDomainValidationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.CivoDomainValidationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.JSONFailureResponse"
                        }
                    }
                }
            }
        },
        "/cluster": {
            "get": {
                "description": "Return all known configured Kubefirst clusters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cluster"
                ],
                "summary": "Return all known configured Kubefirst clusters",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.Cluster"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.JSONFailureResponse"
                        }
                    }
                }
            }
        },
        "/cluster/:cluster_name": {
            "get": {
                "description": "Return a configured Kubefirst cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cluster"
                ],
                "summary": "Return a configured Kubefirst cluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cluster name",
                        "name": "cluster_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Cluster"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.JSONFailureResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a Kubefirst cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cluster"
                ],
                "summary": "Create a Kubefirst cluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cluster name",
                        "name": "cluster_name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Cluster create request in JSON format",
                        "name": "definition",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ClusterDefinition"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/types.JSONSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.JSONFailureResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Kubefirst cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cluster"
                ],
                "summary": "Delete a Kubefirst cluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cluster name",
                        "name": "cluster_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/types.JSONSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.JSONFailureResponse"
                        }
                    }
                }
            }
        },
        "/cluster/:cluster_name/export": {
            "post": {
                "description": "Export a Kubefirst cluster database entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cluster"
                ],
                "summary": "Export a Kubefirst cluster database entry",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cluster name",
                        "name": "cluster_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/types.JSONSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.JSONFailureResponse"
                        }
                    }
                }
            }
        },
        "/cluster/:cluster_name/reset_progress": {
            "post": {
                "description": "Remove a cluster progress marker from a cluster entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cluster"
                ],
                "summary": "Remove a cluster progress marker from a cluster entry",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Cluster name",
                        "name": "cluster_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/types.JSONSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.JSONFailureResponse"
                        }
                    }
                }
            }
        },
        "/cluster/import": {
            "post": {
                "description": "Import a Kubefirst cluster database entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cluster"
                ],
                "summary": "Import a Kubefirst cluster database entry",
                "parameters": [
                    {
                        "description": "Cluster import request in JSON format",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/types.ImportClusterRequest"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/types.JSONSuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.JSONFailureResponse"
                        }
                    }
                }
            }
        },
        "/domain/:cloud_provider": {
            "get": {
                "description": "Return a configured Kubefirst cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "domain"
                ],
                "summary": "Return a configured Kubefirst cluster",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The cloud provider to return registered domains/zones from",
                        "name": "cloud_provider",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.Cluster"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.JSONFailureResponse"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Return health status if the application is running.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Return health status if the application is running.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.JSONHealthResponse"
                        }
                    }
                }
            }
        },
        "/stream": {
            "get": {
                "description": "Stream API server logs",
                "tags": [
                    "logs"
                ],
                "summary": "Stream API server logs",
                "responses": {}
            }
        }
    },
    "definitions": {
        "types.AWSAuth": {
            "type": "object",
            "properties": {
                "access_key_id": {
                    "type": "string"
                },
                "secret_access_key": {
                    "type": "string"
                },
                "session_token": {
                    "type": "string"
                }
            }
        },
        "types.AWSDomainValidateResponse": {
            "type": "object",
            "properties": {
                "validated": {
                    "type": "boolean"
                }
            }
        },
        "types.AWSProfilesResponse": {
            "type": "object",
            "properties": {
                "profiles": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "types.CivoAuth": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "types.CivoDomainValidationRequest": {
            "type": "object",
            "properties": {
                "cloud_region": {
                    "type": "string"
                }
            }
        },
        "types.CivoDomainValidationResponse": {
            "type": "object",
            "properties": {
                "validated": {
                    "type": "boolean"
                }
            }
        },
        "types.Cluster": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string"
                },
                "alerts_email": {
                    "type": "string"
                },
                "argocd_auth_token": {
                    "type": "string"
                },
                "argocd_create_registry_check": {
                    "type": "boolean"
                },
                "argocd_delete_registry_check": {
                    "type": "boolean"
                },
                "argocd_initialize_check": {
                    "type": "boolean"
                },
                "argocd_install_check": {
                    "type": "boolean"
                },
                "argocd_password": {
                    "type": "string"
                },
                "argocd_username": {
                    "type": "string"
                },
                "atlantis_webhook_secret": {
                    "type": "string"
                },
                "atlantis_webhook_url": {
                    "type": "string"
                },
                "aws_account_id": {
                    "description": "kms",
                    "type": "string"
                },
                "aws_auth": {
                    "description": "Auth",
                    "allOf": [
                        {
                            "$ref": "#/definitions/types.AWSAuth"
                        }
                    ]
                },
                "aws_kms_key_detokenized_check": {
                    "type": "boolean"
                },
                "aws_kms_key_id": {
                    "type": "string"
                },
                "civo_auth": {
                    "$ref": "#/definitions/types.CivoAuth"
                },
                "cloud_provider": {
                    "type": "string"
                },
                "cloud_region": {
                    "type": "string"
                },
                "cloud_terraform_apply_check": {
                    "type": "boolean"
                },
                "cloud_terraform_apply_failed_check": {
                    "type": "boolean"
                },
                "cluster_id": {
                    "type": "string"
                },
                "cluster_name": {
                    "type": "string"
                },
                "cluster_secrets_created_check": {
                    "type": "boolean"
                },
                "cluster_type": {
                    "type": "string"
                },
                "creation_timestamp": {
                    "type": "string"
                },
                "do_auth": {
                    "$ref": "#/definitions/types.DigitaloceanAuth"
                },
                "domain_liveness_check": {
                    "type": "boolean"
                },
                "domain_name": {
                    "type": "string"
                },
                "git_host": {
                    "type": "string"
                },
                "git_init_check": {
                    "type": "boolean"
                },
                "git_owner": {
                    "type": "string"
                },
                "git_provider": {
                    "type": "string"
                },
                "git_terraform_apply_check": {
                    "type": "boolean"
                },
                "git_token": {
                    "type": "string"
                },
                "git_user": {
                    "type": "string"
                },
                "gitlab_owner_group_id": {
                    "type": "integer"
                },
                "gitops_pushed_check": {
                    "type": "boolean"
                },
                "gitops_ready_check": {
                    "type": "boolean"
                },
                "in_progress": {
                    "type": "boolean"
                },
                "install_tools_check": {
                    "description": "Checks",
                    "type": "boolean"
                },
                "kbot_setup_check": {
                    "type": "boolean"
                },
                "kubefirst_team": {
                    "type": "string"
                },
                "last_condition": {
                    "type": "string"
                },
                "private_key": {
                    "type": "string"
                },
                "public_key": {
                    "type": "string"
                },
                "public_keys": {
                    "type": "string"
                },
                "state_store_create_check": {
                    "type": "boolean"
                },
                "state_store_credentials": {
                    "$ref": "#/definitions/types.StateStoreCredentials"
                },
                "state_store_creds_check": {
                    "type": "boolean"
                },
                "state_store_details": {
                    "$ref": "#/definitions/types.StateStoreDetails"
                },
                "status": {
                    "type": "string"
                },
                "useTelemetry": {
                    "description": "Telemetry",
                    "type": "boolean"
                },
                "users_terraform_apply_check": {
                    "type": "boolean"
                },
                "vault_initialized_check": {
                    "type": "boolean"
                },
                "vault_terraform_apply_check": {
                    "type": "boolean"
                },
                "vultr_auth": {
                    "$ref": "#/definitions/types.VultrAuth"
                }
            }
        },
        "types.ClusterDefinition": {
            "type": "object",
            "required": [
                "admin_email",
                "cloud_provider",
                "cloud_region",
                "domain_name",
                "git_owner",
                "git_provider",
                "git_token",
                "type"
            ],
            "properties": {
                "admin_email": {
                    "type": "string"
                },
                "aws_auth": {
                    "$ref": "#/definitions/types.AWSAuth"
                },
                "civo_auth": {
                    "$ref": "#/definitions/types.CivoAuth"
                },
                "cloud_provider": {
                    "type": "string",
                    "enum": [
                        "aws",
                        "civo",
                        "digitalocean",
                        "vultr"
                    ]
                },
                "cloud_region": {
                    "type": "string"
                },
                "cluster_name": {
                    "type": "string"
                },
                "do_auth": {
                    "$ref": "#/definitions/types.DigitaloceanAuth"
                },
                "domain_name": {
                    "type": "string"
                },
                "git_owner": {
                    "type": "string"
                },
                "git_provider": {
                    "type": "string",
                    "enum": [
                        "github",
                        "gitlab"
                    ]
                },
                "git_token": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "mgmt",
                        "workload"
                    ]
                },
                "vultr_auth": {
                    "$ref": "#/definitions/types.VultrAuth"
                }
            }
        },
        "types.DigitaloceanAuth": {
            "type": "object",
            "properties": {
                "spaces_key": {
                    "type": "string"
                },
                "spaces_secret": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "types.ImportClusterRequest": {
            "type": "object",
            "properties": {
                "cloud_provider": {
                    "type": "string"
                },
                "cloud_region": {
                    "type": "string"
                },
                "cluster_name": {
                    "type": "string"
                },
                "state_store_credentials": {
                    "$ref": "#/definitions/types.StateStoreCredentials"
                },
                "state_store_details": {
                    "$ref": "#/definitions/types.StateStoreDetails"
                }
            }
        },
        "types.JSONFailureResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "err"
                }
            }
        },
        "types.JSONHealthResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "healthy"
                }
            }
        },
        "types.JSONSuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "types.StateStoreCredentials": {
            "type": "object",
            "properties": {
                "access_key_id": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "secret_access_key": {
                    "type": "string"
                }
            }
        },
        "types.StateStoreDetails": {
            "type": "object",
            "properties": {
                "aws_artifacts_bucket": {
                    "type": "string"
                },
                "aws_state_store_bucket": {
                    "type": "string"
                },
                "hostname": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "types.VultrAuth": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:port",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Kubefirst API",
	Description:      "Kubefirst API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
