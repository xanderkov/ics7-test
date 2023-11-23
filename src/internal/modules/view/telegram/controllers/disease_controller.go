package controllers

import (
	"context"
	dto1 "hospital/internal/modules/domain/disease/dto"
)

func (r *Controller) Disease(ctx context.Context, id int) (*dto1.Disease, error) {
	user, err := r.diseaseService.GetById(ctx, id)

	return user, err
}

func (r *Controller) AddDisease(ctx context.Context, disease *dto1.CreateDisease) (*dto1.Disease, error) {
	user, err := r.diseaseService.Create(ctx, disease)
	return user, err
}

func (r *Controller) GetAllDiseases(ctx context.Context) (dto1.Diseases, error) {
	disease, err := r.diseaseService.List(ctx)
	return disease, err
}
