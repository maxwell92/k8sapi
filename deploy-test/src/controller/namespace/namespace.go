package namespacecontroller

import (
	"encoding/json"
	"fmt"
	hc "httpclient"
	"log"
	namespace "model/namespace"
	"strings"
)

type NamespaceController struct {
	Name string
}

func (nc *NamespaceController) Get(url string) {
	var nslist namespace.NamespaceList
	client := hc.NewHttpClient("", "")

	response, err := client.Get(url)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(response, &nslist)
	if err != nil {
		log.Println(err)
	}

	num := len(nslist.Items)
	fmt.Println(num)

	for i := 0; i < num; i++ {
		fmt.Println(nslist.Items[i].Metadata.Name)
	}
}

func (nc *NamespaceController) Post(url string) {
	client := hc.NewHttpClient("", "")

	labels := make(map[string]string, 1)
	labels["name"] = "ops"

	ns := new(namespace.Namespace)
	ns.ApiVersion = "v1"
	ns.Kind = "Namespace"
	ns.Metadata.Name = "ops"
	ns.Metadata.Labels = labels

	result, _ := json.MarshalIndent(ns, "", "")
	fmt.Println(string(result))
	rep, err := client.Post(url, strings.NewReader(string(result)))
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(rep))
	}
}

func (nc *NamespaceController) Delete(url string) {
	client := hc.NewHttpClient("", "")
	resp, err := client.Delete(url)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(resp))
	}
}

func (nc *NamespaceController) Describe() {

}

func (nc *NamespaceController) EncodeJson() {

}

func (nc *NamespaceController) DecodeJson() {

}
