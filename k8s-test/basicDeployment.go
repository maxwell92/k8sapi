package main

type basicDeployment struct {
    ApiVersion string `json:"apiVersion"`
    Kind string `json:"kind"`
    Metadata metadataDL `json:"metadata"`
    Spec specDL `json:"spec"`
}

type metadataDL struct {
    Name string `json:"name"`
    Namespace string `json:"namespace"`
    Labels map[string] string `json:"labels"`
    Annotations map[string] string `json:"annotations"`
}

type specDL struct {
    Replicas float64 `json:"replicas"`
    Template templateSpec `json:"template"`
}

type templateSpec struct {
    Metadata metadataTSDL `json:"metadata"`
    Spec specTSDL `json:"spec"`
}

type metadataTSDL struct {
    Name string `json:"name"`
    Labels map[string] string `json:"labels"`
}

type specTSDL struct {
    Volumes []volumesSTSDL `json:"volumes"`
    Containers []containersSTSDL `json:"containers"` 
}

type volumesSTSDL struct {
    Name string `json:"name"`
    HostPath hostPathVSTSDL `json:"hostPath"`
    EmptyDir emptyDirVSTSDL `json:"emptyDir"`
    PersistentVolumeClaim pvVolumeClaimVSTSDL `json:"persistentVolumeClaim"`
    Rbd rbdVSTSDL `json:"rbd"`
    ConfigMap configMapVSTSDL `json:"configMap"`
}

type hostPathVSTSDL struct {
    Path string `json:"path"`
}

type emptyDirVSTSDL struct {
    Medium string `json:"medium"`
}

type pvVolumeClaimVSTSDL struct {
    ClaimName string `json:"claimName"`
    ReadOnly bool `json:"readOnly"`
}

type rbdVSTSDL struct {
    Monitors []string `json:"monitors"`
    Image string `json:"image"`
    FsType string `json:"fsType"`
    Pool string `json:"pool"`
    User string `json:"user"`
    Keyring string `json:"keyring"`
    SecretRef secretRefRVSTSDL `json:"secretRef"`
    ReadOnly bool `json:"readOnly"` 
}

type secretRefRVSTSDL struct {
    Name string `json: name"`
}

type configMapVSTSDL struct {
    Name string `json:"name"` 
    Items []itemsConfigMap `json:"items"`
}

type itemsConfigMap struct {
    Key string `json:"key"`
    Path string `json:"Path"`
}

type containersSTSDL struct {
    Name string `json:"name"`
    Image string `json:"image"`
    Command []string `json:"command"`
    Args []string `json:"args"`
    Ports []portContainer `json:"ports"`     
    Env []envContainer `json:"env"`
    Resources resourcesContainer `json:"resources"`
    VolumeMounts []volumeMountsContainer `json:"volumeMounts"` 
    LivenessProbe livenessProbeContainer `json:"livenessProbe"`
    ReadinessProbe readinessProbeContainer `json:"readinessProbe"`
    Lifecycle lifecycleContainer `json:"lifecycle"`
 
}

type portContainer struct {
    Name string `json:"name"`
    HostPort float64 `json:"hostPort"`
    ContainerPort float64 `json:"containerPort"`
    Protocol string `json:"protocol"`
    HostIP string `json:"hostIP"` 
}

type envContainer struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

type resourcesContainer struct {
    Limits map[string] string `json:"limits"`
    Requests map[string] string `json:"requests"`
}
    
type volumeMountsContainer struct {
    Name string `json:"name"`
    ReadOnly bool `json:"readOnly"`
    MountPath string `json:"mountPath"` 
}

type livenessProbeContainer struct {
    Exec execLiveProbeType `json:"exec"`
    HttpGet httpGetLiveProbeType `json:"httpGet"`
    TcpSocket tcpLiveProbeType `json:"tcpSocket"`
    InitialDelaySeconds float64 `json:"initialDelaySeconds"`
    TimeoutSeconds float64 `json:"timeoutSeconds"`
    PeriodSeconds float64 `json:"periodSeconds"`
    SuccessThreshold float64 `json:"successThreshold"`
    FailureThreshold float64 `json:"failureThreshold"`
}

type execLiveProbeType struct {
    Command []string `json:"command"` 
}

type httpGetLiveProbeType struct {
    Path string `json:"path"`
    Port float64 `json:"port"`
    Host string `json:"host"`
    Scheme string `json:"scheme"`
    HttpHeaders []headersGLPType `json:"httpHeaders"`    
}

type headersGLPType struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

type tcpLiveProbeType struct {
    Port float64 `json:"port"`
}

type readinessProbeContainer struct {
    HttpGet httpGetReadProbeType `json:"httpGet"`
    InitialDelaySeconds float64 `json:"initialDelaySeconds"`
    TimeoutSeconds float64 `json:"timeoutSeconds"`
    PeriodSeconds float64 `json:"periodSeconds"`
    SuccessThreshold float64 `json:"successThreshold"`
    FailureThreshold float64 `json:"failureThreshold"`
}


type httpGetReadProbeType struct {
    Path string `json:"path"`
    Port float64 `json:"port"`
    Host string `json:"host"`
    Scheme string `json:"scheme"`
    HttpHeaders []headersGRPType `json:"httpHeaders"`    
}

type headersGRPType struct {
    Name string `json:"name"`
    Value string `json:"value"`
}


type lifecycleContainer struct {
    PostStart postStartLCType `json:"postStart"`
    PreStop preStopLCType `json:"preStop"`    
}

type postStartLCType struct {
    Exec execPSLCType `json:"exec"`
}

type execPSLCType struct {
    Command []string `json:"command"`
}




type preStopLCType struct {
    Exec execPrSLCType `json:"exec"`
}

type execPrSLCType struct {
    Command []string `json:"command"`
}
