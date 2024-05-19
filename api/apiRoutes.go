package api

import (
	"net/http"
	"portfolio/api/handlers"
)

func ApiRoutes() *http.ServeMux {
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching webHandler
	mux := http.NewServeMux()

	// routes
	mux.HandleFunc("/", handlers.CatchAllHandler)
	mux.HandleFunc("GET /check", handlers.CheckRoleHandler)

	//user
	mux.HandleFunc("GET /user/{id}", handlers.GetUser)

	//auth
	mux.HandleFunc("POST /login", handlers.Login)
	mux.HandleFunc("POST /register", handlers.CreateUser)

	return mux
}
