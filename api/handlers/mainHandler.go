package handlers

import (
	"net/http"
	"portfolio/api/service/jwt"
	"strconv"
)

func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusGone)
	_, err := w.Write([]byte("Bad endpoint"))
	if err != nil {
		InternalServerErrorHandler(w, err)
	}
}

func CheckRoleHandler(w http.ResponseWriter, r *http.Request) {
	jwtCookie, _ := r.Cookie("jwt")

	if jwtCookie != nil {
		uid, audience, err := jwt.VerifyJWT(jwtCookie.Value)
		if err != nil {
			InternalServerErrorHandler(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("id: " + strconv.Itoa(uid) + "\naudience: " + audience))
		return
	}
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write([]byte("Cookie not found"))
}
