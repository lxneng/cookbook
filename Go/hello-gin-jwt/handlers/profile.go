package handlers

import (
	"hello-gin-jwt/config"
	"hello-gin-jwt/db"
	"hello-gin-jwt/jwt"
	"hello-gin-jwt/models"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Profile returns user data
func Profile(c *gin.Context) {
	var user models.User

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
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error signing token",
		})
		c.Abort()
		return
	}

	db := db.GetDB()
	result := db.Where("email = ?", email).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"msg": "user not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": "could not get user profile",
		})
		c.Abort()
		return
	}

	user.Password = ""

	c.JSON(200, user)

	return
}
