package service

// ServiceList for Get @/api/v1/namespaces/{namespace}/services
type ServiceList struct {
    Kind string `json:"kind"`
    ApiVersion string `json:"apiVersion"`
    Metadata MetadataListType `json:"metadata"`
    Items []ItemsListType `json:"items"`
}

type MetadataListType struct {
    SelfLink string `json:"selfLink"`
    ResourceVersion string `json:"resourceVersion"`
}

type ItemsListType Service

// Service for Post @/api/v1/namespaces/{namespace}/services
type ItemsListType struct {
    Kind string `json:"kind"`
    ApiVersion string `json:"apiVersion"`
    Metadata MetadataSL `json:"metadata"`
    Spec SpecSL `json:"spec"`
    Status StatusSL `json:"status"`
}

type MetadataSL struct {
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
    Annotations map[string] string `json:"annotations"`
}

type SpecSL struct {
    Ports []PortsSL `json:"ports"`
    Selector map[string] string `json:"selector"`
    ClusterIP string `json:"clusterIP"`
    Type string `json:"type"`
    ExternalIPs []string `json:"externalIPs"`
    DeprecatedPublicIPs []string `json:"deprecatedPublicIPs"`
    SessionAffinity string `json:"sessionAffinity"`
    LoadBalancerIP string `json:"loadBalancerIP"`
}

type PortsSL struct {
    Name string `json:"name"`
    Protocol string `json:"protocol"`
    Port float64 `json:"port"`
    TargetPort float64 `json:"targetPort"`
    NodePort float64 `json:"nodePort"` 
}

type StatusSL struct {
    LoadBalancer LBStatusSL `json:"loadBalancer"`
}

type LBStatusSL struct {
    Ingress []IngressLBSSL `json:"ingress"`
}

type Ingress struct {
    IP string `json:"ip"`
    Hostname string `json:"hostname"`
}
