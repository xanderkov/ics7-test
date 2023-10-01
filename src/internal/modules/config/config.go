package config

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"hospital/internal/modules/app"
)

type Config struct {
	Secret string `envconfig:"SECRET"`

	DBDriver     string `envconfig:"DB_DRIVER" default:"postgres"`
	DBConnection string `envconfig:"DB_CONN_STRING"`

	DBUser string `envconfig:"DB_USER"`
	DBPass string `envconfig:"DB_PASS"`
	DBHost string `envconfig:"DB_HOST" default:"db"`
	DBPort string `envconfig:"DB_PORT" default:"5432"`
	DBName string `envconfig:"DB_NAME" default:"main"`

	SQLSlowThreshold int    `envconfig:"SQL_SLOW_THRESHOLD" default:"600"`
	AutoMigrate      bool   `envconfig:"AUTO_MIGRATE" default:"true"`
	LogLevel         string `envconfig:"LOG_LEVEL" default:"info" validate:"oneof=debug info warn error dpanic panic fatal"`
	TraceSQLCommands bool

	TelegramToken string `envconfig:"TELEGRAM_APITOKEN"`
}

func NewConfig(app app.App, logger *zap.Logger, logLevel zap.AtomicLevel) (Config, error) {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		return Config{}, err
	}

	logger.Info("Configurated", zap.Any("config", config))

	err = logLevel.UnmarshalText([]byte(config.LogLevel))
	if err != nil {
		return Config{}, err
	}

	return config, err
}
