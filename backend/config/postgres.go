package config

import (
	"fmt"
	"os"
)

// TODO: Add possibility for not all parameters to be present
type PostgresConfig struct {
	Host     string
	Db       string
	User     string
	Password string
	SslMode  string
}

func (p *PostgresConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s", p.Host, p.User, p.Db, p.Password, p.SslMode)
}

func postgreConfigFromEnv() PostgresConfig {
	return PostgresConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Db:       os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		SslMode:  os.Getenv("POSTGRES_SSL"),
	}
}
