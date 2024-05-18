package services

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"portfolio/web/types"
)

func ReadProjectsJson() []types.Project {
	// Open our jsonFile
	jsonFile, err := os.Open("web/assets/json/projects.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var projects types.Projects

	json.Unmarshal(byteValue, &projects)

	return projects.Projects
}
