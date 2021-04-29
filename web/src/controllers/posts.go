package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"web/src/config"
	"web/src/models"
	"web/src/responses"
)

// CreatePost handles the creation of a post request to the API
func CreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusBadRequest,
			responses.APIError{Error: err.Error()},
		)
		return
	}

	url := fmt.Sprintf("%s/posts", config.APIURL)
	response, err := UserAuthenticatedRequest(
		r,
		http.MethodPost,
		url,
		bytes.NewBuffer(post),
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

	// Check if the status code returned from API is an error
	if response.StatusCode >= 400 {
		responses.HandleStatusCodeError(w, response)
		return
	}

	responses.ReturnJSON(w, response.StatusCode, nil)
}

// GetPosts decodes the and returns the posts in JSON format
func GetPosts(w http.ResponseWriter, response *http.Response) ([]models.Post, error) {
	var posts []models.Post
	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		return nil, err
	}
	return posts, nil
}
