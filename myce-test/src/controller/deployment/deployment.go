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
	}
}

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

func (dc *DeploymentController) Describe(url string) {

}

func (dc *DeploymentController) EncodeJson() {

}

func (dc *DeploymentController) DecodeJson() {

}
