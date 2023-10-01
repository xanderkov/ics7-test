package logger

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() (*zap.Logger, zap.AtomicLevel, error) {
	config := zap.NewDevelopmentConfig() // Стандартное local time
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()

	loggerLevel := zap.NewAtomicLevelAt(zap.InfoLevel)

	return logger, loggerLevel, err
}

func InvokeLogger(logger *zap.Logger, lifecycle fx.Lifecycle) {
	lifecycle.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return logger.Sync()
		},
	})
}
