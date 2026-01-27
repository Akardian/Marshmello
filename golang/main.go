package main

import (
	"log"
	"lookup/routes"
	"net/http"
)

func main() {
	// API routes for htmx
	routes.RegisterRoutes()

	// Start server
	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
