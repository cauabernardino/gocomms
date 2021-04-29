package routes

import (
	"net/http"
	"web/src/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

// Configure handles the configuration of the routes to the router
func Configure(router *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, homeRoute)
	routes = append(routes, userRoutes...)
	routes = append(routes, postRoutes...)

	for _, route := range routes {

		if route.AuthRequired {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(
					middlewares.AuthenticateSession(route.Function),
				),
			).Methods(route.Method)
		} else {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(
					route.Function,
				),
			).Methods(route.Method)
		}

	}

	// Configuration to handle static files
	fileServer := http.FileServer(http.Dir("./static/"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	return router
}
