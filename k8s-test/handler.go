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
	"io"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, This is k8sapi!", html.EscapeString(r.URL.Path))
}
/*
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




func Podlist(w http.ResponseWriter, r *http.Request) {
	
	var response []byte
	var err error

	response, err = Get("http://172.21.1.11:8080/api/v1/pods")
	if err != nil {
		panic(err)
		log.Println(err)	
	}
		
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

		
	json.NewEncoder(w).Encode(applist)
    
}

func resolveToStruct(response []byte) (s *podList, err error) {
	err = json.Unmarshal(response, &s)
	return s, err 
}



*/

func Post(url string, body io.Reader) (rtn []byte, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},	
	}

	client := &http.Client{Transport: tr}
	//resp, err := client.Get(url)
	resp, err := client.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		log.Println(err)
		panic(err)
		return nil, err	
	}

	defer resp.Body.Close()

	rtn, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err	
	}
	return rtn, nil
}




func appDeploy(w http.ResponseWriter, r *http.Request) {
	//tempApp := new(appDeployment)
    var tempApp appDeployment
	tempApp.Name = "nginx-deployment-test"
	tempApp.Namespace = "default"

    var dc []string
    dc = make([]string, 5)
    dc[0] = "shijilulian"
    dc[1] = "dianxin"

	tempApp.Datacenter = dc 
	tempApp.Image = "nginx:1.7.9"
	tempApp.Replicas = 3

	spec := make(map[string]string, 2)
	spec["cpu"] = "2"
	spec["mem"] = "8G"

	tempApp.Spec.Request = spec

    configfile := make([]string, 1)
    configfile[0] = "/home/ubuntu/nginx.conf"
	tempApp.ConfigFile = configfile

    mountpoints := make([]string, 1)
    mountpoints[0] = "/usr/share/nginx"
    tempApp.MountPoints = make([]mountPointsType, 1)
	tempApp.MountPoints[0].MountPath = "/usr/share/nginx" 

	env := make(map[string]string, 2)
	env["username"] = "maxwell"
	env["password"] = "123456"

    tempApp.Env = make([]envType, 2)
	tempApp.Env[0].Name = "username"
    tempApp.Env[0].Value = "maxwell"
    tempApp.Env[1].Name = "password"
    tempApp.Env[1].Value = "123456" 

	labels := make(map[string]string, 1)
	labels["app"] = "nginx"

	tempApp.Label = labels


	tempDeploy := new(basicDeployment)
	tempDeploy.ApiVersion = "extensions/v1beta1"
	tempDeploy.Kind = "Deployment"
	tempDeploy.Metadata.Name = tempApp.Name
    tempDeploy.Metadata.Labels = tempApp.Label 
	tempDeploy.Spec.Replicas = tempApp.Replicas
    tempDeploy.Spec.Template.Metadata.Labels = tempApp.Label
    tempDeploy.Spec.Template.Spec.Containers = make([]containersSTSDL, 1)
	tempDeploy.Spec.Template.Spec.Containers[0].Name = tempApp.Name
    tempDeploy.Spec.Template.Spec.Containers[0].Image = tempApp.Image

	//var body io.Writer
//	json.NewEncoder(w).Encode(tempDeploy)
    //var bdy io.Reader
    
    //io.Copy(bdy, body)
    //io.Copy(w, bdy)

    var result []byte
    result, _ = json.Marshal(tempDeploy)
    
    
	_, err := Post("http://172.21.1.11:8080/apis/extensions/v1beta1/namespaces/default/deployments", strings.NewReader(string(result)))
	if err != nil {
		panic(err)
		log.Println(err)
	}


	fmt.Fprintln(w, "created")

    fmt.Fprintln(w, string(result))

}
