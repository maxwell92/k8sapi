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
    Kind string `json:"kind"`
    ApiVersion string `json:"apiVersion"`
    Item []itemsType `json:"items"`
}

type itemsType struct {
    MetadataType metadataType1 `json:"metadata"`
    Spec specType `json:"spec"`
    Status statusType `json:"status"`
}

type metadataType1 struct {
    Name string `json:"name"`
    Namespace string `json:"namespace"`
}

type specType struct {
    Volumes []volumesType `json:"volumes"`
    Containers []containersType `json:"containers"`
}

type volumesType struct {
    Name string `json:"name"`
}

type containersType struct {
    Name string `json:"name"`
    Image string `json:"image"`
    Command []string `json:"command"`
    Resources resourcesType `json:"resources"`
    VolumeMounts []volumeMountsType `json:"volumeMounts"`
}

type resourcesType struct {
    Limits limitsType `json:"limits"`
}

type limitsType struct {
    Cpu string `json:"cpu"`
    Mem string `json:"mem"`
}

type volumeMountsType struct {
    Name string `json:"name"`
    ReadOnly bool `json:"readOnly"`
    MountPath string `json:"mountPath"`
}

type statusType struct {
    HostIP string `json:"hostIP"`
    PodIP string `json:"podIP"`
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
/*
func resolveToStruct(response []byte) (err error) {
    var f interface{}
    err = json.Unmarshal(response, &f)
    m := f.(map[string]interface{})
    for k, v := range m {
	   switch vv := v.(type) {
	   case string:
		   fmt.Println(k, "is string", vv)
	   case int:
		   fmt.Println(k, "is int", vv)
	   case float64:
		   fmt.Println(k, "is float64", vv)
	   case []interface{}:
		   fmt.Println(k, "is an array:")
		   for i, u := range vv {
			   fmt.Println(i, u)
		   }
	   default:
		   fmt.Println(k, "is nothing.")

	   }
    }
    return nil
}
*/

func resolveToStruct(response []byte) (s *podlist, err error) {
    err = json.Unmarshal(response, &s)
    return s, err
}

func main() {
    url := os.Args[1]

    var response []byte
    var err error

    response, err = Get(url)
    if err != nil {
	panic(err)
 	log.Println(err)
    }


    rs, err := resolveToStruct(response)
    if err != nil {
	panic(err)
	log.Println(err)
    }

    for i := 0; i < len(rs.Item); i++ {
	fmt.Printf("%d\n", i + 1)
        fmt.Println(rs.Item[i].MetadataType.Name)
	fmt.Println(rs.Item[i].MetadataType.Namespace)
	fmt.Println(rs.Item[i].Spec.Containers[0].Image)
	fmt.Println(rs.Item[i].Status.HostIP)
	fmt.Println(rs.Item[i].Status.PodIP)
    }
}
