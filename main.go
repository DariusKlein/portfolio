package main

import "net/http"

func main() {

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	//connect to database and migrate
	DB()

	// Register the routes and handlers
	mux.HandleFunc("/", catchAllHandler)
	mux.HandleFunc("/test", testHandler)
	mux.HandleFunc("/test2", test2Handler)

	// Run the server
	http.ListenAndServe(":4002", mux)
}

func catchAllHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusGone)
	_, err := w.Write([]byte("Bad endpoint"))
	if err != nil {
		InternalServerErrorHandler(w, r)
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("test"))
	if err != nil {
		InternalServerErrorHandler(w, r)
	}
}

func test2Handler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("test2"))
	if err != nil {
		InternalServerErrorHandler(w, r)
	}
}
