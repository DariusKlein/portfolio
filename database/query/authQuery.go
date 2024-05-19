package query

import (
	"context"
	"fmt"
	"log"
	"portfolio/database"
	"portfolio/database/ent"
	"portfolio/database/ent/user"
)

func GetLogin(ctx context.Context, U *ent.User) (*ent.User, error) {
	u, err := database.Client.User.
		Query().
		Where(user.Name(U.Name)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
