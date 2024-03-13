package api

import (
	"net/http"
	handler2 "portfolio/api/handler"
)

func ApiRoutes(mux **http.ServeMux) {
	m := *mux
	// Register the routes and webHandler
	m.HandleFunc("/api/", handler2.CatchAllHandler)
	m.HandleFunc("POST /api/user", handler2.CreateUser)
	m.HandleFunc("GET /api/user/{id}", handler2.GetUser)
}
