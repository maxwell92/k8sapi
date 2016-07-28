package main

import (
	"net/http"
    "testing"
)



func Test_appDeploy1(t *testing.T) {
    
    http.HandleFunc("/another", appDeploy) 
	err := http.ListenAndServe(":10000", nil)
    if err != nil {
        t.Error(err)
    }
}

