package repo

import "go.uber.org/fx"

var (
	Module     = fx.Provide(NewPatientRepo)
	Invokables = fx.Invoke()
)
