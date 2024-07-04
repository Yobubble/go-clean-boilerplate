package data

import "Github.com/Yobubble/go-clean-boilerplate/pkg/entities"

var (
	Users = []entities.User{
		{
			Username: "Cyan",
			Email:    "cyan@gmail.com",
			Password: "$2a$10$CcQYc3lTkAz9zgL7.A1lU.HISfIjsSoyyou5GZQ90eAoq.Qi5JfFW", // cyanpassword
		},
		{
			Username: "Blue",
			Email:    "blue@gmail.com",
			Password: "$2a$10$pLvr5EkdfGkeh55gjXC7cuPuKZWs70jncxvHxYmTU0sFAzr4rz6xe", // bluepassword
		},
	}
)
