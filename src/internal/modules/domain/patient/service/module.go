package service

import "go.uber.org/fx"

var (
	Module     = fx.Provide(NewPatientService)
	Invokables = fx.Invoke()
)
