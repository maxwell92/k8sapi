package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"Podlist",
		"GET",
		"/podlist",
		Podlist,
	},

	Route{
		"Podetails",
		"GET",
		"/podlist/{podname}",
		Podetails,
	},

	Route{
		"delPod",
		"DELETE",
		"/podlist/{podname}",
		delPod,
	},
}
