package models

import (
	"hello-gin-jwt/db"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	user := User{
		Password: "passw0rd",
	}

	err := user.HashPassword(user.Password)
	assert.NoError(t, err)

	os.Setenv("passwordHash", user.Password)
}

func TestCreateUserRecord(t *testing.T) {
	var userResult User

	db := db.GetDB()
	err := db.AutoMigrate(&User{})
	assert.NoError(t, err)

	user := User{
		Name:     "Test User",
		Email:    "test@email.com",
		Password: os.Getenv("passwordHash"),
	}

	err = user.CreateUserRecord()
	assert.NoError(t, err)

	db.Where("email = ?", user.Email).Find(&userResult)

	db.Unscoped().Delete(&user)

	assert.Equal(t, "Test User", userResult.Name)
	assert.Equal(t, "test@email.com", userResult.Email)

}

func TestCheckPassword(t *testing.T) {
	hash := os.Getenv("passwordHash")

	user := User{
		Password: hash,
	}

	err := user.CheckPassword("passw0rd")
	assert.NoError(t, err)
}
