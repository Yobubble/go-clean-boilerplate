package repositories

import "gorm.io/gorm"

type authPostgresqlRepository struct {
	db *gorm.DB
}

// FindEmail implements AuthRepository.
func (a *authPostgresqlRepository) FindEmail() error {
	panic("unimplemented")
}

// FindPassword implements AuthRepository.
func (a *authPostgresqlRepository) FindPassword() error {
	panic("unimplemented")
}

func NewAuthPostgresqlRepository(db *gorm.DB) AuthRepository {
	return &authPostgresqlRepository{
		db: db,
	}
}
