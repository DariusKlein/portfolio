package api

import (
	"net/http"
	"portfolio/api/handler"
)

func ApiRoutes(mux **http.ServeMux) {
	m := *mux
	// Register the routes and webHandler
	m.HandleFunc("/api/", handler.CatchAllHandler)
	m.HandleFunc("POST /api/user", handler.CreateUser)
	m.HandleFunc("GET /api/user/{id}", handler.GetUser)
	m.HandleFunc("POST /api/login", handler.Login)
}
