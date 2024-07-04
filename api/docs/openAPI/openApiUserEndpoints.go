package main

import (
	"github.com/a-h/respond"
	"github.com/a-h/rest"
	"net/http"
	"portfolio/api/types"
)

func RegisterUserEndpoints() {
	api.Get("/user/{uid}").
		HasPathParameter("id", rest.PathParam{
			Description: "id of the user",
			Regexp:      `\d+`,
		}).
		HasDescription("Get user by uid.").
		HasResponseModel(http.StatusOK, rest.ModelOf[types.User]()).
		HasResponseModel(http.StatusBadRequest, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]())

	api.Post("/register").
		HasDescription("Register.").
		HasRequestModel(rest.ModelOf[types.User]()).
		HasResponseModel(http.StatusCreated, rest.ModelOf[string]()).
		HasResponseModel(http.StatusBadRequest, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]())

	api.Post("/login").
		HasDescription("Login.").
		HasRequestModel(rest.ModelOf[types.LoginUser]()).
		HasResponseModel(http.StatusOK, rest.ModelOf[string]()).
		HasResponseModel(http.StatusInternalServerError, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnprocessableEntity, rest.ModelOf[respond.Error]()).
		HasResponseModel(http.StatusUnauthorized, rest.ModelOf[string]())

}
