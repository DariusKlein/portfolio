package web

import (
	"net/http"
	handlers2 "portfolio/api/handlers"
)

func WebRoutes() *http.ServeMux {
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching webHandler
	mux := http.NewServeMux()

	// Register the routes and webHandler
	mux.HandleFunc("/", handlers2.CatchAllHandler)

	return mux
}
