package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	api2 "portfolio/api"
	"portfolio/database"
)

func main() {
	// load .env in runtime environment
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env not found: %v", err)
		return
	}
	print("test")
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching webHandler
	mux := http.NewServeMux()

	api2.WebRoutes(&mux)

	api2.ApiRoutes(&mux)

	//connect to database and migrate
	database.DB()

	// Run the server
	http.ListenAndServe(":4002", mux)
}
