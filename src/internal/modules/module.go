package modules

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"hospital/internal/modules/app"
	"hospital/internal/modules/config"
	"hospital/internal/modules/db"
	"hospital/internal/modules/domain"
	"hospital/internal/modules/logger"
	"hospital/internal/modules/view/telegram"
)

var (
	AppModule = fx.Options(
		app.Module,
		logger.Module,
		config.Module,
		db.Module,

		domain.Module,
		telegram.Module,

		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)

	AppInvokables = fx.Options(
		app.Invokables,
		logger.Invokables,
		config.Invokables,
		db.Invokables,

		domain.Invokables,
		telegram.Invokables,
	)
)
