package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"portfolio/api/service/jwt"
	"portfolio/api/service/parse"
	"portfolio/database/ent"
	"portfolio/database/query"
	"strconv"
)

func CreateProjectHandler(w http.ResponseWriter, r *http.Request) {

	id, _, err := jwt.VerifyUser(r)
	if err != nil {
		UnauthorizedHandler(w)
		return
	}

	var p *ent.Project

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}

	err = query.CreateProject(context.Background(), *p, id)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}
}

func UpdateProjectHandler(w http.ResponseWriter, r *http.Request) {

	_, _, err := jwt.VerifyUser(r)
	if err != nil {
		UnauthorizedHandler(w)
		return
	}

	var p *ent.Project

	projectID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		BadRequestHandler(w)
	}

	p, err = query.GetFullProject(context.Background(), projectID)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		InternalServerErrorHandler(w, err)
		return
	}

	err = query.UpdateProject(context.Background(), *p, projectID)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}
}

func UpdateProjectsHandler(w http.ResponseWriter, r *http.Request) {

	isHtmx := r.Header.Get("HX-Request")

	_, _, err := jwt.VerifyUser(r)
	if err != nil {
		UnauthorizedHandler(w)
		return
	}

	var newProjects []*ent.Project
	var p []*ent.Project

	if isHtmx == "true" {
		p = parse.ParseProjectInput(r)
	} else {

		p, err = query.GetFullProjects(context.Background())
		if err != nil {
			UnprocessableEntityHandler(w, err)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&newProjects)
		if err != nil {
			InternalServerErrorHandler(w, err)
			return
		}

		for _, project := range p {
			for _, newProject := range newProjects {
				if project.ID == newProject.ID {
					// todo add update from api
				}
			}
		}
	}

	for _, project := range p {
		err = query.UpdateProject(context.Background(), *project, project.ID)
		if err != nil {
			UnprocessableEntityHandler(w, err)
			return
		}
	}

}

func GetProjectHandler(w http.ResponseWriter, r *http.Request) {

	projectID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		BadRequestHandler(w)
	}

	p, err := query.GetProject(context.Background(), projectID)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		return
	}
}

func GetProjectsHandler(w http.ResponseWriter, r *http.Request) {

	p, err := query.GetProjects(context.Background())
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		return
	}
}
