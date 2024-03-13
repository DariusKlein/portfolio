package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"portfolio/backend/database/ent"
	"portfolio/backend/database/query"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var u *ent.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		InternalServerErrorHandler(w, r)
	}

	err = query.CreateUser(context.Background(), *u)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("user created"))

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	userID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		BadRequestHandler(w, r)
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
