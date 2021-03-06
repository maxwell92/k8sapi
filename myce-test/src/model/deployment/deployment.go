package deployment

// DeploymentList for Get @/apis/extensions/v1beta1/namespaces/{namespace}/deployments
type DeploymentList struct {
	Kind       string           `json:"kind"`
	ApiVersion string           `json:"apiVersion"`
	Metadata   MetadataListType `json:'metadata"`
	Items      []ItemsListType  `json:"items"`
}

type MetadataListType struct {
	SelfLink        string `json:"selfLink"`
	ResourceVersion string `json:"resourceVersion"`
}

type ItemsListType Deployment

// Deployment for Post @/apis/extensions/v1beta1/namespaces/{namespace}/deployments
type Deployment struct {
	ApiVersion string     `json:"apiVersion"`
	Kind       string     `json:"kind"`
	Metadata   MetadataDL `json:"metadata"`
	Spec       SpecDL     `json:"spec"`
}

type MetadataDL struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	CreationTimestamp string            `json:"creationTimestamp,omitempty"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations,omitempty"`
}

type SpecDL struct {
	Replicas float64      `json:"replicas"`
	Template TemplateSpec `json:"template"`
}

type TemplateSpec struct {
	Metadata MetadataTSDL `json:"metadata"`
	Spec     SpecTSDL     `json:"spec"`
}

type MetadataTSDL struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
}

type SpecTSDL struct {
	Volumes []VolumesSTSDL `json:"volumes,omitempty"`
	//Containers []ContainersSTSDL `json:"containers"`
	Containers []ContainerType `json:"containers"`
}

type VolumesSTSDL struct {
	Name                  string           `json:"name"`
	HostPath              *HostPathVSTSDL  `json:"hostPath,omitempty"`
	EmptyDir              *EmptyDirVSTSDL  `json:"emptyDir,omitempty"`
	PersistentVolumeClaim *PvClaimVSTSDL   `json:"persistentVolumeClaim,omitempty"`
	Rbd                   *RbdVSTSDL       `json:"rbd,omitempty"`
	ConfigMap             *ConfigMapVSTSDL `json:"configMap,omitempty"`
}

type HostPathVSTSDL struct {
	Path string `json:"path"`
}

type EmptyDirVSTSDL struct {
	Medium string `json:"medium"`
}

type PvClaimVSTSDL struct {
	ClaimName string `json:"claimName"`
	ReadOnly  bool   `json:"readOnly"`
}

type RbdVSTSDL struct {
	Monitors  []string          `json:"monitors"`
	Image     string            `json:"image"`
	FsType    string            `json:"fsType"`
	Pool      string            `json:"pool"`
	User      string            `json:"user"`
	Keyring   string            `json:"keyring"`
	SecretRef *SecretRefRVSTSDL `json:"secretRef"`
	ReadOnly  bool              `json:"readOnly"`
}

type SecretRefRVSTSDL struct {
	Name string `json: name"`
}

type ConfigMapVSTSDL struct {
	Name  string           `json:"name"`
	Items []ItemsConfigMap `json:"items"`
}

type ItemsConfigMap struct {
	Key  string `json:"key"`
	Path string `json:"Path"`
}

type ContainersSTSDL struct {
	Name           string                   `json:"name"`
	Image          string                   `json:"image"`
	Command        []string                 `json:"command,omitempty"`
	Args           []string                 `json:"args,omitempty"`
	Ports          []PortContainer          `json:"ports"`
	Env            []EnvContainer           `json:"env,omitempty"`
	Resources      *ResourcesContainer      `json:"resources,omitempty"`
	VolumeMounts   []VolumeMountsContainer  `json:"volumeMounts,omitempty"`
	LivenessProbe  *LivenessProbeContainer  `json:"livenessProbe,omitempty"`
	ReadinessProbe *ReadinessProbeContainer `json:"readinessProbe,omitempty"`
	Lifecycle      *LifecycleContainer      `json:"lifecycle,omitempty"`
}

type ContainerType ContainersSTSDL

type PortContainer struct {
	Name          string  `json:"name"`
	HostPort      float64 `json:"hostPort"`
	ContainerPort float64 `json:"containerPort"`
	Protocol      string  `json:"protocol"`
	HostIP        string  `json:"hostIP"`
}

type EnvContainer struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ResourcesContainer struct {
	Limits   map[string]string `json:"limits"`
	Requests map[string]string `json:"requests"`
}

type VolumeMountsContainer struct {
	Name      string `json:"name"`
	ReadOnly  bool   `json:"readOnly"`
	MountPath string `json:"mountPath"`
}

type LivenessProbeContainer struct {
	Exec                *ExecLiveProbeType    `json:"exec,omitempty"`
	HttpGet             *HttpGetLiveProbeType `json:"httpGet,omitempty"`
	TcpSocket           *TcpLiveProbeType     `json:"tcpSocket,omitempty"`
	InitialDelaySeconds float64               `json:"initialDelaySeconds"`
	TimeoutSeconds      float64               `json:"timeoutSeconds"`
	PeriodSeconds       float64               `json:"periodSeconds"`
	SuccessThreshold    float64               `json:"successThreshold"`
	FailureThreshold    float64               `json:"failureThreshold"`
}

type ExecLiveProbeType struct {
	Command []string `json:"command"`
}

type HttpGetLiveProbeType struct {
	Path        string           `json:"path"`
	Port        float64          `json:"port"`
	Host        string           `json:"host"`
	Scheme      string           `json:"scheme"`
	HttpHeaders []HeadersGLPType `json:"httpHeaders"`
}

type HeadersGLPType struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TcpLiveProbeType struct {
	Port float64 `json:"port"`
}

type ReadinessProbeContainer struct {
	HttpGet             *HttpGetReadProbeType `json:"httpGet,omitempty"`
	InitialDelaySeconds float64               `json:"initialDelaySeconds"`
	TimeoutSeconds      float64               `json:"timeoutSeconds"`
	PeriodSeconds       float64               `json:"periodSeconds"`
	SuccessThreshold    float64               `json:"successThreshold"`
	FailureThreshold    float64               `json:"failureThreshold"`
}

type HttpGetReadProbeType struct {
	Path        string           `json:"path"`
	Port        float64          `json:"port"`
	Host        string           `json:"host"`
	Scheme      string           `json:"scheme"`
	HttpHeaders []HeadersGRPType `json:"httpHeaders"`
}

type HeadersGRPType struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type LifecycleContainer struct {
	PostStart *PostStartLCType `json:"postStart,omitempty"`
	PreStop   *PreStopLCType   `json:"preStop,omitempty"`
}

type PostStartLCType struct {
	Exec *ExecPSLCType `json:"exec,omitempty"`
}

type ExecPSLCType struct {
	Command []string `json:"command"`
}

type PreStopLCType struct {
	Exec *ExecPrSLCType `json:"exec,omitempty"`
}

type ExecPrSLCType struct {
	Command []string `json:"command"`
}
