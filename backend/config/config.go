package config

type Config struct {
	PostgresConfig PostgresConfig
}

func FromEnv() Config {
	return Config{
		PostgresConfig: postgreConfigFromEnv(),
	}
}
