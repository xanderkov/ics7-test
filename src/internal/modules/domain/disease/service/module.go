package service

import (
	"go.uber.org/fx"
)

var (
	Module     = fx.Provide(NewDiseaseService)
	Invokables = fx.Invoke()
)
