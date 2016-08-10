package replicaset

// ReplicaSetList for Get @/apis/extensions/v1beta1/namespaces/{namespace}/replicasets
type ReplicaSetList struct {
	Kind       string           `json:"kind"`
	ApiVersion string           `json:"apiVersion"`
	Metadata   MetadataListType `json:"metadata"`
	Items      []ItemsListType  `json:"items"`
}

type MetadataListType struct {
	SelfLink        string `json:"selfLink"`
	ResourceVersion string `json:"resourceVersion"`
}

// ReplicaSet for Post @/apis/extensions/v1beta1/namespaces/{namespace}/replicasets
type ItemsListType struct {
	Kind       string     `json:"kind"`
	ApiVersion string     `json:"apiVersion"`
	Metadata   MetadataIT `json:"metadata"`
	Spec       SpecType   `json:"spec"`
	Status     StatusType `json:"status,omitempty"`
}

type ReplicaSet ItemsListType

type MetadataIT struct {
	Name                      string            `json:"name"`
	GenerateName              string            `json:"generateName,omitempty"`
	Namespace                 string            `json:"namespace,omitempty"`
	SelfLink                  string            `json:"selfLink,omitempty"`
	Uid                       string            `json:"uid,omitempty"`
	ResourceVersion           string            `json:"resourceVersion,omitempty"`
	Generation                float64           `json:"generation,omitempty"`
	CreationTimestamp         string            `json:"creationTimestamp,omitempty"`
	DeletionTimestamp         string            `json:"deletionTimestamp,omitempty"`
	DeletionGracePeriodSecond float64           `json:"deletionGracePeriodSeconds,omitempty"`
	Labels                    map[string]string `json:"labels"`
	Annotations               map[string]string `json:"annotations,omitempty"`
}

type SpecType struct {
	Replicas float64       `json:"replicas"`
	Selector *SelectorType `json:"selector,omitempty"`
	Template TemplateType  `json:"template"`
}

type SelectorType struct {
	MatchLabels      map[string]string `json:"matchLabels,omitempty"`
	MatchExpressions []MatchExpType    `json:"matchExpressions,omitempty"`
}

type MatchExpType struct {
	Key      string   `json:"key,omitempty"`
	Operator string   `json:"operator,omitempty"`
	Values   []string `json:"values,omitempty"`
}

type TemplateType struct {
	Metadata MetadataTemplateType `json:"metadata"`
	Spec     SpecTemplateType     `json:"spec"`
}

type MetadataTemplateType struct {
	Name                       string            `json:"name,omitempty"`
	GenerateName               string            `json:"generateName,omitempty"`
	Namespace                  string            `json:"namespace,omitempty"`
	SelfLink                   string            `json:"selfLink,omitempty"`
	Uid                        string            `json:"uid,omitempty"`
	ResourceVersion            string            `json:"resourceVersion,omitempty"`
	Generation                 float64           `json:"generation,omitempty"`
	CreationTimestamp          string            `json:"creationTimestamp,omitempty"`
	DeletionTimestamp          string            `json:"deletionTimestamp,omitempty"`
	DeletionGracePeriodSeconds float64           `json:"deletionGracePeriodSeconds,omitempty"`
	Labels                     map[string]string `json:"labels,omitempty"`
	Annotations                map[string]string `json:"annotations,omitempty"`
}

type SpecTemplateType struct {
	Volumes                       []VolumesSTT      `json:"volumes,omitempty"`
	Containers                    []ContainerType   `json:"containers"`
	RestartPolicy                 string            `json:"restartPolicy,omitempty"`
	TerminationGracePeriodSeconds float64           `json:"terminationGracePeriodSeconds,omitempty"`
	ActiveDeadlineSeconds         float64           `json:"activeDeadlineSeconds,omitempty"`
	DnsPolicy                     string            `json:"dnsPolicy,omitempty"`
	NodeSelector                  map[string]string `json:"nodeSelector,omitempty"`
	NodeName                      string            `json:"nodeName,omitempty"`
	HostNetwork                   bool              `json:"hostNetwork,omitempty"`
	HostPID                       bool              `json:"hostPID,omitempty"`
	HostIPC                       bool              `json:"hostIPC,omitempty"`
}

type VolumesSTT struct {
	Name                  string         `json:"name,omitempty"`
	HostPath              *HostPathVSTT  `json:"hostPath,omitempty"`
	EmptyDir              *EmptyDirVSTT  `json:"emptyDir,omitempty"`
	PersistentVolumeClaim *PVClaimVSTT   `json:"persistentVolumeClaim,omitempty"`
	Rbd                   *RbdVSTT       `json:"rbd,omitempty"`
	ConfigMap             *ConfigMapVSTT `json:"configMap,omitempty"`
}

type HostPathVSTT struct {
	Path string `json:"path,omitempty"`
}

type EmptyDirVSTT struct {
	Medium string `json:"medium,omitempty"`
}

type PVClaimVSTT struct {
	ClaimName string `json:"claimName,omitempty"`
	ReadOnly  bool   `json:"readOnly,omitempty"`
}

type RbdVSTT struct {
	Monitors  []string        `json:"monitors,omitempty"`
	Image     string          `json:"image,omitempty"`
	FsType    string          `json:"fsType,omitempty"`
	Pool      string          `json:"pool,omitempty"`
	User      string          `json:"user,omitempty"`
	Keyring   string          `json:"keyring,omitempty"`
	SecretRef *SecretRefRVSTT `json:"secretRef,omitempty"`
	ReadOnly  bool            `json:"readOnly,omitempty"`
}

type SecretRefRVSTT struct {
	Name string `json:"name,omitempty"`
}

type ConfigMapVSTT struct {
	Name  string       `json:"name,omitempty"`
	Items []ItemsCVSTT `json:"items,omitempty"`
}

type ItemsCVSTT struct {
	Key  string `json:"key,omitempty"`
	Path string `json:"path,omitempty"`
}

type ContainerType struct {
	Name                   string              `json:"name,omitempty"`
	Image                  string              `json:"image"`
	Command                []string            `json:"command,omitempty"`
	Args                   []string            `json:"args,omitempty"`
	WorkingDir             string              `json:"workingDir,omitempty"`
	Ports                  []PortsCSTT         `json:"ports,omitempty"`
	Env                    []EnvCSTT           `json:"env,omitempty"`
	Resources              *ResourceCSTT       `json:"resource,omitempty"`
	VolumeMounts           []VolumeMountsCSTT  `json:"volumeMounts,omitempty"`
	LivenessProbe          *LivenessProbeCSTT  `json:"livenessProbe,omitempty"`
	ReadinessProbe         *ReadinessProbeCSTT `json:"readinessProbe,omitempty"`
	LifeCycle              *LifeCycleCSTT      `json:"lifecycle,omitempty"`
	TerminationMessagePath string              `json:"terminationMessagePath,omitempty"`
	ImagePullPolicy        string              `json:"imagePullPolicy,omitempty"`
	Stdin                  bool                `json:"stdin,omitempty"`
	StdinOnce              bool                `json:"stdinOnce,omitempty"`
	Tty                    bool                `json:"tty,omitempty"`
}

type ContainersSTT ContainerType

type PortsCSTT struct {
	Name          string  `json:"name,omitempty"`
	HostPort      float64 `json:"hostPort,omitempty"`
	ContainerPort float64 `json:"containerPort,omitempty"`
	Protocol      string  `json:"protocol,omitempty"`
	HostIP        string  `json:"hostIP,omitempty"`
}

type EnvCSTT struct {
	Name      string          `json:"name,omitempty"`
	Value     string          `json:"value,omitempty"`
	ValueFrom *ValueFromECSTT `json:"valueFrom,omitempty"`
}

type ValueFromECSTT struct {
	FieldRef        *FieldRefVFECSTT        `json:"fieldRef,omitempty"`
	ConfigMapKeyRef *ConfigMapKeyRefVFECSTT `json:"configMapKeyRef,omitempty"`
	SecretKeyRef    *SecretKeyRefVFECSTT    `json:"secretKeyRef,omitempty"`
}

type FieldRefVFECSTT struct {
	ApiVersion string `json:"apiVersion,omitempty"`
	FieldPath  string `json:"fieldPath,omitempty"`
}

type ConfigMapKeyRefVFECSTT struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
}

type SecretKeyRefVFECSTT struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
}

type ResourceCSTT struct {
	Limits   map[string]string `json:"limits,omitempty"`
	Requests map[string]string `json:"requests,omitempty"`
}

type VolumeMountsCSTT struct {
	Name      string `json:"name,omitempty"`
	ReadOnly  bool   `json:"readOnly,omitempty"`
	MountPath string `json:"mountPath,omitempty"`
}

type LivenessProbeCSTT struct {
	Exec                *ExecLPCSTT      `json:"exec,omitempty"`
	HttpGet             *HttpGetLPCSTT   `json:"httpGet,omitempty"`
	TcpSocket           *TcpSocketLPCSTT `json:"tcpSocket,omitempty"`
	InitialDelaySeConds float64          `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds      float64          `json:"timeoutSeconds,omitempty"`
	PeriodSeconds       float64          `json:"periodSeconds,omitempty"`
	SuccessThreshold    float64          `json:"successThreshold,omitempty"`
	FailureThreshold    float64          `json:"failureThreshold,omitempty"`
}

type ExecLPCSTT struct {
	Command []string `json:"command,omitempty"`
}

type HttpGetLPCSTT struct {
	Path        string                `json:"path,omitempty"`
	Port        string                `json:"port,omitempty"`
	Host        string                `json:"host,omitempty"`
	Scheme      string                `json:"scheme,omitempty"`
	HttpHeaders []HttpHeadersHGLPCSTT `json:"httpHeaders,omitempty"`
}

type HttpHeadersHGLPCSTT struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type TcpSocketLPCSTT struct {
	Port string `json:"port,omitempty"`
}

type ReadinessProbeCSTT struct {
	Exec                *ExecRPCSTT      `json:"exec,omitempty"`
	HttpGet             *HttpGetRPCSTT   `json:"httpGet,omitempty"`
	TcpSocket           *TcpSocketRPCSTT `json:"tcpSocket,omitempty"`
	InitialDelaySeConds float64          `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds      float64          `json:"timeoutSeconds,omitempty"`
	PeriodSeconds       float64          `json:"periodSeconds,omitempty"`
	SuccessThreshold    float64          `json:"successThreshold,omitempty"`
	FailureThreshold    float64          `json:"failureThreshold,omitempty"`
}

type ExecRPCSTT struct {
	Command []string `json:"command,omitempty"`
}

type HttpGetRPCSTT struct {
	Path        string                `json:"path,omitempty"`
	Port        string                `json:"port,omitempty"`
	Host        string                `json:"host,omitempty"`
	Scheme      string                `json:"scheme,omitempty"`
	HttpHeaders []HttpHeadersHGRPCSTT `json:"httpHeaders,omitempty"`
}

type HttpHeadersHGRPCSTT struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type TcpSocketRPCSTT struct {
	Port string `json:"port,omitempty"`
}

type LifeCycleCSTT struct {
	PostStart *PostStartLCCSTT `json:"postStart",omitempty`
	PreStop   *PreStopLCCSTT   `json:"preStop,omitempty"`
}

type PostStartLCCSTT struct {
	Exec *ExecPostLCCSTT `json:"exec,omitempty"`
}

type ExecPostLCCSTT struct {
	Command []string `json:"command,omitempty"`
}

type PreStopLCCSTT struct {
	Exec *ExecPreLCCSTT `json:"exec,omitempty"`
}

type ExecPreLCCSTT struct {
	Command []string `json:"command,omitempty"`
}

type StatusType struct {
	Replicas             float64 `json:"replicas,omitempty"`
	FullyLabeledReplicas float64 `json:"fullyLabeledReplicas,omitempty"`
	ObservedGeneration   float64 `json:"observedGeneration,omitempty"`
}
