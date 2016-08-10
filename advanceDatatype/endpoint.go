package main

//import "fmt"
type EndpointType struct {
	Kind       string     `json: "kind"`
	ApiVersion string     `json: "apiVersion"`
	Items      []ItemType `json: "items"`
}

type ItemType struct {
	Metadata MetadataType `json: "metadata"`
	Subsets  []SubsetType `json: "subsets"`
}

type MetadataType struct {
	Name              string    `json: "name"`
	Namespace         string    `json: "namespace"`
	CreationTimestamp string    `json: "creationTimestamp"`
	Labels            LabelType `json: "labels"`
}

type LabelType struct {
	Name string `json: "name"`
}

type SubsetType struct {
	Addresses []AddressType `json: "addresses"`
	Ports     []PortType    `json: "ports"`
}

type AddressType struct {
	Ip string `json: "ip"`
}

type PortType struct {
	Port int `json: "port"`
}

/*
func main() {
	var ep EndpointTypeRemote
	its := make([]EndpointType, 1)
	its[0].Kind="the kind"

	ep.Items=its
	fmt.Println(ep)
}*/
