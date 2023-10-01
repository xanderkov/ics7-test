package repo

import (
	"context"
	"hospital/internal/modules/db"
	"hospital/internal/modules/db/ent"
	"hospital/internal/modules/domain/room/dto"
)

type RoomRepo struct {
	client *ent.Client
}

func NewRoomRepo(client *ent.Client) *RoomRepo {
	return &RoomRepo{
		client: client,
	}
}

func (r *RoomRepo) GetByNum(ctx context.Context, id int) (*dto.Room, error) {
	Room, err := r.client.Room.Get(ctx, id)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToRoomDTO(Room), nil
}

func (r *RoomRepo) List(ctx context.Context) (dto.Rooms, error) {
	Rooms, err := r.client.Room.Query().All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToRoomDTOs(Rooms), nil
}

func (r *RoomRepo) Create(ctx context.Context, dtm *dto.CreateRoom) (*dto.Room, error) {
	Room, err := r.client.Room.Create().
		SetNumber(dtm.Num).
		SetNumberPatients(dtm.NumberPatients).
		SetFloor(dtm.Floor).
		SetNumberBeds(dtm.NumberBeds).
		SetTypeRoom(dtm.TypeRoom).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToRoomDTO(Room), nil
}

func (r *RoomRepo) Update(ctx context.Context, id int, dtm *dto.UpdateRoom) (*dto.Room, error) {
	Room, err := r.client.Room.UpdateOneID(id).
		SetNumberPatients(dtm.NumberPatients).
		SetFloor(dtm.Floor).
		SetNumberBeds(dtm.NumberBeds).
		SetTypeRoom(dtm.TypeRoom).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToRoomDTO(Room), nil
}

func (r *RoomRepo) Delete(ctx context.Context, id int) error {
	err := r.client.Room.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return db.WrapError(err)
	}

	return nil
}

func (r *RoomRepo) Restore(ctx context.Context, id int) (*dto.Room, error) {
	Room, err := r.client.Room.UpdateOneID(id).Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToRoomDTO(Room), nil
}

func ToRoomDTO(model *ent.Room) *dto.Room {
	if model == nil {
		return nil
	}
	return &dto.Room{
		Id:             model.ID,
		TypeRoom:       model.TypeRoom,
		NumberBeds:     model.NumberBeds,
		Floor:          model.Floor,
		NumberPatients: model.NumberPatients,
		Num:            model.Number,
	}
}

func ToRoomDTOs(models ent.Rooms) dto.Rooms {
	if models == nil {
		return nil
	}
	dtms := make(dto.Rooms, len(models))
	for i := range models {
		dtms[i] = ToRoomDTO(models[i])
	}
	return dtms
}
