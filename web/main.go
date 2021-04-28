package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/config"
	"web/src/router"
	"web/src/views"
)

func main() {

	views.LoadTemplates()
	config.LoadEnvs()

	r := router.Generate()

	fmt.Printf("Listening to port %d...\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
