package benchamark

import (
	"hospital/internal/modules/db"
	"hospital/internal/modules/db/ent"
	"hospital/internal/modules/view/telegram"
	"hospital/internal/modules/view/telegram/controllers"
	"testing"
)

func AddDoctorBenchmark(t *testing.B, client *ent.Client, controller *controllers.Controller) {
	err := db.TruncateAll(client)
	if err != nil {
		return
	}
	user := telegram.UsersMessage{
		ChatId:       1,
		UserMessages: []string{"Kovel", "Психотерапевт", "Глав врач"},
	}

	for i := 0; i < t.N; i++ {
		_ = telegram.EndSingUp(&user, int64(i), controller)
	}

	for i := 0; i < t.N; i++ {
		_ = telegram.GetInfoAboutDoctor(int64(i), controller)
	}

}
