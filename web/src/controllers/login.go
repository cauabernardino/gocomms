package controllers

import "net/http"

// LoginPage handles the loading of login page
func LoginPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login page!"))
}