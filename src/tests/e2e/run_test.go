package e2e

import (
	"context"
	"go.uber.org/fx"
	"hospital/internal/modules/db/ent"
	auth_serv "hospital/internal/modules/domain/auth/service"
	doctor_servis "hospital/internal/modules/domain/doctor/service"
	patient_servis "hospital/internal/modules/domain/patient/service"
	room_servis "hospital/internal/modules/domain/room/service"
	"hospital/internal/modules/view/telegram/controllers"
	"testing"
)

func TestServices(t *testing.T) {
	fx.New(
		testModule,
		testInvokables,

		fx.Supply(t),
		fx.Invoke(execTests),
	).Run()

}

func execTests(
	t *testing.T,
	doctorService *doctor_servis.DoctorService,
	authService *auth_serv.AuthService,
	patientService *patient_servis.PatientService,
	roomService *room_servis.RoomService,

	client *ent.Client,
	lifecycle fx.Lifecycle,
	shutdowner fx.Shutdowner,
	controller *controllers.Controller,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				telegramServiceTest(t, client, controller)

				_ = shutdowner.Shutdown()
			}()

			return nil
		},
	})
}
