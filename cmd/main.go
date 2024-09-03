package main

import (
	"Github.com/Yobubble/go-clean-boilerplate/configs"
	"Github.com/Yobubble/go-clean-boilerplate/databases"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/migrations"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/utils/commons"
	"Github.com/Yobubble/go-clean-boilerplate/servers"
)

func main() {
	commons.NewLogger()
	cfg := configs.LoadConfig()
	db := databases.NewPostgresqlDatabase(cfg).GetDB()

	migrations.NewPostgresqlMigration(db).TableMigrate()
	migrations.NewPostgresqlMigration(db).MockDataMigrate()

	servers.NewGinServer(cfg, db).Start()
}
