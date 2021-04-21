package middlewares

import (
	"api/src/responses"
	"api/src/utils"
	"log"
	"net/http"
)

// Authenticate verifies if the user making the request is authenticated
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := utils.ValidateToken(r); err != nil {
			responses.ReturnError(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}

// Logger writes the request information to the STDOUT
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}
