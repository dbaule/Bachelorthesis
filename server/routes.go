package main

import "net/http"

// Route Defines a route object
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes Defines a collection of route objects
type Routes []Route

// Array with all available routes
var routes = Routes{
	// handles every "/" server request with a hello statement connected with the given value
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"CreateViewController",
		"POST",
		"/viewcontrollers",
		CreateViewController(session),
	},
	Route{
		"GetViewController",
		"GET",
		"/getViewController",
		GetViewController(session),
	},
	Route{
		"GetDifferences",
		"GET",
		"/getDifferences",
		GetDifferences(session),
	},
	Route{
		"CreateEvent",
		"POST",
		"/events",
		CreateEvent(session),
	},
}
