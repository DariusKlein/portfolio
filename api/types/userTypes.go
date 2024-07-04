package types

import (
	"portfolio/database/ent/user"
)

type Username struct {
	Name string
}

type LoginUser struct {
	Email    string
	Password string
}

type User struct {
	ID       int       `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"-"`
	Role     user.Role `json:"role,omitempty"`
}
