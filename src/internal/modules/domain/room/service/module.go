package service

import "go.uber.org/fx"

var (
	Module     = fx.Provide(NewRoomService)
	Invokables = fx.Invoke()
)
