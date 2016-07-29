package main
import (
    "log"
    "net/http"
    "html/template"
    "fmt"
)

func Hello(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        t, err := template.ParseFiles("template/html/homepage.html")
        if err != nil {
            return 
        }

        err = t.Execute(w, nil)
    }
}

func App(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello app")
}

func StaticServer(prefix string, staticDir string) {
    http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(staticDir))))
    return 
}

func main() {
    StaticServer("/template/", "./template")
    http.HandleFunc("/", Hello)
    http.HandleFunc("/app", App)
    err := http.ListenAndServe(":10000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}
