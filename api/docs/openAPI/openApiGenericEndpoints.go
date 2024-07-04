package main

import (
	"github.com/a-h/respond"
	"github.com/a-h/rest"
	"net/http"
)

func RegisterGenericEndpoints() {

	api.Get("/check").
		HasDescription("check for user jwt cookie").
		HasResponseModel(http.StatusOK, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]())

	api.Get("/htmx/canEdit").
		HasDescription("check if user is allowed to edit").
		HasResponseModel(http.StatusOK, rest.ModelOf[string]())
}
