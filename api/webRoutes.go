package api

import (
	"net/http"
	"portfolio/api/webHandler"
)

func WebRoutes(mux **http.ServeMux) {
	m := *mux
	// Register the routes and webHandler
	m.HandleFunc("GET /{$}", webHandler.InitHomepage)
	m.HandleFunc("GET /projecten/{$}", webHandler.InitProjectpage)
	m.HandleFunc("POST /theme", webHandler.UpdateTheme)

	m.HandleFunc("GET /test", webHandler.Test)

	m.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./templates/assets"))))
}
