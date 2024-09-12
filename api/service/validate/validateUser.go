package validate

import (
	"portfolio/api/types"
)

func UserIsValid(u *types.RegisterUser) bool {
	if len(u.Name) > 0 &&
		len(u.Email) > 0 &&
		len(u.Password) > 0 {
		return true
	}
	return false
}
