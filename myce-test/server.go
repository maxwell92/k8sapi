package main
import (
    "fmt"
    "log"
    "net/http"
    
    deployc "controller/deploymentcontroller"
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


func main() {
    
    http.HandleFunc("/", sayhelloName)
    http.HandleFunc("/deploy", Deploy)
    err := http.ListenAndServe(":10000", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
