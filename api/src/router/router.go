package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate will return the configured routes.
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Configure(r)
}
