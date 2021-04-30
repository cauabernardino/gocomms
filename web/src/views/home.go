package views

import (
	"fmt"
	"net/http"
	"strconv"
	"web/src/config"
	"web/src/controllers"
	"web/src/models"
	"web/src/responses"
	"web/src/utils"
)

// HomePage handles the loading of home page
func HomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)
	response, err := controllers.UserAuthenticatedRequest(r, http.MethodGet, url, nil)
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusInternalServerError,
			responses.APIError{Error: err.Error()},
		)
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleAPIStatusCodeError(w, response)
		return
	}

	posts, err := controllers.GetPosts(w, response)
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusUnprocessableEntity,
			responses.APIError{Error: err.Error()},
		)
		return
	}

	// Check for loading the custom buttons (edit, delete, etc)
	cookie, _ := utils.CheckCookie(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	RenderTemplate(
		w,
		"home.html",
		struct {
			Posts  []models.Post
			UserID uint64
		}{
			Posts:  posts,
			UserID: userID,
		},
	)
}
