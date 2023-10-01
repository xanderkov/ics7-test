package repo

import "go.uber.org/fx"

var (
	Module     = fx.Provide(NewDoctorRepo)
	Invokables = fx.Invoke()
)
