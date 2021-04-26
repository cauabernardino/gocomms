package router

import (
	"web/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns a new router with all configured routes
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Configure(r)
}
