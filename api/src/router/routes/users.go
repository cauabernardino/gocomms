package routes

import "net/http"

var userRoutes = []Route{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: func(rw http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Function: func(rw http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodGet,
		Function: func(rw http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodPut,
		Function: func(rw http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodDelete,
		Function: func(rw http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
}
