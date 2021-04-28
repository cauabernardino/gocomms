package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Error represents the error response from API
type APIError struct {
	Error string `json:"error"`
}

// ReturnJSON returns a response in JSON format for the request
func ReturnJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// HandleStatusCodeError returns the API status code error message in a JSON format
func HandleStatusCodeError(w http.ResponseWriter, r *http.Response) {
	var err APIError

	json.NewDecoder(r.Body).Decode(&err)
	ReturnJSON(w, r.StatusCode, err)
}
