package routes

import (
	"net/http"
	"web/src/controllers"
	"web/src/views"
)

var postRoutes = []Route{
	{
		URI:          "/posts",
		Method:       http.MethodPost,
		Function:     controllers.CreatePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}/like",
		Method:       http.MethodPost,
		Function:     controllers.LikePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}/unlike",
		Method:       http.MethodPost,
		Function:     controllers.UnlikePost,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}/edit",
		Method:       http.MethodGet,
		Function:     views.EditPostPage,
		AuthRequired: true,
	},
	{
		URI:          "/posts/{postID}",
		Method:       http.MethodPut,
		Function:     controllers.EditPost,
		AuthRequired: true,
	},
}
