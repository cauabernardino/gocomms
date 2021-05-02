package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"web/src/config"
	"web/src/models"
	"web/src/responses"
	"web/src/utils"

	"github.com/gorilla/mux"
)

// CreateUser handles the call to API for registering an user in database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"username": r.FormValue("username"),
		"password": r.FormValue("password"),
	})
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusBadRequest,
			responses.APIError{Error: err.Error()},
		)
		return
	}

	url := fmt.Sprintf("%s/users", config.APIURL)
	response, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(user),
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
		responses.HandleAPIStatusCodeError(w, response)
		return
	}

	responses.ReturnJSON(w, response.StatusCode, nil)
}

// UserAuthenticatedRequest handles the token insertion into request
func UserAuthenticatedRequest(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	// Creates the request
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	// Get cookie from browser and add it to request header
	cookie, _ := utils.CheckCookie(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	// Execution of request and catching response
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetAllUserInformation calls concurrently for the API to get all
// platform information of the user
func GetAllUserInformation(userID uint64, r *http.Request) (models.User, error) {
	userChannel := make(chan models.User)
	followersChannel := make(chan []models.User)
	followingChannel := make(chan []models.User)
	postsChannel := make(chan []models.Post)

	go GetUserData(userChannel, userID, r)
	go GetFollowersData(followersChannel, userID, r)
	go GetFollowingData(followingChannel, userID, r)
	go GetPostsData(postsChannel, userID, r)

	var (
		user      models.User
		followers []models.User
		following []models.User
		posts     []models.Post
	)

	for i := 0; i < 4; i++ {
		select {
		case userLoad := <-userChannel:
			if userLoad.ID == 0 {
				return models.User{}, errors.New("error on getting user data")
			}
			user = userLoad

		case followersLoad := <-followersChannel:
			if followersLoad == nil {
				return models.User{}, errors.New("error on getting followers")
			}
			followers = followersLoad

		case followingLoad := <-followingChannel:
			if followingLoad == nil {
				return models.User{}, errors.New("error on getting following")
			}
			following = followingLoad

		case postsLoad := <-postsChannel:
			if postsLoad == nil {
				return models.User{}, errors.New("error on getting user posts")
			}
			posts = postsLoad
		}
	}

	user.Followers = followers
	user.Following = following
	user.Posts = posts

	return user, nil
}

// GetUserData is a helper function to call API in users route
func GetUserData(channel chan models.User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, err := UserAuthenticatedRequest(
		r,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		channel <- models.User{}
		return
	}
	defer response.Body.Close()

	var user models.User
	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		channel <- models.User{}
		return
	}

	channel <- user
}

// GetFollowersData is a helper function to call API in followers route
func GetFollowersData(channel chan []models.User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)
	response, err := UserAuthenticatedRequest(
		r,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var followers []models.User
	if err = json.NewDecoder(response.Body).Decode(&followers); err != nil {
		channel <- nil
		return
	}

	if followers == nil {
		channel <- make([]models.User, 0)
		return
	}

	channel <- followers
}

// GetFollowingData is a helper function to call API in following route
func GetFollowingData(channel chan []models.User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)
	response, err := UserAuthenticatedRequest(
		r,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var following []models.User
	if err = json.NewDecoder(response.Body).Decode(&following); err != nil {
		channel <- nil
		return
	}

	if following == nil {
		channel <- make([]models.User, 0)
		return
	}

	channel <- following
}

// GetFollowingData is a helper function to call API in posts route
func GetPostsData(channel chan []models.Post, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.APIURL, userID)
	response, err := UserAuthenticatedRequest(
		r,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		channel <- nil
		return
	}
	defer response.Body.Close()

	var posts []models.Post
	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		channel <- nil
		return
	}

	if posts == nil {
		channel <- make([]models.Post, 0)
		return
	}

	channel <- posts
}

// CreateUser handles the call to API for unfollowing an user
func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusBadRequest,
			responses.APIError{Error: err.Error()},
		)
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollow", config.APIURL, userID)
	response, err := UserAuthenticatedRequest(
		r,
		http.MethodPost,
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

	responses.ReturnJSON(w, response.StatusCode, nil)
}

// FollowUser handles the call to API for following an user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userID"], 10, 64)
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusBadRequest,
			responses.APIError{Error: err.Error()},
		)
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.APIURL, userID)
	response, err := UserAuthenticatedRequest(
		r,
		http.MethodPost,
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

	responses.ReturnJSON(w, response.StatusCode, nil)
}

// EditProfile handles the call to API to edit the user data
func EditProfile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"username": r.FormValue("username"),
	})
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusBadRequest,
			responses.APIError{Error: err.Error()},
		)
		return
	}

	cookie, _ := utils.CheckCookie(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, err := UserAuthenticatedRequest(
		r,
		http.MethodPut,
		url,
		bytes.NewBuffer(user),
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

	responses.ReturnJSON(w, response.StatusCode, nil)
}

// ChangePassword handles the call to API to change user password
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	passwords, err := json.Marshal(map[string]string{
		"new":     r.FormValue("new"),
		"current": r.FormValue("current"),
	})
	if err != nil {
		responses.ReturnJSON(
			w,
			http.StatusBadRequest,
			responses.APIError{Error: err.Error()},
		)
		return
	}

	cookie, _ := utils.CheckCookie(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/reset", config.APIURL, userID)
	response, err := UserAuthenticatedRequest(
		r,
		http.MethodPost,
		url,
		bytes.NewBuffer(passwords),
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

	responses.ReturnJSON(w, response.StatusCode, nil)
}
