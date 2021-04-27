package routes

import (
	"net/http"
	"web/src/views"
)

var userRoutes = []Route{
	{
		URI:          "/signup",
		Method:       http.MethodGet,
		Function:     views.LoadSignUpPage,
		AuthRequired: false,
	},
}
