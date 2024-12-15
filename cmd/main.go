package main

import (
	"learn_http_interface/implement/routes"
	"log"
	"net/http"
)

func main() {
	mux := routes.RegisterRoutes()

	log.Println("Server is running on http://127.0.0.1:3000")
	if err := http.ListenAndServe("127.0.0.1:3000", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
