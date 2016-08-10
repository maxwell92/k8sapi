package endpointcontroller

import (
	"encoding/json"
	"fmt"
	hc "httpclient"
	"log"
	endpoint "model/endpoint"
	"strings"
)

type EndpointController struct {
	Name string
}

func (ec *EndpointController) Get(url string) {
	//TODO:

	var eplist endpoint.EndpointList
	client := hc.NewHttpClient("", "")

	response, err := client.Get(url)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(response, &eplist)
	if err != nil {
		log.Println(err)
	}

	num := len(eplist.Items)
	fmt.Println(num)

	for i := 0; i < num; i++ {
		fmt.Println(eplist.Items[i].Metadata.Name)
	}

}

func (ec *EndpointController) Post(url string) {
	client := hc.NewHttpClient("", "")

	labels := make(map[string]string, 1)
	labels["name"] = "nginx-ep-test"

	ep := new(endpoint.Endpoint)
	ep.ApiVersion = "v1"
	ep.Kind = "Endpoints"
	ep.Metadata.Name = "nginx-ep-test"
	ep.Metadata.Labels = labels
	ep.SubSets = make([]endpoint.SubSetsEL, 1)
	ep.SubSets[0].Addresses = make([]endpoint.AddressesType, 1)
	ep.SubSets[0].Addresses[0].IP = "192.168.1.100"
	ep.SubSets[0].Ports = make([]endpoint.PortsType, 1)
	ep.SubSets[0].Ports[0].Port = 8080

	result, _ := json.MarshalIndent(ep, "", " ")
	fmt.Println(string(result))
	rep, err := client.Post(url, strings.NewReader(string(result)))
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(rep))
	}

}
func (ec *EndpointController) Delete(url string) {
	client := hc.NewHttpClient("", "")
	resp, err := client.Delete(url)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(resp))
	}
}

func (ec *EndpointController) Describe() {

}

func (ec *EndpointController) EncodeJson() {

}
func (ec *EndpointController) DecodeJson() {

}
