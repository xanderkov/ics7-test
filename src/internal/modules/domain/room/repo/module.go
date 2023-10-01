package repo

import "go.uber.org/fx"

var (
	Module     = fx.Provide(NewRoomRepo)
	Invokables = fx.Invoke()
)
