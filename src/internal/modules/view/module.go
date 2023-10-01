package telegram

import (
	"go.uber.org/fx"
	"hospital/internal/modules/view/telegram"
)

var (
	Module     = fx.Provide(telegram.Module)
	Invokables = fx.Invoke(telegram.Invokables)
)
