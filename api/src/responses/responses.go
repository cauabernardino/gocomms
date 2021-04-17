package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// ReturnJSON returns a JSON response for the request
func ReturnJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// ReturnError throws the errors in JSON format
func ReturnError(w http.ResponseWriter, statusCode int, err error) {

	ReturnJSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})

}
