package migrations

import (
	"Github.com/Yobubble/go-clean-boilerplate/pkg/entities"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/utils/data"
	"gorm.io/gorm"
)

type postgresqlMigration struct {
	db *gorm.DB
}

// MockDataMigrate implements Migration.
func (p *postgresqlMigration) MockDataMigrate() error {
	result := p.db.Create(data.Users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// TableMigrate implements Migration.
func (p *postgresqlMigration) TableMigrate() error {
	err := p.db.AutoMigrate(entities.User{})
	if err != nil {
		return err
	}
	return nil
}

func NewPostgresqlMigration(db *gorm.DB) Migration {
	return &postgresqlMigration{
		db: db,
	}
}
