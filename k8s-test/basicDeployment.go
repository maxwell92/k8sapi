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
    Annotations map[string] string `json:"annotations,omitempty"`
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
    Volumes []volumesSTSDL `json:"volumes,omitempty"`
    Containers []containersSTSDL `json:"containers"` 
}

type volumesSTSDL struct {
    Name string `json:"name"`
    HostPath hostPathVSTSDL `json:"hostPath"`
    EmptyDir emptyDirVSTSDL `json:"emptyDir"`
    PersistentVolumeClaim *pvVolumeClaimVSTSDL `json:"persistentVolumeClaim,omitempty"`
    Rbd *rbdVSTSDL `json:"rbd,omitempty"`
    ConfigMap *configMapVSTSDL `json:"configMap,omitempty"`
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
    SecretRef *secretRefRVSTSDL `json:"secretRef,omitempty"`
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
    Command []string `json:"command,omitempty"`
    Args []string `json:"args,omitempty"`
    Ports []portContainer `json:"ports,omitempty"`     
    Env []envContainer `json:"env,omitempty"`
    Resources resourcesContainer `json:"resources,omitempty"`
    VolumeMounts []volumeMountsContainer `json:"volumeMounts,omitempty"` 
    LivenessProbe *livenessProbeContainer `json:"livenessProbe,omitempty"`
    ReadinessProbe *readinessProbeContainer `json:"readinessProbe,omitempty"`
    Lifecycle *lifecycleContainer `json:"lifecycle,omitcycle"`
 
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
    Limits map[string] string `json:"limits,omitempty"`
    Requests map[string] string `json:"requests,omitempty"`
}
    
type volumeMountsContainer struct {
    Name string `json:"name"`
    ReadOnly bool `json:"readOnly"`
    MountPath string `json:"mountPath"` 
}

type livenessProbeContainer struct {
    Exec *execLiveProbeType `json:"exec,omitempty"`
    HttpGet *httpGetLiveProbeType `json:"httpGet,omitempty"`
    TcpSocket *tcpLiveProbeType `json:"tcpSocket,omitempty"`
    InitialDelaySeconds float64 `json:"initialDelaySeconds,omitempty"`
    TimeoutSeconds float64 `json:"timeoutSeconds,omitempty"`
    PeriodSeconds float64 `json:"periodSeconds,omitempty"`
    SuccessThreshold float64 `json:"successThreshold,omitempty"`
    FailureThreshold float64 `json:"failureThreshold,omitempty"`
}

type execLiveProbeType struct {
    Command []string `json:"command,omitempty"` 
}

type httpGetLiveProbeType struct {
    Path string `json:"path,omitempty"`
    Port float64 `json:"port,omitempty"`
    Host string `json:"host,omitempty"`
    Scheme string `json:"scheme,omitempty"`
    HttpHeaders []headersGLPType `json:"httpHeaders,omitempty"`    
}

type headersGLPType struct {
    Name string `json:"name,omitempty"`
    Value string `json:"value,omitempty"`
}

type tcpLiveProbeType struct {
    Port float64 `json:"port,omitempty"`
}

type readinessProbeContainer struct {
    HttpGet *httpGetReadProbeType `json:"httpGet,omitempty"`
    InitialDelaySeconds float64 `json:"initialDelaySeconds,omitempty"`
    TimeoutSeconds float64 `json:"timeoutSeconds,omitempty"`
    PeriodSeconds float64 `json:"periodSeconds,omitempty"`
    SuccessThreshold float64 `json:"successThreshold,omitempty"`
    FailureThreshold float64 `json:"failureThreshold,omitempty"`
}


type httpGetReadProbeType struct {
    Path string `json:"path,omitempty"`
    Port float64 `json:"port,omitempty"`
    Host string `json:"host,omitempty"`
    Scheme string `json:"scheme,omitempty"`
    HttpHeaders []headersGRPType `json:"httpHeaders,omitempty"`    
}

type headersGRPType struct {
    Name string `json:"name"`
    Value string `json:"value"`
}


type lifecycleContainer struct {
    PostStart *postStartLCType `json:"postStart,omitempty"`
    PreStop *preStopLCType `json:"preStop,omitempty"`    
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

