package handler

import (
	"net/http"
)

func CatchAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusGone)
	_, err := w.Write([]byte("Bad endpoint"))
	if err != nil {
		InternalServerErrorHandler(w, r)
	}
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("test"))
	if err != nil {
		InternalServerErrorHandler(w, r)
	}
}

func Test2Handler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("test2"))
	if err != nil {
		InternalServerErrorHandler(w, r)
	}
}
