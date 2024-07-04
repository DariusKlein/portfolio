package query

import (
	"context"
	"fmt"
	"log"
	"portfolio/database"
	"portfolio/database/ent"
	"portfolio/database/ent/project"
)

func CreateProject(ctx context.Context, project ent.Project, id int) error {

	_, err := database.Client.Project.
		Create().
		SetName(project.Name).
		SetDescription(project.Description).
		SetURL(project.URL).
		SetDocURL(project.DocURL).
		SetImageURL(project.ImageURL).
		AddUserIDs(id).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create project: %w", err)
	}
	log.Println("project created: ")
	return err
}

func UpdateProject(ctx context.Context, project ent.Project, projectID int) error {

	_, err := database.Client.Project.
		UpdateOneID(projectID).
		SetName(project.Name).
		SetDescription(project.Description).
		SetURL(project.URL).
		SetDocURL(project.DocURL).
		SetImageURL(project.ImageURL).
		AddUsers(project.Edges.Users...).
		AddTeams(project.Edges.Teams...).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to update project: %w", err)
	}
	log.Println("project updated: ")
	return err
}

func GetProject(ctx context.Context, projectID int) (*ent.Project, error) {

	p, err := database.Client.Project.Get(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}
	return p, err
}

func GetFullProject(ctx context.Context, projectID int) (*ent.Project, error) {

	p, err := database.Client.Project.
		Query().
		WithUsers().
		WithTeams().
		Where(project.ID(projectID)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}
	return p, err
}

func GetFullProjects(ctx context.Context) ([]*ent.Project, error) {

	p, err := database.Client.Project.
		Query().
		WithUsers().
		WithTeams().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get project: %w", err)
	}
	return p, err
}

func GetProjects(ctx context.Context) (ent.Projects, error) {

	p, err := database.Client.Project.Query().
		Order(ent.Asc(project.FieldID)).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get projects: %w", err)
	}
	return p, err
}
