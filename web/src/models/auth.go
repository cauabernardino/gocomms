package models

// AuthData represents the struct of authentication data sent by API
type AuthData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
