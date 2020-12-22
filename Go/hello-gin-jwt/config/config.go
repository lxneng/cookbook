package config

import (
	"sync"

	"github.com/koding/multiconfig"
)

type Config struct {
	ServerPort int    `default:"8100"`
	Dsn        string `default:"hello-gin-jwt.db"`

	SecretKey       string `default:"SecretKeySecretKeySecretKeySecretKey"`
	Issuer          string `default:"AuthService"`
	ExpirationHours int64  `default:"24"`
}

var (
	once sync.Once
	c    *Config
)

func GetConfig() *Config {
	once.Do(func() {
		c = new(Config)
		m := multiconfig.New()
		m.MustLoad(c)
	})
	return c
}
