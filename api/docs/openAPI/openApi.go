package main

import (
	"encoding/json"
	"github.com/a-h/rest"
	"log"
	"os"
)

var api *rest.API

func main() {
	// Configure the models.
	api = rest.NewAPI("portfolio")
	api.StripPkgPaths = []string{"github.com/a-h/rest/example", "github.com/a-h/respond"}

	// register models
	RegisterModels()
	// register endpoints
	RegisterGenericEndpoints()
	RegisterUserEndpoints()
	RegisterProjectEndpoints()

	// Create the specification.
	spec, err := api.Spec()
	if err != nil {
		log.Fatalf("failed to create spec: %v", err)
	}

	// create file
	file, _ := os.OpenFile("common/docs/openApi.json", os.O_CREATE, os.ModePerm)
	defer file.Close()

	// Write to file
	enc := json.NewEncoder(file)
	enc.SetIndent("", " ")
	enc.Encode(spec)
}
