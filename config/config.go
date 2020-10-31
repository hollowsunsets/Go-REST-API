package config

import "os"

type Config struct {
	PostgresUser string
	PostgresPass string
	PostgresDBName string
	TokenSecret string
}

func New() *Config {
	return &Config{
		PostgresUser: os.Getenv("POSTGRES_USER"),
		PostgresPass: os.Getenv("POSTGRES_PASS"),
		PostgresDBName: os.Getenv("POSTGRES_DB_NAME"),
		TokenSecret: os.Getenv("JWT_TOKEN_SECRET"),
	}
}