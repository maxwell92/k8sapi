package pod

import (
	"encoding/json"
	"fmt"
	hc "httpclient"
	"log"
	pod "model/pod"
	"strings"
)

type PodController struct {
	Name string
}

func (pc *PodController) Get(url string) {

	// TODO:
	// 1. Validation of url and make the right url
	// 2. Return the response json

	var plist pod.PodList
	client := hc.NewHttpClient("", "")

	response, err := client.Get(url)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(response, &plist)
	if err != nil {
		log.Println(err)
	}

	num := len(plist.Items)
	fmt.Println(num)

	for i := 0; i < num; i++ {
		fmt.Println(plist.Items[i].Metadata.Name)
	}
}

func (pc *PodController) Post(url string) {
	// TODO:
	// 1. Validation of both json and url
	// 2. Return the response json

	client := hc.NewHttpClient("", "")

	labels := make(map[string]string, 1)
	labels["name"] = "nginx-pd-test"

	pd := new(pod.Pod)
	pd.ApiVersion = "v1"
	pd.Kind = "Pod"
	pd.Metadata.Name = "nginx-pd-test"
	pd.Metadata.Labels = labels
	pd.Spec.Containers = make([]pod.ContainerSType, 1)
	pd.Spec.Containers[0].Name = "nginx-pd-test"
	pd.Spec.Containers[0].Image = "nginx:1.7.9"

	result, _ := json.MarshalIndent(pd, "", " ")
	rep, err := client.Post(url, strings.NewReader(string(result)))
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(rep))
	}
}

func (pc *PodController) Delete(url string) {
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

func (pc *PodController) Describe(url string) {

}

func (pc *PodController) EncodeJson() {

}

func (pc *PodController) DecodeJson() {

}
