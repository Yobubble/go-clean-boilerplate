package middlewares

import (
	"fmt"
	"net/http"

	"Github.com/Yobubble/go-clean-boilerplate/pkg/utils/commons"
	"Github.com/Yobubble/go-clean-boilerplate/pkg/utils/jwt"
	"github.com/gin-gonic/gin"
)

func JwtAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		claims, err := jwt.VerifyToken(tokenString)
		if err != nil {
			commons.Response(c, nil, commons.VerifyTokenError, http.StatusUnauthorized, err.Error())
		}
		fmt.Print(claims)
		c.Next()
	}
}
