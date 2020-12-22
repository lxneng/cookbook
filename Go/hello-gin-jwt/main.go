package main

import (
	"fmt"
	"hello-gin-jwt/config"
	"hello-gin-jwt/db"
	"hello-gin-jwt/handlers"
	"hello-gin-jwt/models"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		v1.POST("/signup", handlers.Signup)
		v1.POST("/signin", handlers.Signin)

		v1.GET("/profile", handlers.Profile)
	}

	return r
}

func dbMigrate() {
	db := db.GetDB()
	db.AutoMigrate(models.GetModels()...)
}

func main() {
	c := config.GetConfig()
	dbMigrate() // run migrate
	r := setupRouter()
	r.Run(fmt.Sprintf(":%v", c.ServerPort))
}
