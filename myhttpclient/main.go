package main

import (
	"fmt"
	"httpclient"
	"log"
	"strings"
)

func main() {
	/*
	   response, err := client.Get("http://master:8080/api/v1/pods")
	   if err != nil {
	       log.Println(err)
	   }
	*/

	client := httpclient.NewHttpClient("http://master", "8080")
	fmt.Println(rcJson)
	response, err := client.Post("http://master:8080/api/v1/namespaces/default/replicationcontrollers", strings.NewReader(rcJson))
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(response))
}
