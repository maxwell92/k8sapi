package replicaset


// ReplicaSetList for Get @/apis/extensions/v1beta1/namespaces/{namespace}/replicasets
type ReplicaSetList struct {
    Kind string `json:"kind"`
    ApiVersion string `json:"apiVersion"`
    Metadata MetadataListType `json:"metadata"`
    Items []ItemsListType `json:"items"`
} 

type MetadataListType struct {
    SelfLink string `json:"selfLink"`
    ResourceVersion string `json:"resourceVersion"`
}


type ItemsListType ReplicaSet

// ReplicaSet for Post @/apis/extensions/v1beta1/namespaces/{namespace}/replicasets
type ItemsListType struct {
    Kind string `json:"kind"`
    ApiVersion string `json:"apiVersion"`
    Metadata MetadataIT `json:"metadata"`
    Spec SpecType `json:"spec"`
    Status StatusType `json:"status"`
}


type MetadataIT struct {
    Name string `json:"name"`
    GenerateName string `json:"generateName"`
    Namespace string `json:"namespace"`
    SelfLink string `json:"selfLink"`
    Uid string `json:"uid"` 
    ResourceVersion string `json:"resourceVersion"`
    Generation float64 `json:"generation"`
    CreationTimestamp string `json:"creationTimestamp"`
    DeletionTimestamp string `json:"deletionTimestamp"`
    DeletionGracePeriodSecond float64 `json:"deletionGracePeriodSeconds"`
    Labels map[string] string `json:"labels"`
    Annotations map[string] string `json:"annotations"` 
}


type SpecType struct {
    Replicas float64 `json:"replicas"`
    Selector SelectorType `json:"selector"`
    Template TemplateType `json:"template"`
}

type SelectorType struct {
    MatchLabels map[string] string `json:"matchLabels"`
    MatchExpressions []MatchExpType `json:"matchExpressions"`
}

type MatchExpType struct {
    Key string `json:"key"`
    Operator string `json:"operator"`
    Values []string `json:"values"`
}

type TemplateType struct {
    Metadata MetadataTemplateType `json:"metadata"`
    Spec SpecTemplateType `json:"spec"`
}

type MetadataTemplateType struct {
    Name string `json:"name"`
    GenerateName string `json:"generateName"`
    Namespace string `json:"namespace"`
    SelfLink string `json:"selfLink"`
    Uid string `json:"uid"`
    ResourceVersion string `json:"resourceVersion"`
    Generation float64 `json:"generation"`
    CreationTimestamp string `json:"creationTimestamp"`
    DeletionTimestamp string `json:"deletionTimestamp"`
    DeletionGracePeriodSeconds float64 `json:"deletionGracePeriodSeconds"`
    Labels map[string] string `json:"labels"`
    Annotations map[string] string `json:"annotations"`
}


type SpecTemplateType struct {
    Volumes []VolumesSTT `json:"volumes"`
    Containers []ContainersSTT `json:"containers"`
    RestartPolicy string `json:"restartPolicy"`
    TerminationGracePeriodSeconds float64 `json:"terminationGracePeriodSeconds"`
    ActiveDeadlineSeconds float64 `json:"activeDeadlineSeconds"`
    DnsPolicy string `json:"dnsPolicy"`
    NodeSelector map[string] string `json:"nodeSelector"`
    NodeName string `json:"nodeName"`
    HostNetwork bool `json:"hostNetwork"`
    HostPID bool `json:"hostPID"`
    HostIPC bool `json:"hostIPC"`
}

type VolumesSTT struct {
    Name string `json:"name"`
    HostPath HostPathVSTT `json:"hostPath"`
    EmptyDir EmptyDirVSTT `json:"emptyDir"`
    PersistentVolumeClaim PVClaimVSTT `json:"persistentVolumeClaim"`  
    Rbd RbdVSTT `json:"rbd"`
    ConfigMap ConfigMapVSTT `json:"configMap"`
}

type HostPathVSTT struct {
    Path string `json:"path"` 
}

type EmptyDirVSTT struct {
    Medium string `json:"medium"`
}

type PVClaimVSTT struct {
    ClaimName string `json:"claimName"`
    ReadOnly bool `json:"readOnly"`
}

type RbdVSTT struct {
    Monitors []string `json:"monitors"`
    Image string `json:"image"`
    FsType string `json:"fsType"`
    Pool string `json:"pool"`
    User string `json:"user"`
    Keyring string `json:"keyring"`
    SecretRef SecretRefRVSTT `json:"secretRef"`
    ReadOnly bool `json:"readOnly"`
}

type SecretRefRVSTT struct {
    Name string `json:"name"`
}

type ConfigMapVSTT struct {
    Name string `json:"name"`
    Items []ItemsCVSTT `json:"items"`
}

type ItemsCVSTT struct {
    Key string `json:"key"`
    Path string `json:"path"`
}


type ContainersSTT struct {
    Name string `json:"name"`
    Image string `json:"image"`
    Command []string `json:"command"`
    Args []string `json:"args"`
    WorkingDir string `json:"workingDir"`
    Ports []PortsCSTT `json:"ports"`
    Env []EnvCSTT `json:"env"`
    Resources ResourceCSTT `json:"resource"`
    VolumeMounts []VolumeMountsCSTT `json:"volumeMounts"`  
    LivenessProbe LivenessProbeCSTT `json:"livenessProbe"`
    ReadinessProbe ReadinessProbeCSTT `json:"readinessProbe"`
    LifeCycle LifeCycleCSTT `json:"lifecycle"`
    TerminationMessagePath string `json:"terminationMessagePath"`
    ImagePullPolicy string `json:"imagePullPolicy"`
    Stdin bool `json:"stdin"`
    StdinOnce bool `json:"stdinOnce"`
    Tty bool `json:"tty"`
}

type PortsCSTT struct {
    Name string `json:"name"`
    HostPort float64 `json:"hostPort"`
    ContainerPort float64 `json:"containerPort"`
    Protocol string `json:"protocol"`
    HostIP string `json:"hostIP"`
}

type EnvCSTT struct {
    Name string `json:"name"`
    Value string `json:"value"`
    ValueFrom ValueFromECSTT `json:"valueFrom"`
}

type ValueFromECSTT struct {
    FieldRef FieldRefVFECSTT `json:"fieldRef"` 
    ConfigMapKeyRef ConfigMapKeyRefVFECSTT `json:"configMapKeyRef"`
    SecretKeyRef SecretKeyRefVFECSTT `json:"secretKeyRef"`
}

type FieldRefVFECSTT struct {
    ApiVersion string `json:"apiVersion"`
    FieldPath string `json:"fieldPath"`
}

type ConfigMapKeyRefVFECSTT struct {
    Name string `json:"name"`
    Key string `json:"key"`
}

type SecretKeyRefVFECSTT struct {
    Name string `json:"name"`
    Key string `json:"key"`
}

type ResourceCSTT struct {
    Limits map[string] string `json:"limits"`
    Requests map[string] string `json:"requests"`
}

type VolumeMountsCSTT struct {
    Name string `json:"name"`
    ReadOnly bool `json:"readOnly"`
    MountPath string `json:"mountPath"`
}

type LivenessProbeCSTT struct {
    Exec ExecLPCSTT `json:"exec"`
    HttpGet HttpGetLPCSTT `json:"httpGet"`
    TcpSocket TcpSocketLPCSTT `json:"tcpSocket"`
    InitialDelaySeConds float64 `json:"initialDelaySeconds"`
    TimeoutSeconds float64 `json:"timeoutSeconds"`
    PeriodSeconds float64 `json:"periodSeconds"`
    SuccessThreshold float64 `json:"successThreshold"`
    FailureThreshold float64 `json:"failureThreshold"`
}

type ExecLPCSTT struct {
    Command []string `json:"command"`
}

type HttpGetLPCSTT struct {
    Path string `json:"path"`
    Port string `json:"port"`
    Host string `json:"host"`
    Scheme string `json:"scheme"`
    HttpHeaders []HttpHeadersHGLPCSTT `json:"httpHeaders"`
}

type HttpHeadersHGLPCSTT struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

type TcpSocketLPCSTT struct {
    Port string `json:"port"`
}

type ReadinessProbeCSTT struct {
    Exec ExecRPCSTT `json:"exec"`
    HttpGet HttpGetRPCSTT `json:"httpGet"`
    TcpSocket TcpSocketRPCSTT `json:"tcpSocket"`
    InitialDelaySeConds float64 `json:"initialDelaySeconds"`
    TimeoutSeconds float64 `json:"timeoutSeconds"`
    PeriodSeconds float64 `json:"periodSeconds"`
    SuccessThreshold float64 `json:"successThreshold"`
    FailureThreshold float64 `json:"failureThreshold"`
   
} 

type ExecRPCSTT struct {
    Command []string `json:"command"`
}

type HttpGetRPCSTT struct {
    Path string `json:"path"`
    Port string `json:"port"`
    Host string `json:"host"`
    Scheme string `json:"scheme"`
    HttpHeaders []HttpHeadersHGRPCSTT `json:"httpHeaders"`
}

type HttpHeadersHGRPCSTT struct {
    Name string `json:"name"`
    Value string `json:"value"`
}

type TcpSocketRPCSTT struct {
    Port string `json:"port"`
}

type LifeCycleCSTT struct {
    PostStart PostStartLCCSTT `json:"postStart"`
    PreStop PreStopLCCSTT `json:"preStop"`
}


type PostStartLCCSTT struct {
    Exec ExecPostLCCSTT `json:"exec"`       
}

type ExecPostLCCSTT struct {
    Command []string `json:"command"`
}

type PreStopLCCSTT struct {
    Exec ExecPreLCCSTT `json:"exec"`       
}

type ExecPreLCCSTT struct {
    Command []string `json:"command"`
}


type StatusType struct {
    Replicas float64 `json:"replicas"`
    FullyLabeledReplicas float64 `json:"fullyLabeledReplicas"`
    ObservedGeneration float64 `json:"observedGeneration"`    
}
