package handler

import "net/http"

func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusGone)
	_, err := w.Write([]byte("Bad endpoint"))
	if err != nil {
		InternalServerErrorHandler(w)
	}
}
