package handlers

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"net/http"
)

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {

	err := Page("About page", createAboutBody(w, r)).Render(w)
	if err != nil {
		return
	}

}

func createAboutBody(w http.ResponseWriter, r *http.Request) g.Node {

	return Body()
}
