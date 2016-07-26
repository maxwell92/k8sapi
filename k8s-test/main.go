package main

import (
	"net/http"
	"log"
)



func main() {
	router := NewRouter()

    //http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("/static/"))))
    http.Handle("/statics", http.FileServer(http.Dir("./static/")))
	log.Fatal(http.ListenAndServe(":10000", router))
}

