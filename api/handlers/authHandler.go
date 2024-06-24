package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"portfolio/api/service/bcrypt"
	"portfolio/api/service/jwt"
	"portfolio/database/ent"
	"portfolio/database/query"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var u *ent.User

	isHtmx := r.Header.Get("HX-Request")

	if isHtmx == "true" {
		u = &ent.User{
			Email:    r.PostFormValue("email"),
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

		jwtToken := jwt.CreateUserJWT(User.Name, User.ID, string(User.Role))

		if jwtToken != "" {

			cookie := &http.Cookie{Name: "jwt",
				Value: jwtToken,
				//HttpOnly: true,
				//Secure:   true,
				SameSite: http.SameSiteLaxMode,
				Expires:  time.Now().Add(24 * time.Hour),
			}

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

func CanEdit(w http.ResponseWriter, r *http.Request) {
	jwtCookie, _ := r.Cookie("jwt")

	if jwtCookie != nil {
		_, audience, err := jwt.VerifyJWT(jwtCookie.Value)
		if err != nil {
			InternalServerErrorHandler(w, err)
			return
		}
		if audience == "owner" || audience == "visitor" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("<button class=\"button is-link\">Edit</button>"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(""))
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
