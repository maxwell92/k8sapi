package deployment

import (
	"encoding/json"
	"fmt"
	hc "httpclient"
	"log"
	deploy "model/deployment"

	"strings"
)

type DeploymentController struct {
	Name string
}

func (dc *DeploymentController) Get(url string) {

	// TODO:
	// 1. Validation of url and make the right url
	// 2. Return the response json

	var dplist deploy.DeploymentList
	client := hc.NewHttpClient("", "")

	response, err := client.Get(url)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(response, &dplist)
	if err != nil {
		log.Println(err)
	}

	num := len(dplist.Items)
	fmt.Println(num)

	for i := 0; i < num; i++ {
		fmt.Println(dplist.Items[i].Metadata.Name)
		fmt.Println(dplist.Items[i].Metadata.CreationTimestamp)
	}
}

/*
func (dc *DeploymentController) Post(url string) {
	// TODO:
	// 1. Validation of both json and url
	// 2. Return the response json

	client := hc.NewHttpClient("", "")

	labels := make(map[string]string, 1)
	labels["name"] = "nginx-test"

	dp := new(deploy.Deployment)
	dp.ApiVersion = "extensions/v1beta1"
	dp.Kind = "Deployment"
	dp.Metadata.Name = "nginx-test"
	dp.Spec.Replicas = 3
	dp.Spec.Template.Metadata.Labels = labels
	dp.Spec.Template.Spec.Containers = make([]deploy.ContainerType, 1)
	dp.Spec.Template.Spec.Containers[0].Name = "nginx-test"
	dp.Spec.Template.Spec.Containers[0].Image = "nginx:1.7.9"

	result, _ := json.MarshalIndent(dp, "", " ")
	rep, err := client.Post(url, strings.NewReader(string(result)))
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(rep))
	}
}
*/

var myDeploy = `
{
	"kind": "Deployment",
	"apiVersion": "extensions/v1beta1",
	"metadata": {
		"name": "nginx-test",
		"labels": {
			"name": "nginx-test",
			"maintainer": "maxwell"
		}
	},
	"spec": {
		"replicas": 3,
		"template": {
			"metadata": {
				"labels": {
					"name": "spec-template-metadata-labels"	
				}	
			},	
			"spec": {
				"containers": [{
					"name": "nginx-test",
					"image": "nginx:1.7.9"
				}]	
			}
		}
	}
}
`

var myApp = `
{
	"name": "nginx-test",
	"namespace": "default",
	"datacenter": [
		"shijihulian"
	],
	"image": "nginx:1.7.9",
	"replicas": 3,
	"spec": {
		"request": {
			"cpu": "2",
			"mem": "512"
		}			
	},
	"configFile": [
		"/root/nginx.conf"	
	],
	"mountPoints": {
		"name": "disk1",
		"mountPath": "/usr/local/nginx/html"
	},
	"healthCheck": {
		"httpGet": {
			"path": "http://api/v1/healthz",
			"port": 11000
		},
		"initialDelaySeconds": 3,
		"periodSeconds": 2
	},
	"readiness": {
		"exec": {
			"echo readiness"	
		}	
	},
	"postStart": {
		"exec": {
			"echo postStart"	
		}	
	},
	"preStop": {
		"exec": {
			"echo preStop"	
		}	
	},
	"env": {
		"magic": "good",
		"sheep": "mushroom"
	},
	"command": [
		"echo"	
	],
	"args": [
		"abc"	
	]
}			
`

func (dc *DeploymentController) Post(url string) {
	client := hc.NewHttpClient("", "")
	rep, err := client.Post(url, strings.NewReader(string(myDeploy)))
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(rep))
	}
}
func (dc *DeploymentController) Delete(url string) {
	// TODO:
	// 1. Validation of url
	// 2. Return the respone json

	client := hc.NewHttpClient("", "")
	resp, err := client.Delete(url)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(resp))
	}
}

func (dc *DeploymentController) DeployApp(url string) {
	client := hc.NewHttpClient("", "")
	resp, err := client.Post(url, strings.NewReader(string(myApp)))
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(resp))
	}
}

func (dc *DeploymentController) Handle() ([]byte, error) {

	//TODO: Unmarshal
	//appDeploy := new(deploy.AppDeployment)
	var app deploy.AppDeployment
	//err := json.Unmarshal(myApp, &app)
	myAppByte, err := strings.NewReader(myApp).ReadByte()
	if err != nil {
		log.Println(err)
	}

	err := json.Unmarshal(myAppByte, &app)
	if err != nil {
		log.Println(err)
	}

	// Validation of dc
	/* for k, v := range app.Datacenter {
		mysql.Query(v)
	}
	*/
	// for each dc
	// for k, v := range app.Datacenter
	var dp deploy.Deployment // or new()?
	dp.ApiVersion = "extensions/v1beta1"
	dp.Kind = "Deployment"
	dp.Metadata.Name = app.Name
	dp.Metadata.Namespace = app.Namespace
	dp.Metadata.Labels = app.Labels
	dp.Spec.Replicas = app.Replicas
	dp.Spec.Template.Metadata.Name = app.Name
	dp.Spec.Template.Metadata.Labels = app.Labels
	dp.Spec.Template.Spec.Containers[0].Name = app.Name
	dp.Spec.Template.Spec.Containers[0].Image = app.Image
	dp.Spec.Template.Spec.Containers[0].Command = app.Command
	dp.Spec.Template.Spec.Containers[0].Args = app.Args
	dp.Spec.Template.Spec.Containers[0].Env = app.Env
	dp.Spec.Template.Spec.Containers[0].Resources.Requests = app.Spec.Request
	dp.Spec.Template.Spec.Containers[0].VolumeMounts = app.MountPoints
	dp.Spec.Template.Spec.Containers[0].LivenessProbe = app.HealthCheck
	dp.Spec.Template.Spec.Containers[0].ReadinessProbe = app.Readiness
	dp.Spec.Template.Spec.Containers[0].Lifecycle.PostStart = app.PostStart
	dp.Spec.Template.Spec.Containers[0].Lifecycle.PreStop = app.PreStop

	// post to k8s
	var result []byte
	client := hc.NewHttpClient("", "")
	url := "http://master:8080/apis/extensions/v1beta1/namespaces/default/deployments"
	result, _ = json.MarshalIndent(dp, "", "  ")
	resp, err := client.Post(url, strings.NewReader(string(result)))
	return resp, err
	// write back the response
}

func (dc *DeploymentController) Describe(url string) {

}

func (dc *DeploymentController) EncodeJson() {

}

func (dc *DeploymentController) DecodeJson() {

}
