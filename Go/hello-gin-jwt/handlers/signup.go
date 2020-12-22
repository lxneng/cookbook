package handlers

import (
	"hello-gin-jwt/models"
	"log"

	"github.com/gin-gonic/gin"
)

// Signup creates a user in db
func Signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)

		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()

		return
	}

	err = user.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())

		c.JSON(500, gin.H{
			"msg": "error hashing password",
		})
		c.Abort()

		return
	}

	err = user.CreateUserRecord()
	if err != nil {
		log.Println(err)

		c.JSON(500, gin.H{
			"msg": "error creating user",
		})
		c.Abort()

		return
	}

	c.JSON(200, user)
}
