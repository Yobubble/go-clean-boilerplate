package usecases

import (
	"Github.com/Yobubble/go-clean-boilerplate/pkg/features/auth/models"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/features/auth/repositories"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/utils/jwt"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	repo repositories.AuthRepository
}

// InsertNewUser implements AuthUseCase.
func (a *authUseCase) InsertNewUser(newUser models.SignUpBodyModel) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		return err
	}

	newUser.Password = string(hashedPassword)
	if err := a.repo.CreateUser(newUser); err != nil {
		return err
	}
	return nil
}

// ValidateUser implements AuthUseCase.
func (a *authUseCase) ValidateUser(user models.SignInBodyModel) (string, error) {
	username, password, err := a.repo.MatchEmail(user.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password)); err != nil {
		return "", err
	}

	token, err := jwt.CreateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func NewAuthUseCase(repo repositories.AuthRepository) AuthUseCase {
	return &authUseCase{
		repo: repo,
	}
}
