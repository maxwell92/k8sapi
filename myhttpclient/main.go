package main

import (
    "fmt"
    "log"
    "httpclient"
    "crypto/tls"
)

func main() {
    client := httpclient.NewHttpClient("http://usa1", "8080")
    response, err := client.Get("http://usa1:8080/api/v1/pods")
    if err != nil {
        log.Println(err)
    }

    fmt.Println(response)
}
