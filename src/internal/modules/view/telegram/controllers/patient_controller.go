package controllers

import (
	"context"
	dto1 "hospital/internal/modules/domain/patient/dto"
)

func (r *Controller) Patient(ctx context.Context, id int) (*dto1.Patient, error) {
	user, err := r.patientService.GetById(ctx, id)

	return user, err
}

func (r *Controller) AddPatient(ctx context.Context, patient *dto1.CreatePatient) (*dto1.Patient, error) {
	user, err := r.patientService.Create(ctx, patient)
	return user, err
}

func (r *Controller) GetAllPatients(ctx context.Context) (dto1.Patients, error) {
	user, err := r.patientService.List(ctx)
	return user, err
}
