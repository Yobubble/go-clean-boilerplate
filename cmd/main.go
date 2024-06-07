package main

import (
	"Github.com/Yobubble/go-clean-boilerplate/configs"
	"Github.com/Yobubble/go-clean-boilerplate/databases"
	"Github.com/Yobubble/go-clean-boilerplate/servers"
)

func main() {
	cfg := configs.LoadConfig()
	db := databases.NewPostgresqlDatabase(cfg)
	servers.NewFiberServer(db.GetDB())
}
