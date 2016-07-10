package main
import (
	"fmt"
	"html"
	"log"
	"net/http"
	"encoding/json"
	"crypto/tls"
	"io/ioutil"
	"strings"

	"github.com/gorilla/mux"
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

type PodinfoType struct {
	Name 	string 	`json:"Name"`
	Ready 	bool 	`json:"Ready"`
	Status 	string 	`json:"Status"`
	Restarts	int	`json:"Restarts"`
	StartTime	string	`json:"StartTime"`
}

type PodlistType []PodinfoType

type PodetailsType struct {
	Name	string	`json:"Name"`
	Namespace	string	`json:"Namespace"`
	Node	string	`json:"Node"`
	Start_Time	string	`json:"Start Time"`
//	Labels	string	`json:"Labels"`
	Status	string	`json:"Status"`
	IP	string	`json:"IP"`
//	Controllers	string	`json:"Controllers"`
//	Containers	[]ContainerType	`json:"Containers"`
//	Conditions	[]ConditionType	`json:"Conditions"`
} 


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, This is k8sapi!", html.EscapeString(r.URL.Path))
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



func Podlist(w http.ResponseWriter, r *http.Request) {
	
	var response []byte
	var err error

	response, err = Get("http://usa1:8080/api/v1/pods")
	if err != nil {
		panic(err)
		log.Println(err)	
	}
		
	rs, err := resolveToStruct(response)		
	if err != nil {
		panic(err)
		log.Println(err)
	}

	fmt.Println("Name\t\tReady\t\tStatus\t\tRestarts\t\tAge\n")

	for i := 0; i < len(rs.Items); i++ {
		fmt.Printf("%s\t\t", rs.Items[i].Metadata.Name)		
		if rs.Items[i].Status.ContainerStatuses[0].Ready == true {
			fmt.Printf("%d\t\t", 1)	
		} else {
			fmt.Printf("%d\t\t", 0)	
		}

		fmt.Printf("%s\t\t", rs.Items[i].Status.Phase)
		fmt.Printf("%d\t\t", rs.Items[i].Status.ContainerStatuses[0].RestartCount)
		fmt.Printf("%s\t", rs.Items[i].Status.StartTime)
		fmt.Printf("\n")

	}


	podlist := make(PodlistType, 20)

	for i := 0; i < len(rs.Items); i++ {
		podlist[i].Name = rs.Items[i].Metadata.Name
		podlist[i].Ready = rs.Items[i].Status.ContainerStatuses[0].Ready
		podlist[i].Status = rs.Items[i].Status.Phase
		podlist[i].Restarts = rs.Items[i].Status.ContainerStatuses[0].RestartCount
		podlist[i].StartTime = rs.Items[i].Status.StartTime
	}

		
	json.NewEncoder(w).Encode(podlist)

}

func Podetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	podname := vars["podname"]
	fmt.Fprintln(w, "Pod [", podname, "] Details: ")

	var response []byte
	var err error

	response, err = Get("http://usa1:8080/api/v1/pods")
	if err != nil {
		panic(err)
		log.Println(err)	
	}
		
	rs, err := resolveToStruct(response)		
	if err != nil {
		panic(err)
		log.Println(err)
	}

	var i int
	for i = 0; i < len(rs.Items); i++ {
		if strings.Compare(rs.Items[i].Metadata.Name, podname) == 0 {
			fmt.Println("Found")	
			break
		}
	}
	
	if i == len(rs.Items) {
		fmt.Println("Not Found!")	
		fmt.Fprintln(w, "Not Found!")
	}

	var podetails PodetailsType
	podetails.Name = rs.Items[i].Metadata.Name
	podetails.Namespace = rs.Items[i].Metadata.Namespace
	podetails.Node = rs.Items[i].Spec.NodeName
	podetails.Start_Time = rs.Items[i].Status.StartTime
	podetails.Status = rs.Items[i].Status.Phase
	podetails.IP = rs.Items[i].Status.PodIP

	json.NewEncoder(w).Encode(podetails)

}


func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/podlist", Podlist)
	router.HandleFunc("/podlist/{podname}", Podetails)

	log.Fatal(http.ListenAndServe(":10000", router))
	
}
