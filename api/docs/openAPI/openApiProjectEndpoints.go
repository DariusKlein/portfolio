package main

import (
	"github.com/a-h/respond"
	"github.com/a-h/rest"
	"net/http"
	"portfolio/database/ent"
)

func RegisterProjectEndpoints() {
	api.Post("/project").
		HasDescription("Create project").
		HasResponseModel(http.StatusOK, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

	api.Patch("/project/{id}").
		HasPathParameter("id", rest.PathParam{
			Description: "id of the project",
			Regexp:      `\d+`,
		}).
		HasDescription("Update project by id").
		HasRequestModel(rest.ModelOf[ent.Project]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[string]()).
		HasResponseModel(http.StatusBadRequest, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

	api.Patch("/projects").
		HasDescription("Update projects WIP").
		HasResponseModel(http.StatusOK, rest.ModelOf[[]ent.Project]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

	api.Get("/project/{id}").
		HasPathParameter("id", rest.PathParam{
			Description: "id of the project",
			Regexp:      `\d+`,
		}).
		HasDescription("Get project by id").
		HasResponseModel(http.StatusOK, rest.ModelOf[ent.Project]()).
		HasResponseModel(http.StatusBadRequest, rest.ModelOf[string]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]())

	api.Get("/projects").
		HasDescription("Get projects").
		HasResponseModel(http.StatusOK, rest.ModelOf[[]ent.Project]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]())
}
