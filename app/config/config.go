package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/vrischmann/envconfig"
)

type Config struct {
	BindAddress string `envconfig:"BIND_ADDRESS"`
	Port        string `envconfig:"PORT"`
	Debug       bool   `envconfig:"DEBUG"`
}

func NewConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")

	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load() // The Original .env

	config := &Config{
		BindAddress: "0.0.0.0",
		Port:        "8000",
		Debug:       false,
	}

	if err := envconfig.Init(config); err != nil {
		return nil, err
	}

	return config, nil
}
