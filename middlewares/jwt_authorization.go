package middlewares

import (
	"fmt"
	"net/http"

	"Github.com/Yobubble/go-clean-boilerplate/pkg/utils/commons"
	mainConstants "Github.com/Yobubble/go-clean-boilerplate/pkg/utils/constants"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
)

func JwtAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerScehma = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(bearerScehma):]
		claims, err := jwt.VerifyToken(tokenString)
		if err != nil {
			commons.Response(c, nil, mainConstants.VerifyTokenError, http.StatusUnauthorized, err.Error())
		}
		fmt.Print(claims)
		c.Next()
	}
}
