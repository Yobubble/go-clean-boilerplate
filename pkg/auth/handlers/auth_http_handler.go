package handlers

import (
	"fmt"

	"Github.com/Yobubble/go-clean-boilerplate/pkg/auth/usecases"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type authHttpHandler struct {
	usecase usecases.AuthUseCase
}

// SignIn implements Handler.
func (a *authHttpHandler) SignIn(c *fiber.Ctx) error {
	fmt.Println(c.Path(), "SignIn")
	// parse body
	// call ValidateUser from usecase layer
	return utils.Response(c, nil, "", fiber.StatusOK, "")
}

func NewAuthHttpHandler(usecase usecases.AuthUseCase) AuthHandler {
	return &authHttpHandler{
		usecase: usecase,
	}
}
