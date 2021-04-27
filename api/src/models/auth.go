package models

// AuthData creates structure for authentication data response
type AuthData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
