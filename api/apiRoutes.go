package api

import (
	"net/http"
	"portfolio/api/handlers"
)

func ApiRoutes() *http.ServeMux {
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching webHandler
	mux := http.NewServeMux()

	// Register the routes and webHandler
	mux.HandleFunc("/", handlers.CatchAllHandler)
	mux.HandleFunc("POST /register", handlers.CreateUser)
	mux.HandleFunc("GET /user/{id}", handlers.GetUser)
	mux.HandleFunc("POST /login", handlers.Login)

	return mux
}
