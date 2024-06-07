package repositories

type AuthRepository interface {
	FindEmail() error
	FindPassword() error
}
