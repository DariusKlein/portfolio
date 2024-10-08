package query

import (
	"context"
	"fmt"
	"log"
	"portfolio/api/types"
	"portfolio/database"
	"portfolio/database/ent"
	"portfolio/database/ent/user"
)

func GetUser(ctx context.Context, id int) (*ent.User, error) {

	u, err := database.Client.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func CreateUser(ctx context.Context, user types.RegisterUser) error {

	_, err := database.Client.User.
		Create().
		SetName(user.Name).
		SetEmail(user.Email).
		SetPassword(user.Password).
		SetRole("visitor").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	log.Println("user created: ")
	return err
}
