package service

import (
	"context"
	"hospital/internal/modules/domain/patient/dto"
)

//go:generate mockgen -destination mock_test.go -package service . IPatientRepo

type IPatientRepo interface {
	GetById(ctx context.Context, id int) (*dto.Patient, error)
	List(ctx context.Context) (dto.Patients, error)
	Create(ctx context.Context, dtm *dto.CreatePatient) (*dto.Patient, error)
	Update(ctx context.Context, num int, dtm *dto.UpdatePatient) (*dto.Patient, error)
	Delete(ctx context.Context, num int) error
}

type PatientService struct {
	repo IPatientRepo
}

func NewPatientService(repo IPatientRepo) *PatientService {
	return &PatientService{
		repo: repo,
	}
}

func (r *PatientService) GetById(ctx context.Context, id int) (*dto.Patient, error) {
	return r.repo.GetById(ctx, id)
}

func (r *PatientService) List(ctx context.Context) (dto.Patients, error) {
	return r.repo.List(ctx)
}

func (r *PatientService) Create(ctx context.Context, dtm *dto.CreatePatient) (*dto.Patient, error) {
	return r.repo.Create(ctx, dtm)
}

func (r *PatientService) Update(ctx context.Context, id int, dtm *dto.UpdatePatient) (*dto.Patient, error) {
	return r.repo.Update(ctx, id, dtm)
}

func (r *PatientService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
