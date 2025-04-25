// File: main.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AviadKorakin/ReactiveOrganizationalUnitsMicroservice-HTTP-2/repositories"
	"github.com/AviadKorakin/ReactiveOrganizationalUnitsMicroservice-HTTP-2/routers"
)

func main() {
	// 1) MongoDB connection
	client, err := repositories.NewMongoClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("Mongo connect error: %v", err)
	}
	defer repositories.Disconnect(client)

	// 2) Create your router (includes Swagger UI under /swagger/)
	mux := routers.NewRouter(client)

	// 3) Determine port (default to 10000 if not set)
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	addr := "0.0.0.0:" + port

	// 4) Start plain HTTP server; Render handles TLS + HTTP/2
	log.Printf("Starting server on http://%s …", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
