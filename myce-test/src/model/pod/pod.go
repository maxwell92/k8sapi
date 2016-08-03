package pod

// AppInfoType is only used for testing.
type AppInfoType struct {
	Healthz    AppHealthzType    `json:"appHealthz"`
	Name       string            `json:"appName"`
	Label      map[string]string `json:"appLabel"`
	Datacenter []string          `json:"appDatacenter"`
	Replicas   float64           `json:"appReplicas"`
	Worktime   string            `json:"appWorktime"`
}

type AppHealthzType struct {
	PodsAvailable string `json:"podsAvailable"`
}

type AppListType []AppInfoType

// Podlist for Get @/api/v1/namespaces/{namespaces}/pods
type PodList struct {
	Kind       string         `json:"kind"`
	ApiVersion string         `json:"apiVersion"`
	Metadata   MetadataPLType `json:"metadata"`
	Items      []ItemsPLType  `json:"items"`
}

type MetadataPLType struct {
	ResourceVersion string `json:"resourceVersion,omitempty"`
}

// Pod for Post @/api/v1/namespaces/{namespaces}/pods

type ItemsPLType struct {
	Kind       string        `json:"kind,omitempty"`
	ApiVersion string        `json:"apiVersion,omitempty"`
	Metadata   MetadataIType `json:"metadata"`
	Spec       SpecType      `json:"spec"`
	Status     *StatusType   `json:"status,omitempty"`
}

type Pod ItemsPLType

type MetadataIType struct {
	Name            string `json:"name"`
	GenerateName    string `json:"generateName,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
	SelfLink        string `json:"selfLink,omitempty"`
	Uid             string `json:"uid,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
	//Generation float64 `json:"generation"` //float64 string
	//Generation string `json:"generation"` //float64 string
	CreationTimeStamp string            `json:"creationTimeStamp,omitempty"`
	DeletionTimeStamp string            `json:"deletionTimeStamp,omitempty"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations,omitempty"`
}

type SpecType struct {
	Volumes                       []VolumesSType         `json:"volumes,omitempty"`
	Containers                    []ContainerSType       `json:"containers"`
	RestartPolicy                 string                 `json:"restartPolicy,omitempty"`
	TerminationGracePeriodSeconds float64                `json:"terminationGracePeriodSeconds,omitempty"`
	ActiveDeadlineSeconds         float64                `json:"activeDeadlineSeconds,omitempty"`
	DnsPolicy                     string                 `json:"dnsPolicy,omitempty"`
	NodeSelector                  map[string]string      `json:"nodeSelector,omitempty"`
	NodeName                      string                 `json:"nodeName,omitempty"`
	HostNetwork                   bool                   `json:"hostNetwork,omitempty"`
	HostPID                       bool                   `json:"hostPID,omitempty"`
	HostIPC                       bool                   `json:"hostIPC,omitempty"`
	ImagePullSecrets              []ImagePullSecretsType `json:"imagePullSecrets,omitempty"`
}

type VolumesSType struct {
	Name                  string           `json:"name"`
	HostPath              *HostPathVSType  `json:"hostPath",omitempty`
	EmptyDir              *EmptyDirVSType  `json:"emptyDir",omitempty`
	PersistentVolumeClaim *PvClaimVSType   `json:"persistentVolumeClaim,omitempty"`
	Rbd                   *RbdVSType       `json:"rbd,omitempty"`
	ConfigMap             *ConfigMapVSType `json:"configMap,omitempty"`
}

type HostPathVSType struct {
	Path string `json:"path,omitempty"`
}

type EmptyDirVSType struct {
	Medium string `json:"medium,omitempty"`
}

type PvClaimVSType struct {
	ClaimName string `json:"claimName,omitempty"`
	ReadOnly  string `json:"readOnly,omitempty"`
}

type RbdVSType struct {
	Monitors  []string          `json:"monitors,omitempty"`
	Image     string            `json:"image,omitempty"`
	FsType    string            `json:"fsType,omitempty"`
	Pool      string            `json:"pool,omitempty"`
	User      string            `json:"user,omitempty"`
	Keyring   string            `json:"keyring,omitempty"`
	SecretRef *SecretRefrbdType `json:"secretRef,omitempty"`
	ReadOnly  bool              `json:"readOnly,omitempty"`
}

type SecretRefrbdType struct {
	Name string `json:"name,omitempty"`
}

type ConfigMapVSType struct {
	Name  string               `json:"name,omitempty"`
	Items []ItemsConfigMapType `json:"items,omitempty"`
}

type ItemsConfigMapType struct {
	Key  string `json:"key",omitempty`
	Path string `json:"path,omitempty"`
}

type ContainerSType struct {
	Name                   string                       `json:"name"`
	Image                  string                       `json:"image"`
	Command                []string                     `json:"command,omitempty"`
	Args                   []string                     `json:"args,omitempty"`
	WorkingDir             string                       `json:"workingDir,omitempty"`
	Ports                  []PortsContainerType         `json:"ports,omitempty"`
	Env                    []EnvContainerType           `json:"env,omitempty"`
	Resources              *ResourcesContainerType      `json:"resources,omitempty"`
	VolumeMounts           []VolumeMountsContainerType  `json:"volumeMounts,omitempty"`
	LivenessProbe          *LivenessProbeContainerType  `json:"livenessProbe,omitempty"`
	ReadinessProbe         *ReadinessProbeContainerType `json:"readinessProbe,omitempty"`
	Lifecycle              *LifecycleContainerType      `json:"lifecycle,omitempty"`
	TerminationMesasgePath string                       `json:"terminationMessagePath,omitempty"`
	ImagePullPolicy        string                       `json:"imagePullPolicy,omitempty"`
	Stdin                  bool                         `json:"stdin,omitempty"`
	StdinOnce              bool                         `json:"stdinOnce,omitempty"`
	Tty                    bool                         `json:"tty,omitempty"`
}

type PortsContainerType struct {
	Name          string          `json:"name,omitempty"`
	HostPort      *HostPortPCType `json:"hostPort,omitempty"`
	ContainerPort float64         `json:"containerPort,omitempty"`
	Protocol      string          `json:"protocol,omitempty"`
	HostIP        string          `json:"hostIP,omitempty"`
}

type HostPortPCType struct {
}

type EnvContainerType struct {
	Name      string                 `json:"name",omitempty`
	Value     string                 `json:"value",omitempty`
	ValueFrom *ValueFromEnvContainer `json:"valueFrom",omitempty`
}

type ValueFromEnvContainer struct {
	FieldRef        *FieldRefValueFromEnvCon `json:"fieldRef,omitempty"`
	ConfigMapKeyRef *ConfigMapKeyEnvCon      `json:"configMapKeyRef,omitempty"`
	SecretKeyRef    *SecretKeyRefEnvCon      `json:"secretKeyRef,omitempty"`
}

type FieldRefValueFromEnvCon struct {
	ApiVersion string `json:"apiVersion,omitempty"`
	FieldPath  string `json:"fieldPath,omitempty"`
}

type ConfigMapKeyEnvCon struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
}

type SecretKeyRefEnvCon struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
}

type ResourcesContainerType struct {
	LimitsRsConType map[string]string `json:"limits,omitempty"`
	Requests        map[string]string `json:"requests,omitempty"`
}

type VolumeMountsContainerType struct {
	Name      string `json:"name,omitempty"`
	ReadOnly  bool   `json:"readOnly,omitempty"`
	MountPath string `json:"mountPath,omitempty"`
}

type LivenessProbeContainerType struct {
	Exec                *ExecLiveProbeType    `json:"exec,omitempty"`
	HttpGet             *HttpGetLiveProbeType `json:"httpGet,omitempty"`
	TcpSocket           *TcpLiveProbeType     `json:"tcpSocket,omitempty"`
	InitialDelaySeconds float64               `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds      float64               `json:"timeoutSeconds,omitempty"`
	PeriodSeconds       float64               `json:"periodSeconds,omitempty"`
	SuccessThreshold    float64               `json:"successThreshold,omitempty"`
	FailureThreshold    float64               `json:"failureThreshold,omitempty"`
}

type ExecLiveProbeType struct {
	Command []string `json:"command,omitempty"`
}

type HttpGetLiveProbeType struct {
	Path        string           `json:"path,omitempty"`
	Port        float64          `json:"port,omitempty"`
	Host        string           `json:"host,omitempty"`
	Scheme      string           `json:"scheme,omitempty"`
	HttpHeaders []HeadersGLPType `json:"httpHeaders,omitempty"`
}

type HeadersGLPType struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type TcpLiveProbeType struct {
	Port float64 `json:"port,omitempty"`
}

type ReadinessProbeContainerType struct {
	Exec                *ExecReadProbeType    `json:"exec,omitempty"`
	HttpGet             *HttpGetReadProbeType `json:"httpGet,omitempty"`
	TcpSocket           *TcpReadProbeType     `json:"tcpSocket,omitempty"`
	InitialDelaySeconds float64               `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds      float64               `json:"timeoutSeconds,omitempty"`
	PeriodSeconds       float64               `json:"periodSeconds,omitempty"`
	SuccessThreshold    float64               `json:"successThreshold,omitempty"`
	FailureThreshold    float64               `json:"failureThreshold,omitempty"`
}

type ExecReadProbeType struct {
	Command []string `json:"command,omitempty"`
}

type HttpGetReadProbeType struct {
	Path        string           `json:"path,omitempty"`
	Port        float64          `json:"port,omitempty"`
	Host        string           `json:"host,omitempty"`
	Scheme      string           `json:"scheme,omitempty"`
	HttpHeaders []HeadersGRPType `json:"httpHeaders,omitempty"`
}

type HeadersGRPType struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type TcpReadProbeType struct {
	Port float64 `json:"port,omitempty"`
}

type LifecycleContainerType struct {
	PostStart *PostStartLCType `json:"postStart,omitempty"`
	PreStop   *PreStopLCType   `json:"preStop,omitempty"`
}

type PostStartLCType struct {
	Exec      *ExecPSLCType    `json:"exec,omitempty"`
	HttpGet   *HttpGetPSLCType `json:"httpGet,omitempty"`
	TcpSocket *TcpPSLCType     `json:"tcpPSLCType,omitempty"`
}

type ExecPSLCType struct {
	Command []string `json:"command,omitempty"`
}

type HttpGetPSLCType struct {
	Path        string            `json:"path,omitempty"`
	Port        float64           `json:"port,omitempty"`
	Host        string            `json:"host,omitempty"`
	Scheme      string            `json:"scheme,omitempty"`
	HttpHeaders []HeadersPSLCType `json:"httpHeaders,omitempty"`
}

type HeadersPSLCType struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type TcpPSLCType struct {
	Port float64 `json:"port,omitempty"`
}

type PreStopLCType struct {
	Exec      *ExecPrSLCType    `json:"exec,omitempty"`
	HttpGet   *HttpGetPrSLCType `json:"httpGet,omitempty"`
	TcpSocket *TcpPrSLCType     `json:"tcpPSLCType,omitempty"`
}

type ExecPrSLCType struct {
	Command []string `json:"command,omitempty"`
}

type HttpGetPrSLCType struct {
	Path        string             `json:"path,omitempty"`
	Port        float64            `json:"port,omitempty"`
	Host        string             `json:"host,omitempty"`
	Scheme      string             `json:"scheme,omitempty"`
	HttpHeaders []HeadersPrSLCType `json:"httpHeaders,omitempty"`
}

type HeadersPrSLCType struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type TcpPrSLCType struct {
	Port float64 `json:"port,omitempty"`
}

type ImagePullSecretsType struct {
	Name string `json:"name,omitempty"`
}

type StatusType struct {
	Phase             string                     `json:"phase,omitempty"`
	Conditions        []ConditionsStatType       `json:"conditions,omitempty"`
	Message           string                     `json:"message,omitempty"`
	Reason            string                     `json:"reason,omitempty"`
	HostIP            string                     `json:"hostIP,omitempty"`
	StartTime         string                     `json:"startTime,omitempty"`
	ContainerStatuses []ContainerStatuesStatType `json:"containerStatuses,omitempty"`
}

type ConditionsStatType struct {
	Type           string `json:"type,omitempty"`
	Status         string `json:"status,omitempty"`
	LastProbeTime  string `json:"lastProbeTime,omitempty"`
	LastTransition string `json:"lastTransition,omitempty"`
	Reason         string `json:"reason,omitempty"`
	Message        string `json:"mesage,omitempty"`
}

type ContainerStatuesStatType struct {
	Name         string            `json:"name,omitempty"`
	State        *StatCSSType      `json:"state,omitempty"`
	LastState    *LastStateCSSType `json:"lastState,omitempty"`
	Ready        bool              `json:"ready,omitempty"`
	RestartCount float64           `json:"restartCount,omitempty"`
	Image        string            `json:"image,omitempty"`
	ImageID      string            `json:"imageID,omitempty"`
	ContainerID  string            `json:"containerID,omitempty"`
}

type StatCSSType struct {
	Waiting    *WaitSCSSType `json:"waiting,omitempty"`
	Running    *RunSCSSType  `json:"running,omitempty"`
	Terminated *TermCSSType  `json:"terminated,omitempty"`
}

type WaitSCSSType struct {
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

type RunSCSSType struct {
	StartedAt string `json:"startedAt,omitempty"`
}

type TermCSSType struct {
	ExitCode    float64 `json:"exitCode,omitempty"`
	Signal      float64 `json:"signal,omitempty"`
	Reason      string  `json:"reason,omitempty"`
	Message     string  `json:"message,omitempty"`
	StartedAt   string  `json:"startedAt,omitempty"`
	FinishedAt  string  `json:"finishedAt,omitempty"`
	ContainerID string  `json:"containerID,omitempty"`
}

type LastStateCSSType struct {
	Waiting    *WaitLSCSSType `json:"waiting,omitempty"`
	Running    *RunLSCSSType  `json:"running,omitempty"`
	Terminated *TermLSCSSType `json:"terminated,omitempty"`
}

type WaitLSCSSType struct {
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

type RunLSCSSType struct {
	StartedAt string `json:"startedAt,omitempty"`
}

type TermLSCSSType struct {
	ExitCode    float64 `json:"exitCode,omitempty"`
	Signal      float64 `json:"signal,omitempty"`
	Reason      string  `json:"reason,omitempty"`
	Message     string  `json:"message,omitempty"`
	StartedAt   string  `json:"startedAt,omitempty"`
	FinishedAt  string  `json:"finishedAt,omitempty"`
	ContainerID string  `json:"containerID,omitempty"`
}
