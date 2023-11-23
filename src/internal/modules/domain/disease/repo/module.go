package repo

import "go.uber.org/fx"

var (
	Module     = fx.Provide(NewDiseaseRepo)
	Invokables = fx.Invoke()
)
