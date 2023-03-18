package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MarynaMarkova/GO_HELLO_WORLD/pkg/config"
	"github.com/MarynaMarkova/GO_HELLO_WORLD/pkg/handlers"
)

const portNumber = ":8080"


// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.createTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

//To run this program: go run ./cmd/web/.