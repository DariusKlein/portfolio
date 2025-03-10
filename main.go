package main

import (
	"github.com/rs/cors"
	"net/http"
	"portfolio/api"
	"portfolio/database"
	"portfolio/web"
)

func main() {
	//// load .env in runtime environment
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatalf(".env not found: %v", err)
	//	return
	//}

	//connect to database and migrate
	database.DB()

	//init web routes
	webMux := web.WebRoutes()
	// Run web server
	go http.ListenAndServe(":4000", cors.AllowAll().Handler(webMux))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4000", "https://*.dariusklein.nl", "https://*.portfolio.dariusklein.nl", "https://dariusklein.nl"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	//init api routes
	apiMux := api.ApiRoutes()
	//run api server
	http.ListenAndServe(":4001", c.Handler(apiMux))

}
