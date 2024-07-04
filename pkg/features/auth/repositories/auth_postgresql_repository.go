package repositories

import (
	"errors"

	"Github.com/Yobubble/go-clean-boilerplate/pkg/entities"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/features/auth/models"
	"gorm.io/gorm"
)

type authPostgresqlRepository struct {
	db *gorm.DB
}

// CreateUser implements AuthRepository.
func (a *authPostgresqlRepository) CreateUser(newUser models.SignUpBodyModel) error {
	result := a.db.Create(&entities.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// MatchEmail implements AuthRepository.
func (a *authPostgresqlRepository) MatchEmail(email string) (string, string, error) {
	var user entities.User
	result := a.db.Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		if result.Error != nil {
			return "", "", result.Error
		}
		return "", "", errors.New("user not found")
	}
	return user.Username, user.Password, nil
}

func NewAuthPostgresqlRepository(db *gorm.DB) AuthRepository {
	return &authPostgresqlRepository{
		db: db,
	}
}
