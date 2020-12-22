package handlers

import (
	"hello-gin-jwt/config"
	"hello-gin-jwt/db"
	"hello-gin-jwt/jwt"
	"hello-gin-jwt/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SigninPayload login body
type SigninPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SigninResponse token response
type SigninResponse struct {
	Token string `json:"token"`
}

// Signin logs users in
func Signin(c *gin.Context) {
	var payload SigninPayload
	var user models.User

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	db := db.GetDB()
	result := db.Where("email = ?", payload.Email).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	cfg := config.GetConfig()

	jwtWrapper := jwt.JwtWrapper{
		SecretKey:       cfg.SecretKey,
		Issuer:          cfg.Issuer,
		ExpirationHours: cfg.ExpirationHours,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error signing token",
		})
		c.Abort()
		return
	}

	tokenResponse := SigninResponse{
		Token: signedToken,
	}

	c.JSON(200, tokenResponse)

	return
}
