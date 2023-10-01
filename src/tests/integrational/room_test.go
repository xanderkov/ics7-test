package integrational

import (
	"context"
	"github.com/stretchr/testify/assert"
	"hospital/internal/modules/db/ent"
	auth_dto "hospital/internal/modules/domain/auth/dto"
	auth_serv "hospital/internal/modules/domain/auth/service"
	"hospital/internal/modules/domain/room/dto"
	"hospital/internal/modules/domain/room/service"
	"testing"
)

func roomServiceTest(t *testing.T, service *service.RoomService, authService *auth_serv.AuthService, client *ent.Client) {
	err := truncateAll(client)
	// Регистрируем пользователя
	newUser := &auth_dto.NewDoctor{
		TokenId:    "1",
		Surname:    "Kovel",
		Speciality: "Психотерапевт",
		Role:       "Глав врач",
	}
	currentUser, err := authService.SignUp(context.Background(), newUser)
	assert.NoError(t, err)
	// Возврат, т.к. не получится создать сессию с пользователем
	if err != nil {
		return
	}

	// Создаем контекст с сессией
	ctx := makeCtxByUser(currentUser)

	newroom := &dto.CreateRoom{
		Num:            1,
		NumberPatients: 2,
		NumberBeds:     2,
		Floor:          2,
		TypeRoom:       "5",
	}
	room, err := service.Create(ctx, newroom)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	t1, err := service.GetByNum(ctx, room.Id)
	assert.NoError(t, err)
	assert.Equal(t, room, t1)

	ts1, err := service.List(ctx)
	assert.NoError(t, err)
	assert.Equal(t, room, ts1[0])

	room.Floor += 1
	updateUser := &dto.UpdateRoom{
		Num:            room.Num,
		NumberPatients: room.NumberPatients,
		NumberBeds:     room.NumberBeds,
		Floor:          room.Floor,
		TypeRoom:       room.TypeRoom,
	}
	t2, err := service.Update(ctx, room.Id, updateUser)
	assert.NoError(t, err)
	assert.Equal(t, room, t2)

	err = service.Delete(ctx, room.Id)
	assert.NoError(t, err)

	_, err = service.GetByNum(ctx, room.Id)
	assert.Error(t, err)

	ts2, err := service.List(ctx)
	assert.NoError(t, err)
	assert.Empty(t, ts2)

}
