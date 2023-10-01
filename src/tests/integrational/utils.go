package integrational

import (
	"context"
	"github.com/google/uuid"
	"hospital/internal/models/session"
	"hospital/internal/modules/db/ent"
	"hospital/internal/modules/domain/doctor/dto"
)

func makeCtxByUser(doctor *dto.Doctor) context.Context {
	ss := session.Session{
		SessionID: uuid.NewString(),
		UserId:    doctor.Id,
	}

	ctx := context.Background()
	return session.SetSessionToCtx(ctx, ss)
}

func truncateAll(client *ent.Client) error {
	_, err := client.Patient.Delete().Exec(context.Background())
	if err != nil {
		return err
	}

	_, err = client.Doctor.Delete().Exec(context.Background())
	if err != nil {
		return err
	}

	_, err = client.Room.Delete().Exec(context.Background())

	if err != nil {
		return err
	}

	return nil
}
