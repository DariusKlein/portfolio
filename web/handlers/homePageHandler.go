package handlers

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"net/http"
	"portfolio/web/components"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {

	err := Page("Homepage", createHomeBody(w, r)).Render(w)
	if err != nil {
		return
	}

}

func createHomeBody(w http.ResponseWriter, r *http.Request) []g.Node {

	return []g.Node{
		components.Navbar(),
		Main(
			ID("main"),
		),
	}
}
