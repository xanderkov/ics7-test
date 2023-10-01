package integrational

import (
	"context"
	"github.com/stretchr/testify/assert"
	"hospital/internal/modules/db/ent"
	"hospital/internal/modules/domain/doctor/dto"
	"testing"

	auth_dto "hospital/internal/modules/domain/auth/dto"
	auth_serv "hospital/internal/modules/domain/auth/service"
	doctor_server "hospital/internal/modules/domain/doctor/service"
)

func doctorServiceTest(t *testing.T, DoctorService *doctor_server.DoctorService, authService *auth_serv.AuthService, client *ent.Client) {
	err := truncateAll(client)
	// Регистрируем пользователя
	newDoctor := &auth_dto.NewDoctor{
		TokenId:    "1",
		Surname:    "Kovel",
		Speciality: "Психотерапевт",
		Role:       "Глав врач",
	}
	currentDoctor, err := authService.SignUp(context.Background(), newDoctor)
	assert.NoError(t, err)
	// Возврат, т.к. не получится создать сессию с пользователем
	if err != nil {
		return
	}

	// Создаем контекст с сессией
	ctx := makeCtxByUser(currentDoctor)

	u1, err := DoctorService.GetById(ctx, currentDoctor.Id)
	assert.NoError(t, err)
	assert.Equal(t, currentDoctor, u1)

	Doctors1, err := DoctorService.List(ctx)
	assert.NoError(t, err)
	assert.Equal(t, currentDoctor, Doctors1[0])

	currentDoctor.TokenId += "1"
	updateDoctor := &dto.UpdateDoctor{
		Surname:    currentDoctor.Surname,
		Speciality: currentDoctor.Speciality,
		Role:       currentDoctor.Role,
		TokenId:    currentDoctor.TokenId,
	}
	u2, err := DoctorService.Update(ctx, currentDoctor.Id, updateDoctor)
	assert.NoError(t, err)
	assert.Equal(t, currentDoctor, u2)

	err = DoctorService.Delete(ctx, currentDoctor.Id)
	assert.NoError(t, err)

	_, err = DoctorService.GetById(ctx, currentDoctor.Id)
	assert.Error(t, err)

	Doctors2, err := DoctorService.List(ctx)
	assert.NoError(t, err)
	assert.Empty(t, Doctors2)

}
