package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DBUser   string `envconfig:"DB_USER"`
	DBPass   string `envconfig:"DB_PASS"`
	DBHost   string `envconfig:"DB_HOST"`
	DBPort   string `envconfig:"DB_PORT"`
	DBName   string `envconfig:"DB_NAME"`
	GrpcPort string `envconfig:"GRPC_PORT" default:"50051"`
}

func New() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}
