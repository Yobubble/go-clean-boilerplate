package handlers

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	SignIn(c *fiber.Ctx) error
}
