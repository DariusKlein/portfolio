package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"portfolio/database"
	"portfolio/routes"
)

func main() {
	// load .env in runtime environment
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env not found: %v", err)
		return
	}

	//init web routes
	web := routes.WebRoutes()
	//init api routes
	api := routes.ApiRoutes()

	//connect to database and migrate
	database.DB()

	// Run the server
	err = http.ListenAndServe(":4001", web)
	if err != nil {
		log.Fatalf("web failed to start: %v", err)
		return
	}
	err = http.ListenAndServe(":4002", api)
	if err != nil {
		log.Fatalf("api failed to start: %v", err)
		return
	}
}
