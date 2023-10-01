package patient

import (
	"go.uber.org/fx"
	"hospital/internal/modules/domain/patient/repo"
	"hospital/internal/modules/domain/patient/service"
)

var (
	Module = fx.Options(
		service.Module,
		repo.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.PatientRepo) *repo.PatientRepo { return r },
				fx.As(new(service.IPatientRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
		repo.Invokables,
	)
)
