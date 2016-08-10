package replicaset

import (
	"encoding/json"
	"fmt"
	hc "httpclient"
	"log"
	replicaset "model/replicaset"
	"strings"
)

type ReplicaSetController struct {
	Name string
}

func (rc *ReplicaSetController) Get(url string) {

	// TODO:
	// 1. Validation of url and make the right url
	// 2. Return the response json

	var rclist replicaset.ReplicaSetList
	client := hc.NewHttpClient("", "")

	response, err := client.Get(url)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(response, &rclist)
	if err != nil {
		log.Println(err)
	}

	num := len(rclist.Items)
	fmt.Println(num)

	for i := 0; i < num; i++ {
		fmt.Println(rclist.Items[i].Metadata.Name)
	}
}

func (rc *ReplicaSetController) Post(url string) {
	// TODO:
	// 1. Validation of both json and url
	// 2. Return the response json

	client := hc.NewHttpClient("", "")

	labels := make(map[string]string, 1)
	labels["name"] = "nginx-rs-test"

	rp := new(replicaset.ReplicaSet)
	rp.ApiVersion = "extensions/v1beta1"
	rp.Kind = "ReplicaSet"
	rp.Metadata.Name = "nginx-rs-test"
	rp.Spec.Replicas = 3
	rp.Spec.Template.Metadata.Labels = labels
	rp.Spec.Template.Spec.Containers = make([]replicaset.ContainerType, 1)
	rp.Spec.Template.Spec.Containers[0].Name = "nginx-rs-test"
	rp.Spec.Template.Spec.Containers[0].Image = "nginx:1.7.9"

	result, _ := json.MarshalIndent(rp, "", " ")
	rep, err := client.Post(url, strings.NewReader(string(result)))
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(rep))
	}
}

func (rc *ReplicaSetController) Delete(url string) {
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

func (rc *ReplicaSetController) Describe(url string) {

}

func (rc *ReplicaSetController) EncodeJson() {

}

func (rc *ReplicaSetController) DecodeJson() {

}
