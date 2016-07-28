package main

type deploymentList struct {
    Kind string `json: "name"`
    ApiVersion string `json: "apiVersion"`
    Metadata metadataDL `json: "metadata"`
    Spec specDL `json: "spec"`
    Status statusDL `json: "status"`
}

type metadataDL struct {
    Name string `json: "name"`
    GenerateName string `json: "generatename"`
    Namespace string `json: "namespace"`
    SelfLink string `json: "selfLink"`
    Uid string `json: "uid"`
    ResourceVersion string `json: "resourceVersion"`
    CreationTimestamp string `json: "creationTimestamp"`
    DeletionTimestamp string `json: "deletionTimestamp"`
    DeletionGracePeriodSecons float64 `json: "deletionGracePeriodSeconds"`
    Labels map[string] string `json: "labels"`
    Annotations map[string] string `json: "annotations"`
}

type specDL struct {
    Replicas float64 `json: "replicas"`
    Selector selectorSpec `json: "selector"` 
    Template templateSpec `json: "template"`
    Strategy strategySpec `json: "strategy"`
    MinReadySeconds float64 `json: "minReadySeconds"`
    RevisionHistoryLimit float64 `json: "revisionHistoryLimit"`
    Paused bool `json: "paused"`
    RollbackTo rollbackToSpec `json: "rollbackTo"`
}

type selectorSpec struct {
    MatchLabels string `json: "matchLabels"`
    MatchExpressions []matchESSDL `json: "matchExpressions"`
}

type matchESSDL struct {
    Key string `json: "key"`
    Operator string `json: "operator"`
    Values []string `json: "values"`
}

type templateSpec struct {
    Metadata metadataTSDL `json: "metadata"`
    Spec specTSDL `json: "spec"`
}

type metadataTSDL struct {
    Name string `json: "name"`
    GenerateName string `json: "generateName"`
    Namespace string `json: "namespace"`
    SelfLink string `json: "selfLink"`
    Uid string `json: "uid"`
    ResourceVersion string `json: "resourceVersion"`
    CreationTimestamp string `json: "creationTimestamp"`
    DeletionTimestamp string `json: "deletionTimestamp"`
    DeletionGracePeriodSecons float64 `json: "deletionGracePeriodSeconds"`
    Labels map[string] string `json: "labels"`
    Annotations map[string] string `json: "annotations"`    
}

type specTSDL struct {
    Volumes []volumesSTSDL `json: "volumes"`
    Containers []containersSTSDL `json: "containers"` 
    RestartPolicy string `json: "restartPolicy"`
    TerminationGracePeriodSeconds float64 `json: "terminationMessagePath"`
    ActiveDeadlineSeconds float64 `json: "activeDeadlineSeconds"`
    DnsPolicy string `json: "dnsPolicy"`
    NodeSelector string `json: "nodeSelector"`
    NodeName string `json: "nodeName"`
    HostNetwork bool `json: "hostNetwork"`
    HostPID bool `json: "hostPID"`
    HostIPC bool `json: "hostIPC"`
    ImagePullSecrets []ImagePullSecretsSpec `json: "imagePullSecrets"`
}

type ImagePullSecretsSpec struct {
    Name string `json: "name"`
}

type volumesSTSDL struct {
    Name string `json: "name"`
    HostPath hostPathVSTSDL `json: "hostPath"`
    EmptyDir emptyDirVSTSDL `json: "emptyDir"`
    PersistentVolumeClaim pvVolumeClaimVSTSDL `json: "persistentVolumeClaim"`
    Rbd rbdVSTSDL `json: "rbd"`
    ConfigMap configMapVSTSDL `json: "configMap"`
}

type hostPathVSTSDL struct {
    Path string `json: "path"
}

type emptyDirVSTSDL struct {
    Medium string `json: "medium"`
}

type pvVolumeClaimVSTSDL struct {
    ClaimName string `json: "claimName"`
    ReadOnly bool `json: "readOnly"`
}

type rbdVSTSDL struct {
    Monitors []string `json: "monitors"`
    Image string `json: "image"`
    FsType string `json: "fsType"`
    Pool string `json: "pool"`
    User string `json: "user"`
    Keyring string `json: "keyring"`
    SecretRef secretRefRVSTSDL `json: "secretRef"`
    ReadOnly bool `json: "readOnly"` 
}

type secretRefRVSTSDL struct {
    Name string `json: name"`
}

type configMapVSTSDL struct {
    Name string `json: "name"` 
    Items []itemsConfigMap `json: "items"`
}

type itemsConfigMap struct {
    Key string `json: "key"`
    Path string `json: "Path"`
}

type containersSTSDL struct {
    Name string `json: "name"`
    Image string `json: "image"`
    Command []string `json: "command"`
    Args []string `json: "args"`
    WorkingDir string `json: "workingDir"`
    Ports []portContainer `json: "ports"`     
    Env []envContainer `json: "env"`
    Resources resourcesContainer `json: "resources"`
    VolumeMounts []volumeMountsContainer `json: "volumeMounts"` 
    LivenessProbe livenessProbeContainer `json: "livenessProbe"`
    ReadinessProbe readinessProbeContainer `json: "readinessProbeContainer"`
    Lifecycle lifecycleContainer `json: "lifecycle"`
    TerminationMessagePath string `json: "terminationMessagePath"`
    ImagePullPolicy string `json: "imagePullPolicy"`
    Stdin bool `json: "stdin"`
    stdinOnce bool `json: "stdinOnce"`
    Tty bool `json: "tty"`
}

type portContainer struct {
    Name string `json: "name"`
    HostPort float64 `json: "hostPort"`
    ContainerPort float64 `json: "containerPort"`
    Protocol string `json: "protocol"`
    HostIP string `json: "hostIP"` 
}

type envContainer struct {
    Name string `json: "name"`
    Value string `json: "value"`
    ValueFrom valueFromEnv `json: "valueFrom"`
}

type valueFromEnv struct {
    FieldRef fieldRefValue `json: "fieldRef"`
    ConfigMapKeyRef configMapKeyRefValue `json: "configMapKeyRef"`
    SecretKeyRef secretKeyRefValue `json: "secretKeyRef"`
}

type fieldRefValue struct {
    ApiVersion string `json: "apiVersion"`
    FieldPath string `json: fieldPath"`
}

type configMapKeyRefValue struct {
    Name string `json: "name"`
    Key string `json: "key"` 
}

type secretKeyRefValue struct {
    Name string `json: "name"`
    Key string `json: "key"`
}

type resourcesContainer struct {
    Limits map[string] string `json: "limits"`
    Requests map[string] string `json: "requests"`
}
    
type volumeMountsContainer struct {
    Name string `json: "name"
    ReadOnly bool `json: "readOnly"`
    MountPath string `json: "mountPath"` 
}

type livenessProbeContainer struct {
    Exec execLiveProbeType `json: "exec"`
    HttpGet httpGetLiveProbeType `json: "httpGet"`
    TcpSocket tcpLiveProbeType `json: "tcpSocket"`
    InitialDelaySeconds float64 `json: "initialDelaySeconds"`
    TimeoutSeconds float64 `json: "timeoutSeconds"`
    PeriodSeconds float64 `json: "periodSeconds"`
    SuccessThreshold float64 `json: "successThreshold"`
    FailureThreshold float64 `json: "failureThreshold"`
}

type execLiveProbeType struct {
    Command []string `json: "command"` 
}

type httpGetLiveProbeType struct {
    Path string `json: "path"`
    Port float64 `json: "port"`
    Host string `json: "host"`
    Scheme string `json: "scheme"`
    HttpHeaders []headersGLPType `json: "httpHeaders"`    
}

type headersGLPType struct {
    Name string `json: "name"`
    Value string `json: "value"`
}

type tcpLiveProbeType struct {
    Port float64 `json: "port"`
}

type readinessProbeContainer struct {
    Exec execReadProbeType `json: "exec"`
    HttpGet httpGetReadProbeType `json: "httpGet"`
    TcpSocket tcpReadProbeType `json: "tcpSocket"`
    InitialDelaySeconds float64 `json: "initialDelaySeconds"`
    TimeoutSeconds float64 `json: "timeoutSeconds"`
    PeriodSeconds float64 `json: "periodSeconds"`
    SuccessThreshold float64 `json: "successThreshold"`
    FailureThreshold float64 `json: "failureThreshold"`
}

type execReadProbeType struct {
    Command []string `json: "command"` 
}

type httpGetReadProbeType struct {
    Path string `json: "path"`
    Port float64 `json: "port"`
    Host string `json: "host"`
    Scheme string `json: "scheme"`
    HttpHeaders []headersGRPType `json: "httpHeaders"`    
}

type headersGRPType struct {
    Name string `json: "name"`
    Value string `json: "value"`
}

type tcpReadProbeType struct {
    Port float64 `json: "port"`
}

type lifecycleContainerType struct {
    PostStart postStartLCType `json: "postStart"`
    PreStop preStopLCType `json: "preStop"`    
}

type postStartLCType struct {
    Exec execPSLCType `json: "exec"`
    HttpGet httpGetPSLCType `json: "httpGet"`
    TcpSocket tcpPSLCType `json: "tcpPSLCType"`
}

type execPSLCType struct {
    Command []string `json: "command"`
}

type httpGetPSLCType struct {
    Path string `json: "path"`
    Port float64 `json: "port"`
    Host string `json: "host"`
    Scheme string `json: "scheme"`
    HttpHeaders []headersPSLCType `json: "httpHeaders"`    
}

type headersPSLCType struct {
    Name string `json: "name"`
    Value string `json: "value"`
}

type tcpPSLCType struct {
    Port float64 `json: "port"`
}



type preStopLCType struct {
    Exec execPrSLCType `json: "exec"`
    HttpGet httpGetPrSLCType `json: "httpGet"`
    TcpSocket tcpPrSLCType `json: "tcpPSLCType"`
}

type execPrSLCType struct {
    Command []string `json: "command"`
}

type httpGetPrSLCType struct {
    Path string `json: "path"`
    Port float64 `json: "port"`
    Host string `json: "host"`
    Scheme string `json: "scheme"`
    HttpHeaders []headersPrSLCType `json: "httpHeaders"`    
}

type headersPrSLCType struct {
    Name string `json: "name"`
    Value string `json: "value"`
}

type tcpPrSLCType struct {
    Port float64 `json: "port"`
}




type strategySpec struct {
    Type string `json: "type"`
    RollingUpdate rollingUpdateStrategy `json: "rollilngUpdate"`
}

type rollingUpdateStrategy struct {
    MaxUnavailable string `json: "maxUnavailable"`
    MaxSurge string `json: "maxSurge"`
}



type rollbackToSpec struct {
    Revision float64 `json: "revision"`
}

type statusDL struct {
    ObservedGeneration float64 `json: "observedGeneration"`    
    Replicas float64 `json: "replicas"`
    UpdateReplicas float64 `json: "updateReplicas"`
    AvailableReplicas float64 `json: "availableReplicas"`
    UnavailableReplicas float64 `json: "unavailableReplicas"`
}
