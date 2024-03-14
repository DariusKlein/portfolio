package webHandler

import (
	"html/template"
	"net/http"
)

func InitHomepage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles(
		"./templates/homepage/index.html",
		"./templates/navbar/index.html",
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

func Test(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles(
		"./templates/homepage/index.html",
		"./templates/navbar/index.html",
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
