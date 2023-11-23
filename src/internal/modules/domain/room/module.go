package room

import (
	"go.uber.org/fx"
	"hospital/internal/modules/domain/room/repo"
	"hospital/internal/modules/domain/room/service"
)

var (
	Module = fx.Options(
		service.Module,
		repo.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.RoomRepo) *repo.RoomRepo { return r },
				fx.As(new(service.IRoomRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
		repo.Invokables,
	)
)
