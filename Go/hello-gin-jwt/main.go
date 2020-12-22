package main

import (
	"fmt"
	"hello-gin-jwt/config"
	"hello-gin-jwt/db"
	"hello-gin-jwt/handlers"
	"hello-gin-jwt/middlewares"
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
		public := api.Group("/v1")
		{
			public.POST("/signup", handlers.Signup)
			public.POST("/signin", handlers.Signin)
		}

		protected := api.Group("/v1").Use(middlewares.Authz())
		{
			protected.GET("/profile", handlers.Profile)
		}
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
