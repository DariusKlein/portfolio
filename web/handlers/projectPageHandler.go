package handlers

import (
	"context"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
	"net/http"
	"portfolio/database/query"
	"portfolio/web/components"
)

func ProjectPageHandler(w http.ResponseWriter, r *http.Request) {

	err := Page("Project page", createProjectBody(w, r)).Render(w)
	if err != nil {
		return
	}

}

func createProjectBody(w http.ResponseWriter, r *http.Request) g.Node {

	projects, err := query.GetProjects(context.Background())
	if err != nil {
		return nil
	}

	return Body(
		components.ProjectList(projects),
	)

}
