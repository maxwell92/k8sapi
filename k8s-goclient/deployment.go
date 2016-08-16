package main

import (
	"encoding/json"
	"fmt"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"log"
	"os"
	"time"
)

var logger *log.Logger

var (
	str = `
	{
    "kind": "Deployment",
    "apiVersion": "extensions/v1beta1",
    "metadata": {
        "name": "nginx-deploy-k8sclient",
        "namespace": "default",
        "labels": {
            "app": "nginx",
           	"version": "1.7.9"
        },
        "annotations": {
            "deployment.kubernetes.io/revision": "1353"
        }
    },
    "spec": {
        "replicas": 3,
        "selector": {
            "matchLabels": {
                "app": "nginx",
				"version": "1.7.9"
            }
        },
        "template": {
            "metadata": {
                "creationTimestamp": null,
                "labels": {
                    "app": "nginx",
					"version": "1.7.9"
                }
            },
            "spec": {
                "containers": [
                    {
                        "name": "nginx",
                        "image": "registry.test.com:5000/nginx:1.7.9",
                        "ports": [
                            {
                                "containerPort": 80,
                                "protocol": "TCP"
                            }
                        ],
                        "resources": {},
                        "terminationMessagePath": "/dev/termination-log",
                        "imagePullPolicy": "IfNotPresent"
                    }
                ],
                "restartPolicy": "Always",
                "terminationGracePeriodSeconds": 30,
                "dnsPolicy": "ClusterFirst",
                "securityContext": {}
            }
        },
        "strategy": {
            "type": "RollingUpdate",
            "rollingUpdate": {
                "maxUnavailable": 1,
                "maxSurge": 1
            }
        }
    }
}
`
)

func init() {
	logger = log.New(os.Stdout, "", 0)
}

const (
	SERVER string = "http://172.21.1.11:8080"
)

func main() {
	fmt.Println(SERVER)
	config := &restclient.Config{
		Host: SERVER,
	}

	c, err := client.New(config)

	if err != nil {
		logger.Fatalf("Could not connect to k8s api: err=%s\n", err)
	}

	deployment := new(extensions.Deployment)

	err = json.Unmarshal([]byte(str), deployment)

	if err != nil {
		logger.Fatalf("Could not unmarshal from str: err=%s\n", err)
	}

	// Create deployment
	_, err = c.Extensions().Deployments(api.NamespaceDefault).Create(deployment)

	if err != nil {
		logger.Fatalf("Could not create deployment: err=%s\n", err)
	}

	logger.Printf("Deployment was created!")

	time.Sleep(time.Duration(5) * time.Second)
	err = c.Extensions().Deployments(api.NamespaceDefault).Delete("nginx-deploy-k8sclient", &api.DeleteOptions{})
	if err != nil {
		logger.Fatalf("Could not delete deployment: %s, err= %s\n", "nginx-deploy-k8sclient", err)
	}
	logger.Printf("Deployment was deleted!")
}
