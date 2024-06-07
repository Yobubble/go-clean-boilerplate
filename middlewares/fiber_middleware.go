package middlewares

import (
	"os"

	"Github.com/Yobubble/go-clean-boilerplate/pkg/entities"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/utils"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func JwtAuthentication() fiber.Handler {
	godotenv.Load()
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "missing or malformed jwt" {
		return utils.Response(c, nil, "missing or malformed jwt", fiber.StatusBadRequest, err.Error())
	}
	return utils.Response(c, nil, "invalid or expired jwt", fiber.StatusUnauthorized, err.Error())
}

func JwtPayloadExtraction(c *fiber.Ctx, db *gorm.DB) entities.Payload {
	var payload entities.Payload
	token := c.Locals("user").(*jwt.Token)
	payload.Id = uint(token.Claims.(jwt.MapClaims)["id"].(float64))
	payload.Username = token.Claims.(jwt.MapClaims)["username"].(string)
	return payload
}
