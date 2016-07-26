package main 

import (
	"fmt"
	"html"
	"log"
	"net/http"
    "encoding/json"
	"crypto/tls"
	"io/ioutil"

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

func resolveToStruct(response []byte) (s *podList, err error) {
	err = json.Unmarshal(response, &s)
	return s, err 
}



func Podlist(w http.ResponseWriter, r *http.Request) {
	
	var response []byte
	var err error

	//response, err = Get("http://172.21.1.11:8080/api/v1/pods")
	response, err = Get("http://usa1:8080/api/v1/pods")
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


