package types

import (
	"context"
)

type User struct {
	UUID     string   `json:"uuid" form:"-"`
	Username string   `json:"username" form:"username"`
	Groups   []string `json:"groups"`
	Password string   `json:"password" form:"password"`
	IsAdmin  bool     `json:"isAdmin" form:"isAdmin"`
}

type UserCreateOptions struct {
	Username string   `json:"username"`
	Groups   []string `json:"groups"`
	Password string   `json:"is_admin"`
	IsAdmin  bool     `json:"is_admin"`

	// Context can be set with a timeout or can be used to cancel a request.
	Context context.Context `json:"-"`
}
