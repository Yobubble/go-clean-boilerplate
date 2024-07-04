package usecases

import (
	"os"
	"time"

	"Github.com/Yobubble/go-clean-boilerplate/pkg/features/auth/models"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/features/auth/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
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

	token, err := createToken(username)
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

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	godotenv.Load()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
