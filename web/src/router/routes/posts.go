package routes

import (
	"net/http"
	"web/src/controllers"
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
}
