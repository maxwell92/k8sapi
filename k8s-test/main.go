package main

import (
	"net/http"
	"log"
)



func main() {
	router := NewRouter()

    //router.Handle("/statics/index.html", http.StripPrefix("/statics/", http.FileServer(http.Dir("statics/"))))
 
	log.Fatal(http.ListenAndServe(":10000", router))
}

