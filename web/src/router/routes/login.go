package routes

import (
	"net/http"
	"web/src/controllers"
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
	{
		URI:          "/login",
		Method:       http.MethodPost,
		Function:     controllers.UserLogin,
		AuthRequired: false,
	},
}
