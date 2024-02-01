package config

import (
	"context"
	"fmt"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type Settings struct {
	DatabaseName     string `env:"MONGODB_NAME"`
	DatabaseHost     string `env:"MONGODB_HOST"`
	DatabasePort     string `env:"MONGODB_PORT"`
	DatabaseUser     string `env:"MONGODB_USER"`
	DatabasePassword string `env:"MONGODB_PASSWORD"`
}

var Config = Settings{}

func LoadConfig() error {
	ctx := context.Background()

	if err := envconfig.Process(ctx, &Config); err != nil {
		log.Fatal(fmt.Sprintf("Error loading environment variables from .env file: %s", err))
		return err
	}

	return nil
}
