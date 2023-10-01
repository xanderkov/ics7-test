package auth

import (
	"go.uber.org/fx"
	"hospital/internal/modules/domain/auth/service"
	"hospital/internal/modules/domain/doctor/repo"
)

var (
	Module = fx.Options(
		service.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.DoctorRepo) *repo.DoctorRepo { return r },
				fx.As(new(service.IDoctorRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
	)
)
