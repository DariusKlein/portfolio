package handlers

import (
	"html/template"
	"net/http"
)

func InitAboutpage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles(
		"./templates/about/index.html",
		"./templates/navbar/index.html",
		"./templates/themeSelector/index.html",
	))

	err := tmpl.Execute(w, nil)
	if err != nil {
		_, err := w.Write([]byte("failed to load"))
		if err != nil {
			return
		}
		return
	}
}
