package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.SearchUsers,
		AuthRequired: false,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodGet,
		Function:     controllers.SearchUser,
		AuthRequired: false,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: false,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: false,
	},
}
