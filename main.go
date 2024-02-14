package main

import (
	"net/http"
	"portfolio-backend/database"

	"portfolio-backend/api/handler"
)

func main() {

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	//connect to database and migrate
	database.DB()

	// Register the routes and handlers
	mux.HandleFunc("/", handler.CatchAllHandler)
	mux.HandleFunc("/test", handler.TestHandler)
	mux.HandleFunc("/test2", handler.Test2Handler)

	// Run the server
	http.ListenAndServe(":4002", mux)
}
