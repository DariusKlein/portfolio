package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"portfolio/api/service/bcrypt"
	"portfolio/api/service/jwt"
	"portfolio/database/ent"
	"portfolio/database/query"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var u *ent.User

	isHtmx := r.Header.Get("HX-Request")

	if isHtmx == "true" {
		u = &ent.User{
			Name:     r.PostFormValue("name"),
			Password: r.PostFormValue("password"),
		}
	} else {
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			InternalServerErrorHandler(w, err)
		}
	}

	User, err := query.GetLogin(context.Background(), u)
	if err != nil {
		UnprocessableEntityHandler(w, err)
		return
	}

	if bcrypt.CheckPasswordHash(u.Password, User.Password) {

		jwtToken := jwt.CreateUserJWT(u.Name, u.ID, string(u.Role))

		if jwtToken != "" {
			w.Header().Set("HX-Location", "/")

			cookie := &http.Cookie{Name: "jwt", Value: jwtToken, HttpOnly: true, Secure: true, SameSite: http.SameSiteStrictMode}
			http.SetCookie(w, cookie)

			w.WriteHeader(http.StatusOK)
			_, err = w.Write([]byte("login success"))
			return
		} else {
			InternalServerErrorHandler(w, err)
			return
		}

	} else {
		UnauthorizedHandler(w)
		return
	}
}
