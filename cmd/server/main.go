package cmd

import (
	"log"
	"net/http"

	"github.com/AviadKorakin/ReactiveOrganizationalUnitsMicroservice-HTTP-2/repositories"
	"github.com/AviadKorakin/ReactiveOrganizationalUnitsMicroservice-HTTP-2/routers"
)

func main() {
	client, err := repositories.NewMongoClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("Mongo connect error: %v", err)
	}
	defer repositories.Disconnect(client)

	mux := routers.NewRouter(client)
	log.Println("HTTP server starting on :8080 …")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
