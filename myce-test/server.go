package main

import (
	"fmt"
	"log"
	"net/http"

	deployc "controller/deployment"
	endpointc "controller/endpoint"
	namespacec "controller/namespace"
	podc "controller/pod"
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
func Namespace(w http.ResponseWriter, r *http.Request) {
	var nc namespacec.NamespaceController

	nc.Get("http://master:8080/api/v1/namespaces")

	nc.Post("http://master:8080/api/v1/namespaces")

	nc.Get("http://master:8080/api/v1/namespaces")

	nc.Delete("http://master:8080/api/v1/namespaces/nginx-ns-test")

	nc.Get("http://master:8080/api/v1/namespaces")

}
func Pod(w http.ResponseWriter, r *http.Request) {
	var pc podc.PodController

	pc.Get("http://master:8080/api/v1/namespaces/default/pods")

	pc.Post("http://master:8080/api/v1/namespaces/default/pods")

	pc.Get("http://master:8080/api/v1/namespaces/default/pods")

	pc.Delete("http://master:8080/api/v1/namespaces/default/pods/nginx-pd-test")

	pc.Get("http://master:8080/api/v1/namespaces/default/pods")

}
func main() {

	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/deploy", Deploy)
	http.HandleFunc("/endpoint", Endpoint)
	http.HandleFunc("/namespace", Namespace)
	http.HandleFunc("/pod", Pod)
	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
