package handlers

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"net/http"
	"portfolio/web/components"
)

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {

	err := Page("login page", createLoginBody(w, r)).Render(w)
	if err != nil {
		return
	}

}

func createLoginBody(w http.ResponseWriter, r *http.Request) g.Node {

	return Body(components.Login())
}
