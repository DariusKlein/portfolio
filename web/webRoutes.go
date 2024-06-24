package web

import (
	"net/http"
	"portfolio/web/handlers"
)

func WebRoutes() *http.ServeMux {
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching webHandler
	mux := http.NewServeMux()

	// Register the routes and webHandler
	mux.HandleFunc("/", handlers.HomePageHandler)
	mux.HandleFunc("/projects", handlers.ProjectPageHandler)
	mux.HandleFunc("/projects/edit", handlers.CreateProjectEditBody)
	mux.HandleFunc("/about", handlers.AboutPageHandler)
	mux.HandleFunc("/login", handlers.LoginPageHandler)

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets"))))
	return mux
}
