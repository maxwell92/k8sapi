package main

import (
	"os"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"log"
	// "encoding/json"
)

var logger *log.Logger

var (
	str = `{
        "metadata": {
                "name": "redis-master-53yf7",
                "generateName": "redis-master-",
                "namespace": "default",
                "selfLink": "/api/v1/namespaces/default/pods/redis-master-53yf7",
                "uid": "8e5020b6-5d31-11e6-b9d6-44a84240716a",
                "resourceVersion": "3858362",
                "creationTimestamp": "2016-08-08T06:30:10Z",
                "labels": {
                        "name": "redis-master"
                },
                "annotations": {
                        "kubernetes.io/created-by": "{\"kind\":\"SerializedReference\",\"apiVersion\":\"v1\",\"reference\":{\"kind\":\"ReplicationController\",\"namespace\":\"default\",\"name\":\"redis-master\",\"uid\":\"8dbbda83-5d31-11e6-b9d6-44a84240716a\",\"apiVersion\":\"v1\",\"resourceVersion\":\"3858302\"}}\n"
                }
        },
        "spec": {
                "volumes": [
                        {
                                "name": "default-token-7fk3j",
                                "secret": {
                                        "secretName": "default-token-7fk3j"
                                }
                        }
                ],
                "containers": [
                        {
                                "name": "master",
                                "image": "registry.test.com:5000/redis",
                                "ports": [
                                        {
                                                "containerPort": 6379,
                                                "protocol": "TCP"
                                        }
                                ],
                                "resources": {},
                                "volumeMounts": [
                                        {
                                                "name": "default-token-7fk3j",
                                                "readOnly": true,
                                                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
                                        }
                                ],
                                "terminationMessagePath": "/dev/termination-log",
                                "imagePullPolicy": "Always"
                        }
                ],
                "restartPolicy": "Always",
                "terminationGracePeriodSeconds": 30,
                "dnsPolicy": "ClusterFirst",
                "serviceAccountName": "default",
                "nodeName": "172.21.1.21",
                "securityContext": {}
        },
        "status": {
                "phase": "Running",
                "conditions": [
                        {
                                "type": "Ready",
                                "status": "True",
                                "lastProbeTime": null,
                                "lastTransitionTime": "2016-08-08T06:30:27Z"
                        }
                ],
                "hostIP": "172.21.1.21",
                "podIP": "10.0.62.2",
                "startTime": "2016-08-08T06:30:10Z",
                "containerStatuses": [
                        {
                                "name": "master",
                                "state": {
                                        "running": {
                                                "startedAt": "2016-08-08T06:30:26Z"
                                        }
                                },
                                "lastState": {},
                                "ready": true,
                                "restartCount": 0,
                                "image": "registry.test.com:5000/redis",
                                "imageID": "docker://sha256:bab6d3ebc78ce403eac7196cfa7f7625eebbc6341ddd3a807a7f71fd5d94d3df",
                                "containerID": "docker://65e05842cb20dd1d7215d23d5c8358258ca39fac8c1b605f2ec090efad7b0235"
                        }
                ]
        }
}`
)

func init() {
	logger = log.New(os.Stdout, "", 0)
}

const (
	SERVER string = "http://172.21.1.11:8080"
)

func main() {

	config := &restclient.Config{
		Host:        SERVER,
		BearerToken: "",
	}

	c, err := client.New(config)

	if err != nil {
		logger.Fatalf("Could not connect to k8s api: %s\n", err)
	}

	/*
	// List all pods
	list, _ := c.Pods(api.NamespaceDefault).List(api.ListOptions{})

	items := list.Items

	// Foreach every pod
	for _, pod := range items {

		name := pod.ObjectMeta.Name
		logger.Printf("Pod: name=%s\n", name)

		// Get pod named "redis-master-53yf7"
		pod, err := c.Pods(api.NamespaceDefault).Get(name)

		if err != nil {
			logger.Fatalf("Could not connect to k8s api: %s\n", err)
			return
		}

		logger.Printf("\tPod detail info: NodeName=%s, DNSPolicy=%s, RestartPolicy=%s\n", pod.Spec.NodeName, pod.Spec.DNSPolicy, pod.Spec.RestartPolicy)
	}
	*/

	/*
	// Encode to json, or Decode from json-string
	pod, err := c.Pods(api.NamespaceDefault).Get("redis-master-53yf7")
	data, _ := json.MarshalIndent(pod, "", "\t")
	logger.Printf("%s\n", string(data))

	newPod := new(api.Pod)
	json.Unmarshal([]byte(str), newPod)

	logger.Printf("Name: %s\n", newPod.ObjectMeta.Name)
	*/

	// Test Delete
	name := "redis-master-53yf7"
	/*
	pod, err := c.Pods(api.NamespaceDefault).Get(name)
	if err != nil {
		logger.Fatalf("Could not get k8s pod: podName=%s, err=%s\n", name, err)
	}
	*/

	err =  c.Pods(api.NamespaceDefault).Delete(name, &api.DeleteOptions{})

	if err != nil {
		logger.Fatalf("Could not delete k8s pod: podName=%s, err=%s\n", name, err)
	}

	logger.Printf("Pod was deleted: name=%s", name)

}
