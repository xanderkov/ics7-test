package service

import (
	"context"
	"hospital/internal/modules/domain/room/dto"
)

//go:generate mockgen -destination mock_test.go -package service . IRoomRepo

type IRoomRepo interface {
	GetByNum(ctx context.Context, num int) (*dto.Room, error)
	List(ctx context.Context) (dto.Rooms, error)
	Create(ctx context.Context, dtm *dto.CreateRoom) (*dto.Room, error)
	Update(ctx context.Context, num int, dtm *dto.UpdateRoom) (*dto.Room, error)
	Delete(ctx context.Context, num int) error
}

type RoomService struct {
	repo IRoomRepo
}

func NewRoomService(repo IRoomRepo) *RoomService {
	return &RoomService{
		repo: repo,
	}
}

func (r *RoomService) GetByNum(ctx context.Context, num int) (*dto.Room, error) {
	return r.repo.GetByNum(ctx, num)
}

func (r *RoomService) List(ctx context.Context) (dto.Rooms, error) {
	return r.repo.List(ctx)
}

func (r *RoomService) Create(ctx context.Context, dtm *dto.CreateRoom) (*dto.Room, error) {
	return r.repo.Create(ctx, dtm)
}

func (r *RoomService) Update(ctx context.Context, num int, dtm *dto.UpdateRoom) (*dto.Room, error) {
	return r.repo.Update(ctx, num, dtm)
}

func (r *RoomService) Delete(ctx context.Context, num int) error {
	return r.repo.Delete(ctx, num)
}
