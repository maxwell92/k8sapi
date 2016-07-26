package main 
import (
	"net/http"
	"log"
)

func main() {
	//http.Handle("/statics/index.html", http.StripPrefix("/statics/", http.FileServer(http.Dir("/static/index.html"))))
	http.Handle("/", http.FileServer(http.Dir("statics")))
	log.Fatal(http.ListenAndServe(":10001", nil))

}


    
