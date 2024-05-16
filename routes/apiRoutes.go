package routes

import (
	"net/http"
	"portfolio/api/handlers"
)

func ApiRoutes() *http.ServeMux {
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching webHandler
	mux := http.NewServeMux()

	// Register the routes and webHandler
	mux.HandleFunc("/api/", handlers.CatchAllHandler)
	mux.HandleFunc("POST /api/user", handlers.CreateUser)
	mux.HandleFunc("GET /api/user/{id}", handlers.GetUser)
	mux.HandleFunc("POST /api/login", handlers.Login)

	return mux
}
