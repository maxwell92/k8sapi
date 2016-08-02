package namespace

// NamespaceList for Get @/api/v1/namespaces
type NamespaceList struct {
    Kind string `json:"kind"`
    ApiVersion string `json:"apiVersion"`
    Metadata MetadataListType `json:"metadata"`
    Items []ItemsListType `json:"items"`
}

type MetadataListType struct {
    SelfLink string `json:"selfLink"`
    ResourceVersion string `json:"resourceVersion"`
}

type ItemsListType Namespace 

// Namespace for Post @/api/v1/namespace
type Namespace struct {
    Kind string `json:"kind"`
    ApiVersion string `json:"apiVersion"`
    Metadata MetadataType `json:"metadata"`
    Spec SpecType `json:"spec"`
    Status *StatusType `json:"status,omitempty"`
}

type MetadataType struct {
    Name string `json:"name"`
    GenerateName string `json:"generateName,omitempty"`
    Namespace string `json:"namespace,omitempty"`
    SelfLink string `json:"selfLink,omitempty"`
    Uid string `json:"uid,omitempty"`
    ResourceVersion string `json:"resourceVersion,omitempty"`
    Generation float64 `json:"generation,omitempty"`
    CreationTimestamp string `json:"creationTimestamp,omitempty"`
    DeletionTimestamp string `json:"deletionTimestamp,omitempty"`
    DeletionGracePeriodSeconds float64 `json:"deletionGracePeriodSeconds,omitempty"`
    Labels map[string] string `json:"labels"`
    Annotations map[string] string `json:"annotations,omitempty"`
}

type SpecType struct {
//   Finalizers []FinalizersType `json:finalizers"`
    Finalizers []string `json:"finalizers,omitempty"`
}
/*
type FinalizersType struct {

}
*/
type StatusType struct {
    Phase string `json:"phase,omitempty"`
}
