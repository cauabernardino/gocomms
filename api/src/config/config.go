package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DBConnectString is the statement to connect the application with the database
	DBConnectString = ""

	// Port where the API will run
	Port = 0

	// SECRET_KEY is the key used to sign the JWT token
	SECRET_KEY []byte
)

// LoadEnvs will initialize the environment variables for the application
func LoadEnvs() {

	// Load .env file
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Get Port and check port availability
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 5050
	}

	// Get DB related variables
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")

	DBConnectString = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
	)

	SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))
}
