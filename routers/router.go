package routers

import (
	"net/http"

	"github.com/AviadKorakin/ReactiveOrganizationalUnitsMicroservice-HTTP-2/controllers"
	_ "github.com/AviadKorakin/ReactiveOrganizationalUnitsMicroservice-HTTP-2/docs"
	"github.com/AviadKorakin/ReactiveOrganizationalUnitsMicroservice-HTTP-2/repositories"
	"github.com/AviadKorakin/ReactiveOrganizationalUnitsMicroservice-HTTP-2/services"
	httpSwagger "github.com/swaggo/http-swagger" // swagger handler
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// NewRouter sets up all routes and Swagger UI at /swagger/index.html
func NewRouter(client *mongo.Client) http.Handler {
	ur := repositories.NewUnitRepository(client)
	us := services.NewUnitService(ur)
	uc := controllers.NewUnitController(us)

	mux := http.NewServeMux()

	// Swagger UI
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	// Units collection endpoints
	mux.HandleFunc("/units", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			uc.Create(w, r)
		case http.MethodGet:
			uc.List(w, r)
		case http.MethodDelete:
			uc.DeleteAll(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Single-unit endpoints
	mux.HandleFunc("/units/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			uc.GetOne(w, r)
		case http.MethodPut:
			uc.Update(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
