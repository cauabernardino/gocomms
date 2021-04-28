package views

import "net/http"

// HomePage handles the loading of home page
func HomePage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.html", nil)
}
