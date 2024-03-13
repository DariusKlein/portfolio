package api

import (
	"net/http"
	"portfolio/api/handler"
)

func WebRoutes(mux **http.ServeMux) {
	m := *mux
	// Register the routes and webHandler
	m.HandleFunc("GET /{$}", handler.InitHomepage)
}
