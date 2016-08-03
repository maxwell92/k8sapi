package main

import (
	"fmt"
	"log"
	"net/http"

	deployc "controller/deployment"
	endpointc "controller/endpoint"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello myce")
}

func Deploy(w http.ResponseWriter, r *http.Request) {
	var dc deployc.DeploymentController

	dc.Get("http://master:8080/apis/extensions/v1beta1/namespaces/default/deployments")

	dc.Post("http://master:8080/apis/extensions/v1beta1/namespaces/default/deployments")

	dc.Get("http://master:8080/apis/extensions/v1beta1/namespaces/default/deployments")

	dc.Delete("http://master:8080/apis/extensions/v1beta1/namespaces/default/deployments/nginx-test")

	dc.Get("http://master:8080/apis/extensions/v1beta1/namespaces/default/deployments")
}

func Endpoint(w http.ResponseWriter, r *http.Request) {
	var ec endpointc.EndpointController

	ec.Get("http://master:8080/api/v1/namespaces/default/endpoints")

	ec.Post("http://master:8080/api/v1/namespaces/default/endpoints")

	ec.Get("http://master:8080/api/v1/namespaces/default/endpoints")

	ec.Delete("http://master:8080/api/v1/namespaces/default/endpoints/nginx-ep-test")

	ec.Get("http://master:8080/api/v1/namespaces/default/endpoints")

}

func main() {

	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/deploy", Deploy)
	http.HandleFunc("/endpoint", Endpoint)
	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
