package e2e

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"hospital/internal/modules/db"
	"hospital/internal/modules/db/ent"
	auth_dto "hospital/internal/modules/domain/auth/dto"
	telegram "hospital/internal/modules/view/telegram"
	"hospital/internal/modules/view/telegram/controllers"
	"testing"
)

func telegramServiceTest(t *testing.T, client *ent.Client, controller *controllers.Controller) {
	err := db.TruncateAll(client)
	if err != nil {
		return
	}
	user := telegram.UsersMessage{
		ChatId:       1,
		UserMessages: []string{"Kovel", "Психотерапевт", "Глав врач"},
	}

	newUser := &auth_dto.NewDoctor{
		TokenId:    "1",
		Surname:    "Kovel",
		Speciality: "Психотерапевт",
		Role:       "Глав врач",
	}

	reply := telegram.EndSingUp(&user, user.ChatId, controller)
	assert.Equal(t, reply, "Зарегистрирован")

	reply = telegram.GetInfoAboutDoctor(user.ChatId, controller)
	excepted := fmt.Sprintf("Фамилия: %s \nСпециальность: %s \nРоль: %s \n",
		newUser.Surname, newUser.Speciality, newUser.Role)

	assert.Equal(t, reply, excepted)

}
