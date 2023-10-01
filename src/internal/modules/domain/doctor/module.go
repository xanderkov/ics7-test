package doctor

import (
	"go.uber.org/fx"
	"hospital/internal/modules/domain/doctor/repo"
	"hospital/internal/modules/domain/doctor/service"
)

var (
	Module = fx.Options(
		service.Module,
		repo.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.DoctorRepo) *repo.DoctorRepo { return r },
				fx.As(new(service.IDoctorRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
		repo.Invokables,
	)
)
