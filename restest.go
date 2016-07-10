package main
import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
)

func getuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "you are get user %s\n", uid)
}

func modifyuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "your are modify user %s\n", uid)
}

func deleteuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "your are delete user %s\n", uid)
}

func adduser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	fmt.Fprintf(w, "your are add user %s\n", uid)
}

func main() {
	mux := routes.New()
	mux.Get("/user/:uid", getuser)
	mux.Post("/user/:uid", modifyuser)
	mux.Del("/user/:uid", deleteuser)
	mux.Put("/user/", adduser)

	http.Handle("/", mux)
	http.ListenAndServe(":10001", nil)
}
