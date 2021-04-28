package views

import (
	"fmt"
	"net/http"
	"web/src/config"
	"web/src/controllers"
	"web/src/responses"
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
		responses.HandleStatusCodeError(w, response)
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

	RenderTemplate(w, "home.html", posts)
}
