package route

import (
	"CanYouGetTo20_REST-API/pkg/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(routes Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.Logger)

	log.Println("Setting up routes...")
	log.Println("Available routes:")

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		log.Println(route.Name, route.Method, route.Pattern)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
