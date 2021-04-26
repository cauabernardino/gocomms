package main

import (
	"fmt"
	"log"
	"net/http"
	"web/src/router"
	"web/src/views"
)

func main() {
	views.LoadTemplates()
	r := router.Generate()

	fmt.Println("Listening in port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
