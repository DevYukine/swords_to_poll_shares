package app

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ProvideLogger builds a zap.Logger based on Config.
func ProvideLogger(cfg *Config) (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.OutputPaths = []string{"stdout"}

	var level zapcore.Level
	switch cfg.LogLevel {
	case LogLevelDebug:
		level = zap.DebugLevel
	case LogLevelWarn:
		level = zap.WarnLevel
	case LogLevelInfo:
		level = zap.InfoLevel
	case LogLevelError:
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	config.Level = zap.NewAtomicLevelAt(level)
	config.Development = false

	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return logger, nil
}
