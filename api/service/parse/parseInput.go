package parse

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/url"
	"portfolio/database/ent"
	"portfolio/database/query"
	"strconv"
)

func ParseProjectInput(r *http.Request) []*ent.Project {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := url.ParseQuery(string(b))

	var projects []*ent.Project

	for x := range body["project_name"] {
		var p *ent.Project

		projectID, err := strconv.Atoi(body["project_id"][x])

		p, err = query.GetFullProject(context.Background(), projectID)
		if err != nil {
			break
		}
		p.Name = body["project_name"][x]
		p.URL = body["project_repo"][x]
		p.DocURL = body["project_docs"][x]
		p.Description = body["project_description"][x]
		projects = append(projects, p)
	}
	return projects
}
