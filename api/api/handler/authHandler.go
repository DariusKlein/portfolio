package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"portfolio_api/database/ent"
	"portfolio_api/database/ent/user"
	"portfolio_api/database/query"
	"portfolio_api/service/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	var u *ent.User

	isHtmx := r.Header.Get("HX-Request")

	if isHtmx == "true" {
		u = &ent.User{
			Name:     r.PostFormValue("name"),
			Password: r.PostFormValue("password"),
			Role:     user.Role(r.PostFormValue("role")),
		}
	} else {
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			InternalServerErrorHandler(w)
		}
	}

	User, err := query.GetLogin(context.Background(), u.Name)
	if err != nil {
		return
	}

	if bcrypt.CheckPasswordHash(u.Password, User.Password) {
		w.Header().Set("HX-Location", "/")
		return
	} else {

	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(User)
	if err != nil {
		return
	}
}
