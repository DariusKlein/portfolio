package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"portfolio/api/service/jwt"
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
