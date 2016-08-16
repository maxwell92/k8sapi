package main

import (
	//"encoding/json"
	"fmt"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"log"
	"os"
)

var logger *log.Logger

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

	replicasetlist := new(extensions.ReplicaSetList)

	// Get replicasetlist
	replicasetlist, err = c.Extensions().ReplicaSets(api.NamespaceDefault).List(api.ListOptions{})
	if err != nil {
		logger.Fatalf("Could not get replicasetlist: err=%s\n", err)
	}

	fmt.Println("Getting replicaset list")
	count := len(replicasetlist.Items)
	fmt.Println(count)
	for i := 0; i < count; i++ {
		fmt.Printf("[%d]: %s\n", i, replicasetlist.Items[i].ObjectMeta.Name)
	}
}
