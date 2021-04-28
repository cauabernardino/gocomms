package middlewares

import (
	"log"
	"net/http"
	"web/src/utils"
)

// Logger writes the request parameters to the STDOUT
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s\t%s\t%s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// AuthenticateSession verifies if there's a valid cookie present
// in the user browser
func AuthenticateSession(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := utils.CheckCookie(r); err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next(w, r)
	}
}
