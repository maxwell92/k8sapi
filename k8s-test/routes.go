package main
import (
	"net/http"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route


var routes = Routes {
/*	Route {
		"Index",
		"GET",
		"/",
		Index,
	},
*/

	Route {
		"statics",
		"GET",
		"/",
		http.Handle("/", http.FileServer(http.Dir("statics"))),
	},

	Route {
		"Podlist",
		"GET",
		"/podlist",
		Podlist,
	},


}
