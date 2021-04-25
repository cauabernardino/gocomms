package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePost handles the creation of a new post in the platform
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.ExtractUserID(r)
	if err != nil {
		responses.ReturnError(w, http.StatusUnauthorized, err)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ReturnError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(reqBody, &post); err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}
	post.AuthorID = userID

	if err = post.Prepare(); err != nil {
		responses.ReturnError(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostsRepository(db)
	post.ID, err = postRepository.Create(post)
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ReturnJSON(w, http.StatusCreated, post)

}

// SearchPosts handles the searching for all posts
func SearchPosts(w http.ResponseWriter, r *http.Request) {}

// SearchPost handles the searching for a specific post
func SearchPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postID"], 10, 64)
	if err != nil {
		responses.ReturnError(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	postRepository := repositories.NewPostsRepository(db)
	post, err := postRepository.SearchID(postID)
	if err != nil {
		responses.ReturnError(w, http.StatusInternalServerError, err)
		return
	}

	responses.ReturnJSON(w, http.StatusOK, post)
}

// UpdatePost handles the edition/updating of a post
func UpdatePost(w http.ResponseWriter, r *http.Request) {}

// DeletePost handles the deletion of a post
func DeletePost(w http.ResponseWriter, r *http.Request) {}
