package utils

import (
	"net/http"
	"web/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

// ConfigureCookies sets the hashed keys for creating cookies
func ConfigureCookies() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

// SaveCookie encode the cookie and save it to the user browser
func SaveCookie(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	encodedData, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    encodedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

// CheckCookie returns the values stored in user's cookie
func CheckCookie(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	cookieValues := make(map[string]string)
	if err = s.Decode("data", cookie.Value, &cookieValues); err != nil {
		return nil, err
	}

	return cookieValues, nil
}
