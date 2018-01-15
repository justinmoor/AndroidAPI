package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func NewRouter() *mux.Router{
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("\nSetting up routes...")
	fmt.Println("Available routes: \n")
	for _, route := range routes{
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		fmt.Println(route.Name, route.Method, route.Pattern)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	fmt.Println("\nWaiting for incoming requests...\n")
	return router
}

