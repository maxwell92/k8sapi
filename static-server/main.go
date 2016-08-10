package main

import (
	"log"
	"net/http"
)

//succeed one
/*
func main() {
	//http.Handle("/app", http.FileServer(http.Dir("statics")))
    //fs := http.FileServer(http.Dir("/Users/yp-tc-m-7019/mworks/k8sapi/static-server/statics"))

    //relative path works well
    fs := http.FileServer(http.Dir("./statics"))
    http.Handle("/statics/", http.StripPrefix("/statics/", fs))
    //http.Handle("/tmp/", fs)
	log.Fatal(http.ListenAndServe(":10001", nil))
}
*/

/*
func main() {
    //absolute path works well, too
    fs := http.FileServer(http.Dir("/Users/yp-tc-m-7019/mworks/k8sapi/static-server/statics"))
    http.Handle("/statics/", http.StripPrefix("/statics/", fs))
	log.Fatal(http.ListenAndServe(":10001", nil))
}
*/
/*
func main() {
    fs := http.FileServer(http.Dir("./statics"))

    //app cann't work
    //http.Handle("/app/", http.StripPrefix("/app/", fs))


    // /statics also works well
    http.Handle("/statics", http.StripPrefix("/statics/", fs))

	log.Fatal(http.ListenAndServe(":10001", nil))
}
*/

func main() {
	fs := http.FileServer(http.Dir("./statics"))
	http.Handle("/statics/", http.StripPrefix("/statics/", fs))
	log.Fatal(http.ListenAndServe(":10001", nil))
}

//usage
/*
func StaticServer(prefix string, staticDir string) {
    http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(staticDir))))
    return
}
*/
