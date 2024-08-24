package config

import (
	"log"
)

type ConfigValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Environment string

const (
	Development Environment = "development"
	Staging     Environment = "staging"
	Production  Environment = "production"
)

type RedisConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     uint16 `mapstructure:"port" validate:"required,port_range"`
	Password string `mapstructure:"password" validate:"required"`
}

type Conf struct {
	Environment         Environment `mapstructure:"environment" validate:"required,environment"`
	WebServerPort       uint16      `mapstructure:"webServerPort" validate:"required,port_range"`
	GrpcServerPort      uint16      `mapstructure:"grpcServerPort" validate:"required,port_range"`
	DatabaseURL         string      `mapstructure:"databaseURL" validate:"required,url"`
	TempFileStoragePath string      `mapstructure:"tempFileStoragePath" validate:"required"`
	Redis               RedisConfig `mapstructure:"redis" validate:"required"`
}

func NewConfig() (*Conf, error) {
	cfg, err := loadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	return cfg, nil
}
