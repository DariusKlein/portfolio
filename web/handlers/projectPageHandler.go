package handlers

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"net/http"
	"portfolio/web/components"
	"portfolio/web/services"
)

func ProjectPageHandler(w http.ResponseWriter, r *http.Request) {

	err := Page("Project page", createProjectBody(w, r)).Render(w)
	if err != nil {
		return
	}

}

func createProjectBody(w http.ResponseWriter, r *http.Request) g.Node {

	return Body(
		components.ProjectList(services.ReadProjectsJson()),
	)

}
