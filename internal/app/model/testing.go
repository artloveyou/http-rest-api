package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email: "testing@user.io",
		Password: "password",
	}
}
