package endpoint

// Endpoint for Get @/api/v1/namespaces/{namespace}/endpoints
type EndpointList struct {
	Kind       string           `json:"kind"`
	ApiVersion string           `json:"apiVersion"`
	Metadata   MetadataListType `json:"metadata"`
	Items      []ItemsListType  `json:"items"`
}

type MetadataListType struct {
	SelfLink        string `json:"selfLink,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
}

type Endpoint ItemsListType

// Endpoint for Post @/api/v1/namespaces/{namespace}/endpoints
type ItemsListType struct {
	Kind       string      `json:"kind,omitempty"`
	ApiVersion string      `json:"apiVersion,omitempty"`
	Metadata   MetadataEL  `json:"metadata"`
	SubSets    []SubSetsEL `json:"subsets"`
}

type MetadataEL struct {
	Name                       string            `json:"name"`
	GenerateName               string            `json:"generateName,omitempty"`
	Namespace                  string            `json:"namespace,omitempty"`
	SelfLink                   string            `json:"selfLink,omitempty"`
	Uid                        string            `json:"uid,omitempty"`
	ResourceVersion            string            `json:"resourceVersion,omitempty"`
	Generation                 float64           `json:"generation,omitempty"`
	CreationTimestamp          string            `json:"creationTimestamp,omitempty"`
	DeletionTimestamp          string            `json:"deletionTimestamp,omitempty"`
	DeletionGracePeriodSeconds float64           `json:"deletionGracePeriodSeconds,omitempty"`
	Labels                     map[string]string `json:"labels"`
	Annotations                map[string]string `json:"annotations,omitempty"`
}

type SubSetsEL struct {
	Addresses         []AddressesType         `json:"addresses,omitempty"`
	NotReadyAddresses []NotReadyAddressesType `json:"notReadyAddressesType,omitempty"`
	Ports             []PortsType             `json:"ports,omitempty"`
}

type AddressesType struct {
	IP        string         `json:"ip"`
	TargetRef *TargetRefType `json:"targetRef,omitempty"`
}

type TargetRefType struct {
	Kind            string `json:"kind,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
	Name            string `json:"name"`
	Uid             string `json:"uid",omitempty`
	ApiVersion      string `json:"apiVersion,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
	FieldPath       string `json:"fieldPath"`
}

type NotReadyAddressesType struct {
	IP        string         `json:"ip"`
	TargetRef *TargetRefType `json:"targetRef,omitempty"`
}

type PortsType struct {
	Name     string  `json:"name,omitempty"`
	Port     float64 `json:"port"`
	Protocol string  `json:"protocol,omitempty"`
}
