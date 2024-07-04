package usecases

import "Github.com/Yobubble/go-clean-boilerplate/pkg/features/auth/models"

type AuthUseCase interface {
	ValidateUser(user models.SignInBodyModel) (string, error) // return token
	InsertNewUser(newUser models.SignUpBodyModel) error
}
