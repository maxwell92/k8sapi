package main
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"encoding/json"
	"os"
	"log"
)

type podlist struct {
	Kind string `json: "kind"`
	ApiVersion string `json: "apiVersion"`
	Metadata metadataType `json: "metadata"`
	Items []itemsType `json: "items"` 
}

type metadataType struct {
	SelfLink string `json: "selfLink"`
	ResourceVersion string `json: "resourceVersion"`
}

type itemsType struct {
	Metadata metadataType1 `json: "metadata"`
	Spec specType `json: "spec"`
	Status statusType `json: "status"`
}

type metadataType1 struct {
	Name string `json: "name"`
	Namespace string `json: "namespace"`
	SelfLink string `json: "selfLink"`
	Uid string `json: "uid"`
	ResourceVersion string `json: "resourceVersion"`
	CreationTimestamp string `json: "creationTimestamp"`
	Annotations annotationsType `json: "annotations"`
}

type annotationsType struct {
	Kubernetes_io_limit_ranger string `json: "kubernetes.io/limit-ranger"`
}

type specType struct {
	Volumes []volumesType `json: "volumes"`
	Containers []containersType `json: "containers"`
	RestartPolicy string `json: "restartPolicy"`
	TerminationGracePeriodSeconds int `json: "terminationGracePeriodSeconds"`
	DnsPolicy string `json: "dnsPolicy"`
	ServiceAccountName string `json: "serviceAccountName"`
	ServiceAccount string `json: "serviceAccount"`
	NodeName string `json: "nodeName"`
	SecurityContext securityContextType `json: "securityContext"`
}

type volumesType struct {
	Name string `json: "name"`
	Secret secretType `json: "secret"`
}

type secretType struct {
	SecretName string `json: "secretName"`
}

type containersType struct {
	Name string `json: "name"`
	Image string `json: "image"`
	Command []string `json: "command"`
	Env []envType `json: "env"`
	Resources resourcesType `json: "resources"`
	VolumeMounts []volumeMountsType `json: "volumeMounts"`
	TerminationMessagePath string `json: terminationMessagePath"`
	ImagePullPolicy string `json: "imagePullPolicy"`
}

type envType struct {
	Name string `json: "name"`
	ValueFrom valueFromType `json: "valueFrom"`
}

type valueFromType struct {
	ConfigMapKeyRef configMapKeyRefType `json: "configMapKeyRef"`
}

type configMapKeyRefType struct {
	Name string `json: "name"`
	Key string `json: "key"`
}

type resourcesType struct {
	Limits limitsType `json: "limits"`
	Requests requestsType `json: "requests"`
}

type limitsType struct {
	Cpu string `json: "cpu"`
	Memory string `json: "memory"`
}

type requestsType struct {
	Cpu string `json: "cpu"`
	Memory string `json: "memory"`
}

type volumeMountsType struct {
	Name string `json: "name"`
	ReadOnly bool `json: "readOnly"`
	MountPath string `json: "mountPath"`
}	

type securityContextType struct {
	
}

type statusType struct {
	Phase string `json: "phase"`
	Conditions []conditionsType `json: "conditions"`
	HostIP string `json: "hostIP"`
	PodIP string `json: "podIP"`
	StartTime string `json: "startTime"`
	ContainerStatuses []containerStatusesType `json: "containerStatuses"`
}

type conditionsType struct {
	Jtype string `json: "type"`
	Status string `json: "status"`
	LastProbeTime string `json: "lastProbeTime"`
	LastTransitionTime string `json: "lastTransitionTime"`
	Reason string `json: "reason"`
}

type containerStatusesType struct {
	Name string `json: "name"`
	State stateType `json: "state"`
	LastState lastStateType `json: "lastState"`
	Ready bool `json: "ready"`
	RestartCount int `json: "restartCount"`
	Image string `json: "image"`
	ImageID string `json: "imageID"`
	ContainerID string `json: "containerID"`
}

type stateType struct {
	Terminated terminatedType `json: "terminated"`
}

type terminatedType struct {
	ExitCode int `json: "exitCode"`
	Reason string `json: "reason"`
	StartedAt string `json: "startedAt"`
	FinishedAt string `json: "finishedAt"`
	ContainerID string `json: "containerID"`
}

type lastStateType struct {
	
}



func Get(url string) (body []byte, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},	
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		panic(err)
		return nil, err	
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err	
	}

	return body, nil
}

func resolveToStruct(response []byte) (s *podlist, err error) {
	err = json.Unmarshal(response, &s)
	return s, err 
}


func main() {
	url := os.Args[2]
	method := os.Args[1]		

	fmt.Println(method)
	fmt.Println(url)

	var response []byte
	response, err := Get(url)
	if err != nil {
		panic(err)
		log.Println(err)	
	}

	rs, err := resolveToStruct(response)		
	if err != nil {
		panic(err)
		log.Println(err)
	}

	for i := 0; i < len(rs.Items); i++ {
		fmt.Println(rs.Items[i].Metadata.Name)		
	}
}
