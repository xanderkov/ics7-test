package service

import (
	"context"
	"hospital/internal/modules/domain/doctor/dto"
)

//go:generate mockgen -destination mock_test.go -package service . IDoctorRepo

type IDoctorRepo interface {
	GetById(ctx context.Context, id int) (*dto.Doctor, error)
	List(ctx context.Context) (dto.Doctors, error)
	Create(ctx context.Context, dtm *dto.CreateDoctor) (*dto.Doctor, error)
	Update(ctx context.Context, id int, dtm *dto.UpdateDoctor) (*dto.Doctor, error)
	Delete(ctx context.Context, id int) error
	GetByTokenId(ctx context.Context, token string) (*dto.Doctor, error)
}

type DoctorService struct {
	repo IDoctorRepo
}

func NewDoctorService(repo IDoctorRepo) *DoctorService {
	return &DoctorService{
		repo: repo,
	}
}

func (r *DoctorService) GetById(ctx context.Context, id int) (*dto.Doctor, error) {
	return r.repo.GetById(ctx, id)
}

func (r *DoctorService) List(ctx context.Context) (dto.Doctors, error) {
	return r.repo.List(ctx)
}

func (r *DoctorService) Create(ctx context.Context, dtm *dto.CreateDoctor) (*dto.Doctor, error) {
	return r.repo.Create(ctx, dtm)
}

func (r *DoctorService) Update(ctx context.Context, id int, dtm *dto.UpdateDoctor) (*dto.Doctor, error) {
	return r.repo.Update(ctx, id, dtm)
}

func (r *DoctorService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}

func (r *DoctorService) GetByTokenId(ctx context.Context, token string) (*dto.Doctor, error) {
	return r.repo.GetByTokenId(ctx, token)
}
