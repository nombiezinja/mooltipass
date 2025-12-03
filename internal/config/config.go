package config

import (
	"log"
	"log/slog"
	"os"
)

type Config struct {
	ConfigVars *ConfigVars
	// Database                     database.DbInterface
	Logger *slog.Logger
}

// ConfigInterface is asked for everywhere to avoid tight coupling to the concrete Config struct.
type ConfigInterface interface {
	GetConfigVars() *ConfigVars
	GetLogger() *slog.Logger
}

func (c *Config) GetConfigVars() *ConfigVars {
	return c.ConfigVars
}
func (c *Config) GetLogger() *slog.Logger {
	return c.Logger
}

func New() (ConfigInterface, error) {

	config := &Config{
		ConfigVars: MustLoadConfigVars(),
	}

	// Initialize logger.
	config.Logger = slog.Default()

	// Connect to db

	return config, nil
}

func MustLoadConfigVars() *ConfigVars {
	return &ConfigVars{
		Environment:      mustGetEnv("ENVIRONMENT"),
		AppName:          mustGetEnv("APP_NAME"),
		JwtSigningSecret: mustGetEnv("JWT_SIGNING_SECRET"),
		LogLevel:         mustGetEnv("LOG_LEVEL"),
		Port:             mustGetEnv("PORT"),
		DatabaseURL:      mustGetEnv("DATABASE_URL"),
	}
}

func mustGetEnv(key string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		log.Panicf("Unable to get env var %s", key)
	}

	return value
}
