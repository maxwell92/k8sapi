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
			var podetails PodetailsType
			podetails.Name = rs.Items[i].Metadata.Name
			podetails.Namespace = rs.Items[i].Metadata.Namespace
			podetails.Node = rs.Items[i].Spec.NodeName
			podetails.Start_Time = rs.Items[i].Status.StartTime
			podetails.Status = rs.Items[i].Status.Phase
			podetails.IP = rs.Items[i].Status.PodIP

			json.NewEncoder(w).Encode(podetails)

			break
		}
	}
	
	if i == len(rs.Items) {
		fmt.Println("Not Found!")	
		fmt.Fprintln(w, "Not Found!")
	}

	
}

func delPod(w http.ResponseWriter, r *http.Request) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},	
	}

	client := &http.Client{Transport: tr}

	vars := mux.Vars(r)
	podname := vars["podname"]
	fmt.Fprintln(w, "Pod [", podname, "] Details: ")

	
	req, err := http.NewRequest("DELETE", "http://usa1:8080/api/v1/namespaces/default/pods/"+podname, nil)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		panic(err)
	}


/*	var response []byte
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
*/
	defer resp.Body.Close()

/*	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		panic(err)
	}
*/
	fmt.Fprintln(w, "Deleted!")

}
