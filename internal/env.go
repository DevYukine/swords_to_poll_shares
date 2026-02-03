package app

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Environment string

const (
	EnvironmentDevelopment Environment = "development"
	EnvironmentStaging     Environment = "staging"
	EnvironmentProduction  Environment = "production"
)

type LogLevel string

const (
	LogLevelDebug LogLevel = "DEBUG"
	LogLevelInfo  LogLevel = "INFO"
	LogLevelWarn  LogLevel = "WARN"
	LogLevelError LogLevel = "ERROR"
)

// Config is the application configuration loaded from environment variables.
type Config struct {
	LogLevel        LogLevel    `env:"LOG_LEVEL" envDefault:"info"`
	DiscordBotToken string      `env:"DISCORD_BOT_TOKEN,required"`
	Env             Environment `env:"ENV" envDefault:"development"`
}

// ProvideConfig is a Fx provider that ensures configuration is initialized and returns a pointer.
func ProvideConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		// If .env file is not found, we can still proceed with environment variables
	}

	config := Config{}

	err = env.Parse(&config)
	if err != nil {
		panic("Error parsing environment variables " + err.Error())
	}
	return &config
}
