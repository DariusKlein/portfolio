package handlers

import (
	g "github.com/maragudk/gomponents"
	"net/http"
	"portfolio/web/components"
	"portfolio/web/services"
)

func ProjectPageHandler(w http.ResponseWriter, r *http.Request) {

	err := Page("Homepage", createProjectBody(w, r)).Render(w)
	if err != nil {
		return
	}

}

func createProjectBody(w http.ResponseWriter, r *http.Request) []g.Node {

	return []g.Node{
		components.Navbar(),
		components.ProjectList(services.ReadProjectsJson()),
	}
}
