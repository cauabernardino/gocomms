package controllers

import "net/http"

// CreateUser handles the creation of new user in the platform.
func CreateUser(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Creating user!"))
}

// SearchUsers handles the searching for all users of the platform.
func SearchUsers(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Searching for all users!"))
}

// SearchUser handles the searching for a specific user.
func SearchUser(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Searching for a specific user!"))
}

// UpdateUser handles the edition and/or updating of an user.
func UpdateUser(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Updating an user!"))
}

// DeleteUser handles the deletion of an user.
func DeleteUser(w http.ResponseWriter, r *http.Response) {
	w.Write([]byte("Deleting an user!"))
}
