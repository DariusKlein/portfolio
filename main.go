package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"portfolio/api"
	"portfolio/database"
	"portfolio/web"
)

func main() {
	// load .env in runtime environment
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env not found: %v", err)
		return
	}

	//connect to database and migrate
	database.DB()

	//init web routes
	web := web.WebRoutes()
	// Run web server
	go http.ListenAndServe(":4001", web)

	//init api routes
	api := api.ApiRoutes()
	//run api server
	http.ListenAndServe(":4002", api)

}
