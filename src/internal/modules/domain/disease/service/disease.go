package service

import (
	"context"
	"hospital/internal/modules/domain/disease/dto"
)

//go:generate mockgen -destination mock_test.go -package service . IDiseaseRepo

type IDiseaseRepo interface {
	GetById(ctx context.Context, id int) (*dto.Disease, error)
	List(ctx context.Context) (dto.Diseases, error)
	Create(ctx context.Context, dtm *dto.CreateDisease) (*dto.Disease, error)
	Update(ctx context.Context, num int, dtm *dto.UpdateDisease) (*dto.Disease, error)
	Delete(ctx context.Context, num int) error
}

type DiseaseService struct {
	repo IDiseaseRepo
}

func NewDiseaseService(repo IDiseaseRepo) *DiseaseService {
	return &DiseaseService{
		repo: repo,
	}
}

func (r *DiseaseService) GetById(ctx context.Context, id int) (*dto.Disease, error) {
	return r.repo.GetById(ctx, id)
}

func (r *DiseaseService) List(ctx context.Context) (dto.Diseases, error) {
	return r.repo.List(ctx)
}

func (r *DiseaseService) Create(ctx context.Context, dtm *dto.CreateDisease) (*dto.Disease, error) {
	return r.repo.Create(ctx, dtm)
}

func (r *DiseaseService) Update(ctx context.Context, id int, dtm *dto.UpdateDisease) (*dto.Disease, error) {
	return r.repo.Update(ctx, id, dtm)
}

func (r *DiseaseService) Delete(ctx context.Context, id int) error {
	return r.repo.Delete(ctx, id)
}
