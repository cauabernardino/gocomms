package tools

import (
	"encoding/hex"
	"fmt"

	"github.com/gorilla/securecookie"
)

// GenerateKeys helps to generate the Hash and Block keys
func GenerateKeys() {
	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hashKey)

	blockKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(blockKey)
}
