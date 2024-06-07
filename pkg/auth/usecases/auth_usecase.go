package usecases

type AuthUseCase interface {
	ValidateUser() (string, error) // <- return token
}
