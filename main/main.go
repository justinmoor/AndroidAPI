package main

import("log"
		"net/http"
	"fmt"
)

func main(){
	fmt.Println("Setting up HTTP server...\n")
	fmt.Println("Connecting to database...")
	databaseConnection()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))

}
