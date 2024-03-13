package api

import (
	"net/http"
	"portfolio/api/webHandler"
)

func WebRoutes(mux **http.ServeMux) {
	m := *mux
	// Register the routes and webHandler
	m.HandleFunc("GET /{$}", webHandler.InitHomepage)
}
