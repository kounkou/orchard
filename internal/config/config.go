package config

import (
	"os"
)

type Config struct {
	DatabaseName string
}

func LoadConfig() *Config {
	return &Config{
		DatabaseName: os.Getenv("DATABASE_NAME"),
	}
}
