package main

type ServiceType struct {
	Kind       string            `json: "kind"`
	ApiVersion string            `json: "apiVersion"`
	Items      []ServiceItemType `json: "items"`
}

type ServiceItemType struct {
	Metadata ServiceMetadataType `json: "metadata"`
	Spec     SpecType            `json: "spec"`
}

type ServiceMetadataType struct {
	Name   string           `json: "name"`
	Labels ServiceLabelType `json: "labels"`
}

type ServiceLabelType struct {
	Component string `json: "component"`
	Provider  string `json: "provider"`
}

type SpecType struct {
	Ports []ServicePortType `json: "ports"`
}

type ServicePortType struct {
	Protocol   string `json: "protocol"`
	Port       int    `json: "port"`
	TargetPort int    `json: "targetPort"`
}
