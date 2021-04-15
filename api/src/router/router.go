package router

import "github.com/gorilla/mux"

// Generate will return the configured routes.
func Generate() *mux.Router {
	return mux.NewRouter()
}
