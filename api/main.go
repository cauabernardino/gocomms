package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	config.LoadEnvs()

	fmt.Println("Running the API!")
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}
