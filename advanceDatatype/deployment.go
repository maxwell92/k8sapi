package deployment

type DeploymentList struct {
    ApiVersion
}


// Deployment for Post /apis/extensions/v1beta1/namespaces/xxx/deployments

type Deployment struct {
    ApiVersion string `json:"apiVersion"`
    Kind string `json:"kind"`
    Metadata MetadataType `json:"metadata"`
    Spec SpecType `json:"spec"`
}

type MetadataType struct {
    Name string `json:"name"`
    Namespace string `json:"namespace"`
    Labels map[string] string `json:"labels"`
    Annotations map[string] string `json:"annotations,omitempty"`
}

type SpecType struct {
    Replicas float64 `json:"replicas"`
    Template TemplateSpec `json:"template"`
}

type TemplateSpec struct {
    Metadata MetadataTS `json:"metadata"`
    Spec SpecTS `json:"spec"`
}

type MetadataTS struct {
    Name string `json:"name"`
    Labels map[string] string `json:"labels"`
}

type SpecTS struct {
    Volumes []VolumesSTS `json:"volumes,omitempty"`
    Containers []ContainersSTS `json:"containers"` 
}

type VolumesSTS struct {
    Name string `json:"name"`
    HostPath *HostPathVSTS `json:"hostPath,omitempty"`
    EmptyDir *EmptyDirVSTS `json:"emptyDir,omitempty"`
    PersistentVolumeClaim *PvClaimVSTS `json:"persistentVolumeClaim,omitempty"`
    Rbd *RbdVSTS `json:"rbd,omitempty"`
    ConfigMap *ConfigMapVSTS `json:"configMap,omitempty"`
}

type HostPathVSTS struct {
    Path string `json:"path"`
}

type EmptyDirVSTS struct {
    Medium string `json:"medium"`
}

type PvClaimVSTS struct {
    ClaimName string `json:"claimName"`
    ReadOnly bool `json:"readOnly"`
}

type RbdVSTS struct {
    Monitors []string `json:"monitors"`
    Image string `json:"image"`
    FsType string `json:"fsType"`
    Pool string `json:"pool"`
    User string `json:"user"`
    Keyring string `json:"keyring"`
    SecretRef *SecretRefRVSTS `json:"secretRef"`
    ReadOnly bool `json:"readOnly"` 
}

type SecretRefRVSTS struct {
    Name string `json: name"`
}

type ConfigMapVSTS struct {
    Name string `json:"name"` 
    Items []ItemsConfigMap `json:"items"`
}

type ItemsConfigMap struct {
    Key string `json:"key"`
    Path string `json:"Path"`
}


type ContainersSTS struct {
    Name string `json:"name"`
    Image string `json:"image"`
    Command []string `json:"command,omitempty"`
    Args []string `json:"args,omitempty"`
    Ports []PortContainer `json:"ports"`     
    Env []EnvContainer `json:"env,omitempty"`
    Resources *ResourcesContainer `json:"resources,omitempty"`
    VolumeMounts []VolumeMountsContainer `json:"volumeMounts,omitempty"` 
    LivenessProbe *LivenessProbeContainer `json:"livenessProbe,omitempty"`
    ReadinessProbe *ReadinessProbeContainer `json:"readinessProbe,omitempty"`
    Lifecycle *LifecycleContainer `json:"lifecycle,omitempty"`
 
}

type ContainerType ContainersSTS 

type PortContainer struct {
    Name string `json:"name"`
    HostPort float64 `json:"hostPort"`
    ContainerPort float64 `json:"containerPort"`
    Protocol string `json:"protocol"`
    HostIP string `json:"hostIP"` 
}

type EnvContainer struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

type ResourcesContainer struct {
    Limits map[string] string `json:"limits"`
    Requests map[string] string `json:"requests"`
}
    
type VolumeMountsContainer struct {
    Name string `json:"name"`
    ReadOnly bool `json:"readOnly"`
    MountPath string `json:"mountPath"` 
}

type LivenessProbeContainer struct {
    Exec *ExecLiveProbe `json:"exec,omitempty"`
    HttpGet *HttpGetLiveProbe `json:"httpGet,omitempty"`
    TcpSocket *TcpLiveProbe `json:"tcpSocket,omitempty"`
    InitialDelaySeconds float64 `json:"initialDelaySeconds"`
    TimeoutSeconds float64 `json:"timeoutSeconds"`
    PeriodSeconds float64 `json:"periodSeconds"`
    SuccessThreshold float64 `json:"successThreshold"`
    FailureThreshold float64 `json:"failureThreshold"`
}

type ExecLiveProbe struct {
    Command []string `json:"command"` 
}

type HttpGetLiveProbe struct {
    Path string `json:"path"`
    Port float64 `json:"port"`
    Host string `json:"host"`
    Scheme string `json:"scheme"`
    HttpHeaders []HeadersGLP `json:"httpHeaders"`    
}

type HeadersGLP struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

type TcpLiveProbe struct {
    Port float64 `json:"port"`
}

type ReadinessProbeContainer struct {
    HttpGet *HttpGetReadProbe `json:"httpGet,omitempty"`
    InitialDelaySeconds float64 `json:"initialDelaySeconds"`
    TimeoutSeconds float64 `json:"timeoutSeconds"`
    PeriodSeconds float64 `json:"periodSeconds"`
    SuccessThreshold float64 `json:"successThreshold"`
    FailureThreshold float64 `json:"failureThreshold"`
}


type HttpGetReadProbe struct {
    Path string `json:"path"`
    Port float64 `json:"port"`
    Host string `json:"host"`
    Scheme string `json:"scheme"`
    HttpHeaders []HeadersGRP `json:"httpHeaders"`    
}

type HeadersGRP struct {
    Name string `json:"name"`
    Value string `json:"value"`
}


type LifecycleContainer struct {
    PostStart *PostStartLC `json:"postStart,omitempty"`
    PreStop *PreStopLC `json:"preStop,omitempty"`    
}

type PostStartLC struct {
    Exec *ExecPSLC `json:"exec,omitempty"`
}

type ExecPSLC struct {
    Command []string `json:"command"`
}

type PreStopLC struct {
    Exec *ExecPrSLC `json:"exec,omitempty"`
}

type ExecPrSLC struct {
    Command []string `json:"command"`
}

