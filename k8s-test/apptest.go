package main 
import (
    "net/http"
    "net/http/httptest"
    "fmt"
    "log"
)

func main() {
    handler := func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "something failed", http.StatusInternalServerError)
    }

    req, err := http.NewRequest("GET", "http://localhost:10000/deployment", nil)
    if err != nil {
        log.Fatal(err)
    }

    w := httptest.NewRecorder()
    handler(w, req)

    fmt.Printf("%d – %s", w.Code, w.Body.String())
}
