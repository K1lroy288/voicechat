package config

import (
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	ServerHost string
	ServerPort string
	HttpTls    string
}

var (
	instance *Config
	once     sync.Once
)

func loadConfig() *Config {
	once.Do(func() {
		instance = &Config{
			ServerHost: os.Getenv("SERVER_HOST"),
			ServerPort: os.Getenv("SERVER_PORT"),
			HttpTls:    os.Getenv("HTTP_TLS"),
		}
	})

	return instance
}

func GetConfig() *Config {
	if instance == nil {
		return loadConfig()
	}

	return instance
}
