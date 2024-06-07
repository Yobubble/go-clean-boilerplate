package migrations

import "gorm.io/gorm"

type postgresqlMigration struct {
	db *gorm.DB
}

// MockDataMigrate implements Migration.
func (p *postgresqlMigration) MockDataMigrate() error {
	panic("unimplemented")
}

// TableMigrate implements Migration.
func (p *postgresqlMigration) TableMigrate() error {
	panic("unimplemented")
}

func NewPostgresqlMigration(db *gorm.DB) Migration {
	return &postgresqlMigration{
		db: db,
	}
}
