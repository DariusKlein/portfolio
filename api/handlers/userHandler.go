package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"portfolio/api/service/bcrypt"
	"portfolio/api/service/validate"
	"portfolio/database/ent"
	"portfolio/database/ent/user"
	"portfolio/database/query"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var u *ent.User

	isHtmx := r.Header.Get("HX-Request")

	if isHtmx == "true" {
		u = &ent.User{
			Name: r.PostFormValue("name"),
			Role: user.Role(r.PostFormValue("role")),
		}
	} else {
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			InternalServerErrorHandler(w, err)
		}
	}

	if !validate.UserIsValid(u) {
		BadRequestHandler(w)
		return
	}

	//hash password
	u.Password, _ = bcrypt.HashPassword(u.Password)

	err := query.CreateUser(context.Background(), *u)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("user created"))

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	userID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		BadRequestHandler(w)
	}

	User, err := query.GetUser(context.Background(), userID)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(User)
	if err != nil {
		return
	}
}
