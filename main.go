package main

import (
	"net/http"
	"os"

	"meli-redirect-service/controller"
	"meli-redirect-service/repository"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	// Load the .env file
	err = godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// Connect to the database
	err = repository.Connect()
	if err != nil {
		panic("Error connecting to the database " + err.Error())
	}

	// Create the router
	r := mux.NewRouter()
	r.HandleFunc("/{key}", controller.RedirectHandler).Methods("GET")

	// Get the HTTP host and port
	httpHost := os.Getenv("HTTP_HOST")
	httpPort := os.Getenv("HTTP_PORT")

	// Start the HTTP server
	http.ListenAndServe(httpHost+":"+httpPort, r)
}
