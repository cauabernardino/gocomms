package views

import (
	"fmt"
	"net/http"
	"strconv"
	"web/src/config"
	"web/src/controllers"
	"web/src/responses"

	"github.com/gorilla/mux"
)

// EditPostPage handles the loading of edit post page
func EditPostPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID, err := strconv.ParseUint(params["postID"], 10, 64)
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusBadRequest,
			responses.APIError{Error: err.Error()},
		)
		return
	}

	url := fmt.Sprintf("%s/posts/%d", config.APIURL, postID)
	response, err := controllers.UserAuthenticatedRequest(
		r,
		http.MethodGet,
		url,
		nil,
	)
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
	post, err := controllers.GetSinglePost(w, response)
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusUnprocessableEntity,
			responses.APIError{Error: err.Error()},
		)
		return
	}

	RenderTemplate(w, "edit-post.html", post)
}
