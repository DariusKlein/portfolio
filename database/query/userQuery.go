package query

import (
	"context"
	"fmt"
	"log"
	"portfolio-backend/database"
	"portfolio-backend/database/ent"
	"portfolio-backend/database/ent/user"
)

func GetUser(ctx context.Context) (*ent.User, error) {
	u, err := database.DBclient.User.
		Query().
		Where(user.Name("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
