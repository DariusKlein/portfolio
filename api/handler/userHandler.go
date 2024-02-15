package handler

import (
	"context"
	"net/http"
	"portfolio-backend/database/query"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("test2"))
	if err != nil {
		InternalServerErrorHandler(w, r)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user, err := query.GetUser(context.Background())
	if err != nil {
		return
	}
	_, err = w.Write([]byte(r.PathValue(user.Name)))
	if err != nil {
		InternalServerErrorHandler(w, r)
	}
}
