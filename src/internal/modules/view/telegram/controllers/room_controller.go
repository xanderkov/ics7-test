package controllers

import (
	"context"
	dto1 "hospital/internal/modules/domain/room/dto"
)

func (r *Controller) Room(ctx context.Context, id int) (*dto1.Room, error) {
	user, err := r.roomService.GetByNum(ctx, id)

	return user, err
}

func (r *Controller) AddRoom(ctx context.Context, room *dto1.CreateRoom) (*dto1.Room, error) {
	user, err := r.roomService.Create(ctx, room)
	return user, err
}

func (r *Controller) GetAllRooms(ctx context.Context) (dto1.Rooms, error) {
	room, err := r.roomService.List(ctx)
	return room, err
}
