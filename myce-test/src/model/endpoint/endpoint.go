package endpoint

// Endpoint for Get @/api/v1/namespaces/{namespace}/endpoints
type EndpointList struct {
    Kind string `json:"kind"`
    ApiVersion string `json:"apiVersion"`
    Metadata MetadataListType `json:"metadata"`
    Items []ItemsListType `json:"items"`
}

type MetadataListType struct {
    SelfLink string `json:"selfLink"`
    ResourceVersion string `json:"resourceVersion"`
}

type ItemsListType Endpoint

// Endpoint for Post @/api/v1/namespaces/{namespace}/endpoints
type ItemsListType struct {
    Kind string `json:"kind"`
    ApiVersion string `json:"apiVersion"`
    Metadata MetadataEL `json:"metadata"`
    SubSets SubSetsEL `json:"subsets"`
}

type MetadataEL struct {
    Name string `json:"name"`
    GenerateName string `json:generateName"`
    Namespace string `json:"namespace"`
    SelfLink string `json:"selfLink"`
    Uid string `json:uid"`
    ResourceVersion string `json:"resourceVersion"`
    Generation float64 `json:"generation"`
    CreationTimestamp string `json:"creationTimestamp"`
    DeletionTimestamp string `json:"deletionTimestamp"`
    DeletionGracePeriodSeconds float64 `json:"deletionGracePeriodSeconds"`
    Labels map[string] string `json:"labels"`
    Annotations map[string] string `json:"annotations"
}

type SubsetsEL struct {
    Addresses []AddressesType `json:"addresses"`
    NotReadyAddresses []NotReadyAddressesType `json:"notReadyAddressesType"`
    Ports []PortsType `json:"ports"`
} 

type AddressesType struct {
    IP string `json:"ip"`
    TargetRef TargetRefType `json:"targetRef"` 
}

type TargetRefType struct {
    Kind string `json:"kind"`
    Namespace string `json:"namespace"`
    Name string `json:"name"`
    Uid string `json:"uid"`
    ApiVersion string `json:"apiVersion"`
    ResourceVersion string `json:"resourceVersion"`
    FieldPath string `json:"fieldPath"`
}

type NotReadyAddressesType struct {
    IP string `json:"ip"`
    TargetRef TargetRefType `json:"targetRef"`
}

type PortsType struct {
    Name string `json:"name"`
    Port float64 `json:"port"`
    Protocol string `json:"protocol"`
}
