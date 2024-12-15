package main

import (
	"learn_http_interface/implement/routes"
	"log"
	"net/http"
)

func main() {
	mux := routes.RegisterRoutes()

	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
