package app

import (
	"fmt"
	"github.com/kataras/iris"
	hc "httpclient"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
	"log"
	"model/app"
	"strings"
)

const (
	SERVER string = "http://172.21.1.11:8080"
)

var instance *AppController

type AppController struct {
	cli *client.Client
}

func NewAppController(server string) *AppController {
	config := &restclient.Config{
		Host: server,
	}
	cli, err := client.New(config)
	if err != nil {
		log.Fatalf("New AppController err=%s\n", err)
	}

	instance = &AppController{cli: cli}
	return instance
}

func (ac *AppController) List(ctx *iris.Context) {
	oid := ctx.Param("oid")
	//List Deployment
	//deploymentlist, err := ac.cli.Extensions().Deployments(oid).List(api.ListOptions{})

	//List ReplicaSet
	//replicasetlist, err := ac.cli.Extensions().ReplicaSets(oid).List(api.ListOptions{})

	podlist, err := ac.cli.Pods(oid).List(api.ListOptions{})
	if err != nil {
		log.Fatalf("list %s err= %s\n", oid, err)
	}

	//count := len(deploymentlist.Items)
	//count := len(replicasetlist.Items)
	count := len(podlist.Items)
	for i := 0; i < count; i++ {
		//fmt.Printf("[%d]: %s\n", i, deploymentlist.Items[i].ObjectMeta.Name)
		//fmt.Printf("[%d]: %s %d\n", i, replicasetlist.Items[i].ObjectMeta.Name, replicasetlist.Items[i].Spec.Replicas)
		fmt.Printf("[%d]: %s\n", i, podlist.Items[i].ObjectMeta.Name)
	}
}

func (ac *AppController) Create(ctx *iris.Context) {
	//oid := ctx.Param("oid")

	//oid := mysqlClient.Select(user).OId
	oid := api.NamespaceDefault

	App := new(app.AppDeployment)

	//myapp := ctx.PostBody()
	//err := json.Unmarshal([]byte(myapp), App)
	err := ctx.ReadJSON(App)
	if err != nil {
		log.Fatalf("create %s err=%s\n", oid, err)
	}
	deployment := new(extensions.Deployment)
	deployment = &App.Deployment

	dclen := len(App.Datacenters)
	dc := make([]app.AppDc, dclen)
	dc = App.Datacenters

	var Server string
	for _, v := range dc {
		switch v.DcID {
		case 1:
			Server = "http://172.21.1.11:8080"
		case 2:
			Server = "http://172.21.1.11:8080"
		case 3:
			Server = "http://172.21.1.11:8080"
		}

		newconfig := &restclient.Config{
			Host: Server,
		}
		newCli, err := client.New(newconfig)
		if err != nil {
			log.Fatalf("New client err = %s\n", err)
		}

		_, err = newCli.Extensions().Deployments(oid).Create(deployment)
		if err != nil {
			log.Fatalf("craete deployment in %s err = %s\n", oid, err)
		}

		fmt.Printf("create deployment in %d\n", v.DcID)
	}
}

func (ac *AppController) Test(ctx *iris.Context) {
	client := hc.NewHttpClient("", "")
	rep, err := client.Post("http://localhost:10000/api/v1/deployments", strings.NewReader(string(myApp)))
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(rep))
	}
}

func (ac *AppController) Describe(ctx *iris.Context) {
	oid := ctx.Param("oid")
	id := ctx.Param("id")

	//deployment, err := ac.cli.Extensions().Deployments(oid).Get(id)
	//replicaset, err := ac.cli.Extensions().ReplicaSets(oid).Get(id)
	pod, err := ac.cli.Pods(oid).Get(id)
	if err != nil {
		log.Fatalf("describe %s - %s err=%s\n", oid, id, err)
	}

	//fmt.Printf("%s: %d\n", replicaset.ObjectMeta.Name, replicaset.Spec.Replicas)
	fmt.Printf("%s: \n", pod.ObjectMeta.Name)
}

func (ac *AppController) EncodeJson() {

}

func (ac *AppController) DecodeJson() {

}

func (ac *AppController) Delete(c *client.Client) error {
	err := c.Extensions().Deployments(api.NamespaceDefault).Delete("nginx-deploy-app", &api.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

const (
	myApp = `
		{
			"dataCenters": [
				{
					"dcID": 1	
				},
				{
					"dcID": 2	
				}
			],
			"deployment":{
    "kind": "Deployment",
    "apiVersion": "extensions/v1beta1",
    "metadata": {
        "name": "nginx-deploy-app",
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
        "template": {
            "metadata": {
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
                        ]
                    }
                ]
            }
        }
            }
}
		}	
	`
)
