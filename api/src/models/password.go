package models

// Password represents the reset password request format
type Password struct {
	New     string `json:"new"`
	Current string `json:"current"`
}
