package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pah-dev/quik-wallet/utils"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := http.StatusOK
		msg :=""
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			code = http.StatusBadRequest
			msg = "Request parameter error"
		} else {
			tokenString := authHeader[len(BEARER_SCHEMA):]
			_, err := utils.ParseToken(tokenString)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = 20002
					msg = "Token authentication failed"
				default:
					code = 20001
					msg = "Token has timed out"
				}
			}
		}
		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  msg,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
