package middlewares

import (
	"hello-gin-jwt/config"
	"hello-gin-jwt/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authz() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(403, "No Authorization header provided")
			c.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			c.JSON(400, "Incorrect Format of Authorization Token")
			c.Abort()
			return
		}
		cfg := config.GetConfig()
		jwtWrapper := jwt.JwtWrapper{
			SecretKey: cfg.SecretKey,
			Issuer:    cfg.Issuer,
		}

		email, err := jwtWrapper.ParseToken(clientToken)
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}

		c.Set("email", email)

		c.Next()

	}
}
