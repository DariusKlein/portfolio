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
	mux.HandleFunc("GET /user/{id}", handlers.GetUserHandler)

	//auth
	mux.HandleFunc("POST /login", handlers.Login)
	mux.HandleFunc("POST /register", handlers.CreateUserHandler)
	mux.HandleFunc("GET /htmx/canEdit", handlers.CanEdit)

	//Project
	mux.HandleFunc("POST /project", handlers.CreateProjectHandler)
	mux.HandleFunc("PATCH /project/{id}", handlers.UpdateProjectHandler)
	mux.HandleFunc("PATCH /projects", handlers.UpdateProjectsHandler)
	mux.HandleFunc("GET /project/{id}", handlers.GetProjectHandler)
	mux.HandleFunc("GET /projects", handlers.GetProjectsHandler)

	return mux
}
