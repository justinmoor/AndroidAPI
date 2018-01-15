package main

import "net/http"

type Route struct{
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes [] Route

var routes = Routes{
	Route{
		"SubmitScore",
		"POST",
		"/submitscore",
		submitScore,
	},
	Route{
		"ShowScore",
		"GET",
		"/showscore",
		showScore,
	},
}
