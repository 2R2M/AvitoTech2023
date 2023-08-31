package config

import (
	"avitoTech/internal/infrastructure/server/config"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB     DB               `envconfig:"db" validate:"required"`
	Server config.SrvConfig `envconfig:"server" validate:"required"`
}

type DB struct {
	Host         string `envconfig:"host" validate:"required"`
	User         string `envconfig:"user" validate:"required"`
	Password     string `envconfig:"password" validate:"required"`
	Port         int    `envconfig:"port" validate:"required"`
	Name         string `envconfig:"name" validate:"required"`
	MaxOpenConns int    `envconfig:"max_open_conns" validate:"required"`
	MaxIdleConns int    `envconfig:"max_idle_conns" validate:"required"`
}

const prefix = "common_service"

func LoadConfig() (*Config, error) {
	cnfg := &Config{}

	if err := envconfig.Process(prefix, cnfg); err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return cnfg, nil
}
