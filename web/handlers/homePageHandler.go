package handlers

import (
	g "github.com/maragudk/gomponents"
	"net/http"
	"portfolio/web/components"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {

	err := Page("Homepage", createBody(w, r)).Render(w)
	if err != nil {
		return
	}

}

func createBody(w http.ResponseWriter, r *http.Request) []g.Node {

	return []g.Node{
		components.Navbar(r.URL.Path, []components.PageLink{
			{Path: "/contact", Name: "Contact"},
			{Path: "/about", Name: "About"},
		}),
	}
}
