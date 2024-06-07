package usecases

import "Github.com/Yobubble/go-clean-boilerplate/pkg/auth/repositories"

type authUseCase struct {
	repo repositories.AuthRepository
}

// ValidateUser implements AuthUseCase.
func (a *authUseCase) ValidateUser() (string, error) {
	// TODO : call FindEmail
	// TODO : call FindPassword
	panic("unimplemented")
}

func NewAuthUseCase(repo repositories.AuthRepository) AuthUseCase {
	return &authUseCase{
		repo: repo,
	}
}
