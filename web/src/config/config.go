package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL is self explained
	APIURL = ""
	// Port of application
	Port = 0
	// HashKey for authenticate the cookie
	HashKey []byte
	// BlockKey for encrypt the cookie
	BlockKey []byte
)

// LoadEnvs loads all environment variables
func LoadEnvs() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
