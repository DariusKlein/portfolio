package validate

import (
	"portfolio/database/ent"
)

func UserIsValid(u *ent.User) bool {
	if len(u.Name) > 0 &&
		len(u.Role) > 0 {
		return true
	}
	return false
}
