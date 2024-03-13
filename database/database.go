package database

import (
	"context"
	"log"
	"os"
	"portfolio/database/ent"

	_ "github.com/go-sql-driver/mysql"
)

var Client *ent.Client

func DB() {
	err := *new(error)
	Client, err = ent.Open("mysql", os.Getenv("DB_STRING"))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	// Run the auto migration tool.
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
