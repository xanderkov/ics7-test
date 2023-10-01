package telegram

import (
	"go.uber.org/fx"
	"hospital/internal/modules/view/telegram/controllers"
)

var (
	Module     = fx.Provide(controllers.NewController)
	Invokables = fx.Invoke(startBot)
)
