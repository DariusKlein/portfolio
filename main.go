package main

import "net/http"

func main() {

	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.HandleFunc("GET /", test)

	// Run the server
	http.ListenAndServe(":4002", mux)
}

func test(w http.ResponseWriter, r *http.Request) {
	println("test123")
	w.Write([]byte("test"))
}
