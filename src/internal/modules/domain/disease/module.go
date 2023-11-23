package disease

import (
	"go.uber.org/fx"
	"hospital/internal/modules/domain/disease/repo"
	"hospital/internal/modules/domain/disease/service"
)

var (
	Module = fx.Options(
		service.Module,
		repo.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.DiseaseRepo) *repo.DiseaseRepo { return r },
				fx.As(new(service.IDiseaseRepo)),
			),
		),
	)

	Invokables = fx.Options(
		service.Invokables,
		repo.Invokables,
	)
)
