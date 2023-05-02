package main

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DBUser string `envconfig:"DB_USER" required:"true"`
	DBPass string `envconfig:"DB_PASS" required:"true"`
	DBHost string `envconfig:"DB_HOST" required:"true"`
	DBPort string `envconfig:"DB_PORT" required:"true"`
	DBName string `envconfig:"DB_NAME" required:"true"`
}

func newConfig() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}
