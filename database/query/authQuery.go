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

func GetLogin(ctx context.Context, U *types.LoginUser) (*ent.User, error) {
	u, err := database.Client.User.
		Query().
		Where(user.Email(U.Email)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}
