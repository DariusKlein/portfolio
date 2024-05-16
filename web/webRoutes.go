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

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets"))))
	return mux
}
