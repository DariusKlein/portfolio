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
	mux.HandleFunc("GET /{$}", handlers.InitHomepage)
	mux.HandleFunc("GET /projecten/{$}", handlers.InitProjectpage)
	mux.HandleFunc("GET /about/{$}", handlers.InitAboutpage)
	mux.HandleFunc("GET /login/{$}", handlers.InitLoginpage)

	mux.HandleFunc("POST /theme", handlers.UpdateTheme)

	mux.HandleFunc("GET /test", handlers.Test)

	return mux
}
