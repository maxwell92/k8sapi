package main

import (
    "testing"
    "net/http"
    "net/http/httptest"
)



func Test_Podlist1(t *testing.T) {


    handler := func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "something failed", http.StatusInternalServerError)
    }

    //req, err := http.NewRequest("GET", "http://localhost:10000/deployment", nil) 
    req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
    if err != nil {
        t.Error(err)
    }

    t.Error(req)

    w := httptest.NewRecorder()
    handler(w, req)

    t.Error(w.Code, w.Body.String())
}


/*
package main

import (
    "testing"
    "log"
    "net/http"
    "net/http/httptest"
)

func Test_Podlist1(t *testing.T) (err string) {
    handler := func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "something failed", http.StatusInternalServerError)
    }

    req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
    if err != nil {
        log.Fatal(err)
        return "err"
    }

    w := httptest.NewRecorder()
    handler(w, req)

    t.Error(w.Code, w.Body.String())
    return "ok"
}
*/
