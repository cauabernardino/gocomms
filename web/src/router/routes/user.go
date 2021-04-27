package routes

import (
	"net/http"
	"web/src/controllers"
	"web/src/views"
)

var userRoutes = []Route{
	{
		URI:          "/signup",
		Method:       http.MethodGet,
		Function:     views.LoadSignUpPage,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
}
