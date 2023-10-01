package controllers

import (
	auth_serv "hospital/internal/modules/domain/auth/service"
	disease_servis "hospital/internal/modules/domain/disease/service"
	doctor_server "hospital/internal/modules/domain/doctor/service"
	patient_servis "hospital/internal/modules/domain/patient/service"
	room_servis "hospital/internal/modules/domain/room/service"
)

type Controller struct {
	doctorService  *doctor_server.DoctorService
	authService    *auth_serv.AuthService
	patientService *patient_servis.PatientService
	roomService    *room_servis.RoomService
	diseaseService *disease_servis.DiseaseService
}

func NewController(
	doctorService *doctor_server.DoctorService,
	authService *auth_serv.AuthService,
	patientService *patient_servis.PatientService,
	roomService *room_servis.RoomService,
	diseaseService *disease_servis.DiseaseService,
) *Controller {

	r := &Controller{
		authService:    authService,
		doctorService:  doctorService,
		patientService: patientService,
		roomService:    roomService,
		diseaseService: diseaseService,
	}

	return r
}
