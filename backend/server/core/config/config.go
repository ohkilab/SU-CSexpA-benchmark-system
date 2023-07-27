package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DBUser              string `envconfig:"MYSQL_USER" required:"true"`
	DBPass              string `envconfig:"MYSQL_PASSWORD" required:"true"`
	DBHost              string `envconfig:"BACKEND_MYSQL_HOST" required:"true"`
	DBPort              string `envconfig:"BACKEND_MYSQL_PORT" required:"true"`
	DBName              string `envconfig:"MYSQL_DATABASE" required:"true"`
	BenchmarkHost       string `envconfig:"BACKEND_BENCHMARK_HOST" required:"true"`
	BenchmarkPort       string `envconfig:"BENCHMARK_GRPC_PORT" required:"true"`
	GrpcPort            string `envconfig:"BACKEND_GRPC_PORT" default:"50051"`
	JwtSecret           string `envconfig:"BACKEND_JWT_SECRET" required:"true"`
	INIT_ADMIN_NAME     string `envconfig:"BACKEND_INIT_ADMIN_USER_NAME" required:"true"`
	INIT_ADMIN_PASSWORD string `envconfig:"BACKEND_INIT_ADMIN_PASSWORD" required:"true"`
	StoragePath         string `envconfig:"BACKEND_STORAGE_PATH" required:"true"`
}

func New() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}
