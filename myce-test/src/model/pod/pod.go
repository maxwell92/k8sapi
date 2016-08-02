package pod 

// AppInfoType is only used for testing.
type AppInfoType struct {
    Healthz AppHealthzType `json: "appHealthz"`
    Name string `json: "appName"`
    Label map[string] string `json: "appLabel"`
    Datacenter []string `json: "appDatacenter"`
    Replicas float64 `json: "appReplicas"` 
    Worktime string `json: "appWorktime"`
}

type AppHealthzType struct {
    PodsAvailable string `json: "podsAvailable"`
}

type AppListType []AppInfoType


// Podlist for Get @/api/v1/namespaces/{namespaces}/pods
type PodList struct {
	Kind string `json: "kind"`
	ApiVersion string `json: "apiVersion"`
	Metadata MetadataPLType `json: "metadata"`
	Items []ItemsPLType	`json: "items"`
}

type MetadataPLType struct {
	ResourceVersion string `json: "resourceVersion"`
}


// Pod for Post @/api/v1/namespaces/{namespaces}/pods
type ItemsPLType Pod

type ItemsPLType struct {
	Kind string `json: "kind"`
	ApiVersion string `json: "apiVersion"`
	Metadata MetadataIType `json: "metadata"`
	Spec SpecType `json: "spec"`
	Status StatusType `json: "status"`	
}

type MetadataIType struct {
	Name string `json: "name"`
	GenerateName string `json: "generateName"`
	Namespace string `json: "namespace"`
	SelfLink string `json: "selfLink"`
	Uid string `json: "uid"`
	ResourceVersion string `json: "resourceVersion"`
	//Generation float64 `json: "generation"` //float64 string
	//Generation string `json: "generation"` //float64 string
	CreationTimeStamp string `json: "creationTimeStamp"`
	DeletionTimeStamp string `json: "deletionTimeStamp"`
	Labels map[string] string `json: "labels"`
	Annotations map[string] string `json: "annotations"`
}

type SpecType struct {
	Volumes []VolumesSType `json: "volumes"`
	Containers []ContainerSType `json: "containers"`
	RestartPolicy string `json: "restartPolicy"`
	TerminationGracePeriodSeconds float64 `json: "terminationGracePeriodSeconds"`
	ActiveDeadlineSeconds float64 `json: "activeDeadlineSeconds"`
	DnsPolicy string `json: "dnsPolicy"`
	NodeSelector map[string] string  `json: "nodeSelector"`
	NodeName string `json: "nodeName"`
	HostNetwork bool `json: "hostNetwork"`
	HostPID bool `json: "hostPID"`
	HostIPC bool `json: "hostIPC"`
	ImagePullSecrets []ImagePullSecretsType `json: "imagePullSecrets"`

}

type VolumesSType struct {
	Name string `json: "name"`
	HostPath HostPathVSType `json: "hostPath"`
	EmptyDir EmptyDirVSType `json: "emptyDir"`
	PersistentVolumeClaim PvClaimVSType `json: "persistentVolumeClaim"`
	Rbd RbdVSType `json: "rbd"`
	ConfigMap ConfigMapVSType `json: "configMap"`
}

type HostPathVSType struct {
	Path string `json: "path"`
}

type EmptyDirVSType struct {
	Medium string `json: "medium"`
}

type PvClaimVSType struct {
	ClaimName string `json: "claimName"`
	ReadOnly string `json: "readOnly"`
}

type RbdVSType struct {
	Monitors []string `json: "monitors"`
	Image string `json: "image"`
	FsType string `json: "fsType"`
	Pool string `json: "pool"`
	User string `json: "user"`
	Keyring string `json: "keyring"`
	SecretRef SecretRefrbdType `json: "secretRef"`
	ReadOnly bool `json: "readOnly"`
}

type SecretRefrbdType struct {
	Name string `json: "name"`
}

type ConfigMapVSType struct {
	Name string `json: "name"`
	Items []ItemsConfigMapType `json: "items"`
}

type ItemsConfigMapType struct {
	Key string `json: "key"`
	Path string `json: "path"`
}

type ContainerSType struct {
	Name string `json: "name"`
	Image string `json: "image"`
	Command []string `json: "command"`
	Args []string `json: "args"`
	WorkingDir string `json: "workingDir"`
	Ports []PortsContainerType `json: "ports"`
	Env []EnvContainerType `json: "env"`
	Resources ResourcesContainerType `json: "resources"`
	VolumeMounts []VolumeMountsContainerType `json: "volumeMounts"`
	LivenessProbe LivenessProbeContainerType `json: "livenessProbe"`
	ReadinessProbe ReadinessProbeContainerType `json: "readinessProbe"`
	Lifecycle LifecycleContainerType `json: "lifecycle"`
	TerminationMesasgePath string `json: "terminationMessagePath"`
	ImagePullPolicy string `json: "imagePullPolicy"`
	Stdin bool `json: "stdin"`
	StdinOnce bool `json: "stdinOnce"`
	Tty bool `json: "tty"`
}

type PortsContainerType struct {
	Name string `json: "name"`
	HostPort HostPortPCType `json: "hostPort"`
	ContainerPort float64 `json: "containerPort"`
	Protocol string `json: "protocol"`
	HostIP string `json: "hostIP"`
}

type HostPortPCType struct {

}

type EnvContainerType struct {
	Name string `json: "name"`
	Value string `json: "value"`
	ValueFrom ValueFromEnvContainer `json: "valueFrom"`
}

type valueFromEnvContainer struct {
	FieldRef FieldRefValueFromEnvCon `json: "fieldRef"`
	ConfigMapKeyRef ConfigMapKeyEnvCon `json: "configMapKeyRef"`
	SecretKeyRef SecretKeyRefEnvCon  `json: "secretKeyRef"`
}

type FieldRefValueFromEnvCon struct {
	ApiVersion string `json: "apiVersion"`
	FieldPath string `json: "fieldPath"`
}

type ConfigMapKeyEnvCon struct {
	Name string `json: "name"`
	Key string `json: "key"`
}

type SecretKeyRefEnvCon struct {
	Name string `json: "name"`
	Key string `json: "key"`
}

type ResourcesContainerType struct {
    LimitsRsConType map[string] string `json: "limits"`
    Requests map[string] string  `json: "requests"`
}

type VolumeMountsContainerType struct {
    Name string `json: "name"`
    ReadOnly bool `json: "readOnly"`
    MountPath string `json: "mountPath"`
}

type LivenessProbeContainerType struct {
    Exec ExecLiveProbeType `json: "exec"`
    HttpGet HttpGetLiveProbeType `json: "httpGet"`
    TcpSocket TcpLiveProbeType `json: "tcpSocket"`
    InitialDelaySeconds float64 `json: "initialDelaySeconds"`
    TimeoutSeconds float64 `json: "timeoutSeconds"`
    PeriodSeconds float64 `json: "periodSeconds"`
    SuccessThreshold float64 `json: "successThreshold"`
    FailureThreshold float64 `json: "failureThreshold"`
}

type ExecLiveProbeType struct {
    Command []string `json: "command"` 
}

type HttpGetLiveProbeType struct {
    Path string `json: "path"`
    Port float64 `json: "port"`
    Host string `json: "host"`
    Scheme string `json: "scheme"`
    HttpHeaders []HeadersGLPType `json: "httpHeaders"`    
}

type HeadersGLPType struct {
    Name string `json: "name"`
    Value string `json: "value"`
}

type TcpLiveProbeType struct {
    Port float64 `json: "port"`
}


type ReadinessProbeContainerType struct {
    Exec ExecReadProbeType `json: "exec"`
    HttpGet HttpGetReadProbeType `json: "httpGet"`
    TcpSocket TcpReadProbeType `json: "tcpSocket"`
    InitialDelaySeconds float64 `json: "initialDelaySeconds"`
    TimeoutSeconds float64 `json: "timeoutSeconds"`
    PeriodSeconds float64 `json: "periodSeconds"`
    SuccessThreshold float64 `json: "successThreshold"`
    FailureThreshold float64 `json: "failureThreshold"`

}

type ExecReadProbeType struct {
    Command []string `json: "command"` 
}

type HttpGetReadProbeType struct {
    Path string `json: "path"`
    Port float64 `json: "port"`
    Host string `json: "host"`
    Scheme string `json: "scheme"`
    HttpHeaders []HeadersGRPType `json: "httpHeaders"`    
}

type HeadersGRPType struct {
    Name string `json: "name"`
    Value string `json: "value"`
}

type TcpReadProbeType struct {
    Port float64 `json: "port"`
}



type LifecycleContainerType struct {
    PostStart PostStartLCType `json: "postStart"`
    PreStop PreStopLCType `json: "preStop"`    
}

type PostStartLCType struct {
    Exec ExecPSLCType `json: "exec"`
    HttpGet HttpGetPSLCType `json: "httpGet"`
    TcpSocket TcpPSLCType `json: "tcpPSLCType"`
}

type ExecPSLCType struct {
    Command []string `json: "command"`
}

type HttpGetPSLCType struct {
    Path string `json: "path"`
    Port float64 `json: "port"`
    Host string `json: "host"`
    Scheme string `json: "scheme"`
    HttpHeaders []HeadersPSLCType `json: "httpHeaders"`    
}

type HeadersPSLCType struct {
    Name string `json: "name"`
    Value string `json: "value"`
}

type TcpPSLCType struct {
    Port float64 `json: "port"`
}



type PreStopLCType struct {
    Exec ExecPrSLCType `json: "exec"`
    HttpGet HttpGetPrSLCType `json: "httpGet"`
    TcpSocket TcpPrSLCType `json: "tcpPSLCType"`
}

type ExecPrSLCType struct {
    Command []string `json: "command"`
}

type HttpGetPrSLCType struct {
    Path string `json: "path"`
    Port float64 `json: "port"`
    Host string `json: "host"`
    Scheme string `json: "scheme"`
    HttpHeaders []HeadersPrSLCType `json: "httpHeaders"`    
}

type HeadersPrSLCType struct {
    Name string `json: "name"`
    Value string `json: "value"`
}

type TcpPrSLCType struct {
    Port float64 `json: "port"`
}

type ImagePullSecretsType struct {
    Name string `json: "name"`
}

type StatusType struct {
	Phase string `json: "phase"`
	Conditions []ConditionsStatType `json: "conditions"`
	Message string `json: "message"`
	Reason string `json: "reason"`
	HostIP string `json: "hostIP"`
	StartTime string `json: "startTime"`
	ContainerStatuses []ContainerStatuesStatType `json: "containerStatuses"`
}

type ConditionsStatType struct {
    Type string `json: "type"`
    Status string `json: "status"`
    LastProbeTime string `json: "lastProbeTime"`
    LastTransition string `json: "lastTransition"`
    Reason string `json: "reason"`
    Message string `json: "mesage"`
} 

type ContainerStatuesStatType struct {
    Name string `json: "name"`
    State StatCSSType `json: "state"`
    LastState LastStateCSSType `json: "lastState"`
    Ready bool `json: "ready"`
    RestartCount float64 `json: "restartCount"`
    Image string `json: "image"`
    ImageID string `json: "imageID"`
    ContainerID string `json: "containerID"`
}

type StatCSSType struct {
    Waiting WaitSCSSType `json: "waiting"`
    Running RunSCSSType `json: "running"`
    Terminated TermCSSType `json: "terminated"`
}

type WaitSCSSType struct {
    Reason string `json: "reason"`
    Message string `json: "message"` 
}

type RunSCSSType struct {
    StartedAt string `json: "startedAt"`
}

type TermCSSType struct {
    ExitCode float64 `json: "exitCode"`
    Signal float64 `json: "signal"`
    Reason string `json: "reason"`
    Message string `json: "message"`
    StartedAt string `json: "startedAt"`
    FinishedAt string `json: "finishedAt"`
    ContainerID string `json: "containerID"`
}

type LastStateCSSType struct {
    Waiting WaitLSCSSType `json: "waiting"`
    Running RunLSCSSType `json: "running"`
    Terminated TermLSCSSType `json: "terminated"`
}

type WaitLSCSSType struct {
    Reason string `json: "reason"`
    Message string `json: "message"`
}

type RunLSCSSType struct {
    StartedAt string `json: "startedAt"`
}

type TermLSCSSType struct {
    ExitCode float64 `json: "exitCode"`
    Signal float64 `json: "signal"`
    Reason string `json: "reason"`
    Message string `json: "message"`
    StartedAt string `json: "startedAt"`
    FinishedAt string `json: "finishedAt"`
    ContainerID string `json: "containerID"`
}




