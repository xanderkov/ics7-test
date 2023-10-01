package controllers

import (
	"context"
	"hospital/internal/modules/domain/auth/dto"
	dto1 "hospital/internal/modules/domain/doctor/dto"
)

func (r *Controller) SingUp(ctx context.Context, newUser *dto.NewDoctor) (*dto1.Doctor, error) {
	user, err := r.authService.SignUp(ctx, newUser)

	return user, err
}
