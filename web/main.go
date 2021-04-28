package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/config"
	"web/src/router"
	"web/src/utils"
	"web/src/views"
)

func main() {

	config.LoadEnvs()
	utils.ConfigureCookies()
	views.LoadTemplates()

	r := router.Generate()

	fmt.Printf("Listening to port %d...\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
