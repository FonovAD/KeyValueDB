package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port int `env:"KEYVALUE_PORT,required"`
}

func WithDefault(ctx context.Context) *Config {
	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}
	if c.Port == 0 {
		c.Port = 4412
	}
	return &c
}
