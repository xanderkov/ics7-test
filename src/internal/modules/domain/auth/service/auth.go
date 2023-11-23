package service

import (
	"context"
	"hospital/internal/modules/config"
	"hospital/internal/modules/domain/auth/dto"
	doctor_dto "hospital/internal/modules/domain/doctor/dto"
)

//go:generate mockgen -destination mock_test.go -package service . IDoctorRepo

type IDoctorRepo interface {
	// GetByTokenId(ctx context.Context, tokenId string) (*doctor_dto.Doctor, error)
	Create(ctx context.Context, dtm *doctor_dto.CreateDoctor) (*doctor_dto.Doctor, error)
}

type AuthService struct {
	repo    IDoctorRepo
	tokenId string
}

func NewAuthService(repo IDoctorRepo, config config.Config) *AuthService {
	return &AuthService{
		repo:    repo,
		tokenId: config.Secret,
	}
}

func (r *AuthService) SignUp(ctx context.Context, newDoctor *dto.NewDoctor) (*doctor_dto.Doctor, error) {

	createDoctor := &doctor_dto.CreateDoctor{
		Surname:    newDoctor.Surname,
		TokenId:    newDoctor.TokenId,
		Speciality: newDoctor.Speciality,
		Role:       newDoctor.Role,
	}

	// Создаем пользователя
	createdDoctor, err := r.repo.Create(ctx, createDoctor)
	if err != nil {
		return nil, err
	}

	return createdDoctor, nil
}
