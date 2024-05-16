package handlers

import "net/http"

func InternalServerErrorHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte("500 Internal Server Error"))
	if err != nil {
		return
	}
}

func NotFoundHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	_, err := w.Write([]byte("404 Not Found"))
	if err != nil {
		return
	}
}

func BadRequestHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte("400 Bad Request"))
	if err != nil {
		return
	}
}
