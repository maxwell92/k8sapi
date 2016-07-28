type podList struct {
	Kind string `json: "kind"`
	ApiVersion string `json: "apiVersion"`
	Metadata metadataPLType `json: "metadata"`
	Items []itemsPLType	`json: "items"`
}

type metadataPLType struct {
	ResourceVersion string `json: "resourceVersion"`
}

type itemsPLType struct {
	Kind string `json: "kind"`
	ApiVersion string `json: "apiVersion"`
	Metadata metadataIType `json: "metadata"`
	Spec specType `json: "spec"`
	Status statusType `json: "status"`	
}


type metadataIType struct {
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

type specType struct {
	Volumes []volumesSType `json: "volumes"`
	Containers []containerSType `json: "containers"`
	RestartPolicy string `json: "restartPolicy"`
	TerminationGracePeriodSeconds float64 `json: "terminationGracePeriodSeconds"`
	ActiveDeadlineSeconds float64 `json: "activeDeadlineSeconds"`
	DnsPolicy string `json: "dnsPolicy"`
	NodeSelector map[string] string  `json: "nodeSelector"`
	NodeName string `json: "nodeName"`
	HostNetwork bool `json: "hostNetwork"`
	HostPID bool `json: "hostPID"`
	HostIPC bool `json: "hostIPC"`
	ImagePullSecrets []imagePullSecretsType `json: "imagePullSecrets"`

}

type volumesSType struct {
	Name string `json: "name"`
	HostPath hostPathVSType `json: "hostPath"`
	EmptyDir emptyDirVSType `json: "emptyDir"`
	PersistentVolumeClaim pvClaimVSType `json: "persistentVolumeClaim"`
	Rbd rbdVSType `json: "rbd"`
	ConfigMap configMapVSType `json: "configMap"`
}

type hostPathVSType struct {
	Path string `json: "path"`
}

type emptyDirVSType struct {
	Medium string `json: "medium"`
}

type pvClaimVSType struct {
	ClaimName string `json: "claimName"`
	ReadOnly string `json: "readOnly"`
}

type rbdVSType struct {
	Monitors []string `json: "monitors"`
	Image string `json: "image"`
	FsType string `json: "fsType"`
	Pool string `json: "pool"`
	User string `json: "user"`
	Keyring string `json: "keyring"`
	SecretRef secretRefrbdType `json: "secretRef"`
	ReadOnly bool `json: "readOnly"`
}

type secretRefrbdType struct {
	Name string `json: "name"`
}

type configMapVSType struct {
	Name string `json: "name"`
	Items []itemsConfigMapType `json: "items"`
}

type itemsConfigMapType struct {
	Key string `json: "key"`
	Path string `json: "path"`
}

type containerSType struct {
	Name string `json: "name"`
	Image string `json: "image"`
	Command []string `json: "command"`
	Args []string `json: "args"`
	WorkingDir string `json: "workingDir"`
	Ports []portsContainerType `json: "ports"`
	Env []envContainerType `json: "env"`
	Resources resourcesContainerType `json: "resources"`
	VolumeMounts []volumeMountsContainerType `json: "volumeMounts"`
	LivenessProbe livenessProbeContainerType `json: "livenessProbe"`
	ReadinessProbe readinessProbeContainerType `json: "readinessProbe"`
	Lifecycle lifecycleContainerType `json: "lifecycle"`
	TerminationMesasgePath string `json: "terminationMessagePath"`
	ImagePullPolicy string `json: "imagePullPolicy"`
	Stdin bool `json: "stdin"`
	StdinOnce bool `json: "stdinOnce"`
	Tty bool `json: "tty"`
}

type portsContainerType struct {
	Name string `json: "name"`
	HostPort hostPortPCType `json: "hostPort"`
	ContainerPort float64 `json: "containerPort"`
	Protocol string `json: "protocol"`
	HostIP string `json: "hostIP"`
}

type hostPortPCType struct {

}

type envContainerType struct {
	Name string `json: "name"`
	Value string `json: "value"`
	ValueFrom valueFromEnvContainer `json: "valueFrom"`
}

type valueFromEnvContainer struct {
	FieldRef fieldRefValueFromEnvCon `json: "fieldRef"`
	ConfigMapKeyRef configMapKeyEnvCon `json: "configMapKeyRef"`
	SecretKeyRef secretKeyRefEnvCon  `json: "secretKeyRef"`
}

type fieldRefValueFromEnvCon struct {
	ApiVersion string `json: "apiVersion"`
	FieldPath string `json: "fieldPath"`
}

type configMapKeyEnvCon struct {
	Name string `json: "name"`
	Key string `json: "key"`
}

type secretKeyRefEnvCon struct {
	Name string `json: "name"`
	Key string `json: "key"`
}

type resourcesContainerType struct {
    LimitsRsConType map[string] string `json: "limits"`
    Requests map[string] string  `json: "requests"`
}

type volumeMountsContainerType struct {
    Name string `json: "name"`
    ReadOnly bool `json: "readOnly"`
    MountPath string `json: "mountPath"`
}

type livenessProbeContainerType struct {
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


type readinessProbeContainerType struct {
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

type imagePullSecretsType struct {
    Name string `json: "name"`
}

type statusType struct {
	Phase string `json: "phase"`
	Conditions []conditionsStatType `json: "conditions"`
	Message string `json: "message"`
	Reason string `json: "reason"`
	HostIP string `json: "hostIP"`
	StartTime string `json: "startTime"`
	ContainerStatuses []containerStatuesStatType `json: "containerStatuses"`
}

type conditionsStatType struct {
    Type string `json: "type"`
    Status string `json: "status"`
    LastProbeTime string `json: "lastProbeTime"`
    LastTransition string `json: "lastTransition"`
    Reason string `json: "reason"`
    Message string `json: "mesage"`
} 

type containerStatuesStatType struct {
    Name string `json: "name"`
    State statCSSType `json: "state"`
    LastState lastStateCSSType `json: "lastState"`
    Ready bool `json: "ready"`
    RestartCount float64 `json: "restartCount"`
    Image string `json: "image"`
    ImageID string `json: "imageID"`
    ContainerID string `json: "containerID"`
}

type statCSSType struct {
    Waiting waitSCSSType `json: "waiting"`
    Running runSCSSType `json: "running"`
    Terminated termCSSType `json: "terminated"`
}

type waitSCSSType struct {
    Reason string `json: "reason"`
    Message string `json: "message"` 
}

type runSCSSType struct {
    StartedAt string `json: "startedAt"`
}

type termCSSType struct {
    ExitCode float64 `json: "exitCode"`
    Signal float64 `json: "signal"`
    Reason string `json: "reason"`
    Message string `json: "message"`
    StartedAt string `json: "startedAt"`
    FinishedAt string `json: "finishedAt"`
    ContainerID string `json: "containerID"`
}

type lastStateCSSType struct {
    Waiting waitLSCSSType `json: "waiting"`
    Running runLSCSSType `json: "running"`
    Terminated termLSCSSType `json: "terminated"`
}

type waitLSCSSType struct {
    Reason string `json: "reason"`
    Message string `json: "message"`
}

type runLSCSSType struct {
    StartedAt string `json: "startedAt"`
}

type termLSCSSType struct {
    ExitCode float64 `json: "exitCode"`
    Signal float64 `json: "signal"`
    Reason string `json: "reason"`
    Message string `json: "message"`
    StartedAt string `json: "startedAt"`
    FinishedAt string `json: "finishedAt"`
    ContainerID string `json: "containerID"`
}


func Podlist(w http.ResponseWriter, r *http.Request) {
	
	var response []byte
	var err error

	//response, err = Get("http://172.21.1.11:8080/api/v1/pods")
	response, err = Get("http://usa1:8080/api/v1/pods")
    defer response.Body.Close()
/*	
	client := &http.Client{}
	rep, err = http.NewRequest("GET", "http://usa1:8080/api/v1/pods"
	if err != nil {
		panic(err)
		log.Println(err)	
	}
		
	rep.Header.Add("Access-Control-Allow-Origin", "*")
	response, err := client.Do(req)
	if err != nil {
		panic(err)
		log.Println(err)	
	}
	defer response.Body.Close()
*/
	rs, err := resolveToStruct(response)		
	if err != nil {
		panic(err)
		log.Println(err)
	}


    num := len(rs.Items) 

	applist := make(AppListType, num)

	for i := 0; i < len(rs.Items); i++ {
        var dc []string
        dc = make([]string, 5)
        dc[0] = "shijilulian"
        dc[1] = "dianxin"

        applist[i].Healthz.PodsAvailable = "all"
        applist[i].Name = rs.Items[i].Metadata.Name
        applist[i].Label = rs.Items[i].Metadata.Labels
        applist[i].Datacenter = dc 
        applist[i].Replicas = 3
        applist[i].Worktime = rs.Items[i].Metadata.CreationTimeStamp
    }
		
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	json.NewEncoder(w).Encode(applist)
	
}

