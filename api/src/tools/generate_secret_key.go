package tools

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

// GenerateSecretKey helps the creation of project secret key for using in JWT
// and others encryptions
func GenerateSecretKey() {
	key := make([]byte, 64)

	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(key)

	fmt.Println(stringBase64)
}
