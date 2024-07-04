package repositories

import "Github.com/Yobubble/go-clean-boilerplate/pkg/features/auth/models"

type AuthRepository interface {
	MatchEmail(email string) (string, string, error) // return user, password
	CreateUser(newUser models.SignUpBodyModel) error
}
