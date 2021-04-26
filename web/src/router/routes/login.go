package routes

import (
	"net/http"
	"web/src/views"
)

var loginRoutes = []Route{
	{
		URI:          "/",
		Method:       http.MethodGet,
		Function:     views.LoginPage,
		AuthRequired: false,
	},
	{
		URI:          "/login",
		Method:       http.MethodGet,
		Function:     views.LoginPage,
		AuthRequired: false,
	},
}
