package database

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
	"portfolio-backend/database/ent"

	_ "github.com/go-sql-driver/mysql"
)

func DB() {

	err := godotenv.Load()
	if err != nil {
		return
	}

	client, err := ent.Open("mysql", os.Getenv("DB_STRING"))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
