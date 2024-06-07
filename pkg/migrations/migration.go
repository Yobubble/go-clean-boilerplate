package migrations

type Migration interface {
	MockDataMigrate() error
	TableMigrate() error
}
