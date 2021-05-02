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
		Function:     views.SignUpPage,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     views.UsersPage,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}",
		Method:       http.MethodGet,
		Function:     views.ProfilePage,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/unfollow",
		Method:       http.MethodPost,
		Function:     controllers.UnfollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userID}/follow",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/profile",
		Method:       http.MethodGet,
		Function:     views.LoggedUserProfilePage,
		AuthRequired: true,
	},
	{
		URI:          "/edit-profile",
		Method:       http.MethodGet,
		Function:     views.EditProfilePage,
		AuthRequired: true,
	},
	{
		URI:          "/edit-profile",
		Method:       http.MethodPut,
		Function:     controllers.EditProfile,
		AuthRequired: true,
	},
	{
		URI:          "/change-password",
		Method:       http.MethodGet,
		Function:     views.ChangePasswordPage,
		AuthRequired: true,
	},
	{
		URI:          "/change-password",
		Method:       http.MethodPost,
		Function:     controllers.ChangePassword,
		AuthRequired: true,
	},
}
