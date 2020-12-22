package db

import (
	"hello-gin-jwt/config"
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

func GetDB() *gorm.DB {
	once.Do(func() {
		c := config.GetConfig()
		var err error
		db, err = gorm.Open(sqlite.Open(c.Dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect postgres!")
		}
	})
	return db
}
