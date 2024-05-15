package webHandler

import (
	"html/template"
	"net/http"
	"portfolio/types"
)

func InitProjectpage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./templates/index.html"))

	userNames := map[string][]types.Username{
		"Names": {
			{Name: "The Godfather"},
			{Name: "Blade Runner"},
			{Name: "The Thing"},
		},
	}

	err := tmpl.Execute(w, userNames)
	if err != nil {
		_, err := w.Write([]byte("failed to load"))
		if err != nil {
			return
		}
		return
	}
}
