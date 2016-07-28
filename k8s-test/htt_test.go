package main

import (
    "net/http"
	"io/ioutil"
    "crypto/tls"
    "testing"
)

func Test_Deployment1(t *testing.T) {
    tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},	
	}

    url := "http://localhost:10000/deployment"

    client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
       t.Error(err) 
	}

	defer resp.Body.Close()
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        t.Error(err) 
	}

    if resp.StatusCode != 200 {
        t.Error(resp.Status)
        t.Error(string(body))
    }

 
}
