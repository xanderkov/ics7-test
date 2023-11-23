package domain

import (
	"go.uber.org/fx"
	"hospital/internal/modules/domain/auth"
	"hospital/internal/modules/domain/disease"
	"hospital/internal/modules/domain/doctor"
	"hospital/internal/modules/domain/patient"
	"hospital/internal/modules/domain/room"
)

var (
	Module = fx.Options(

		doctor.Module,
		patient.Module,
		room.Module,
		auth.Module,
		disease.Module,
	)
	Invokables = fx.Options(

		doctor.Invokables,
		patient.Invokables,
		room.Invokables,
		auth.Invokables,
		disease.Invokables,
	)
)
