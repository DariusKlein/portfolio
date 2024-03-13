package main

import (
	"net/http"
	"portfolio/backend/api/handler"
	"portfolio/backend/database"
)

func main() {

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	//connect to database and migrate
	database.DB()

	// Register the routes and handlers
	mux.HandleFunc("/", handler.CatchAllHandler)
	mux.HandleFunc("POST /user", handler.CreateUser)
	mux.HandleFunc("GET /user/{id}", handler.GetUser)

	// Run the server
	http.ListenAndServe(":4002", mux)
}
