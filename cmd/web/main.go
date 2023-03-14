package main

import (
	"fmt"
	"net/http"

	"github.com/MarynaMarkova/Go-Learning/tree/main/First_Site_2/pkg/handlers"
)

const portNumber = ":8080"


// main is the main function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

//To run this program: go run ./cmd/web/.