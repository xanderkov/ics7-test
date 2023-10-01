package repo

import (
	"context"
	"hospital/internal/modules/db"
	"hospital/internal/modules/db/ent"
	"hospital/internal/modules/domain/patient/dto"
)

type PatientRepo struct {
	client *ent.Client
}

func NewPatientRepo(client *ent.Client) *PatientRepo {
	return &PatientRepo{
		client: client,
	}
}

func (r *PatientRepo) GetById(ctx context.Context, id int) (*dto.Patient, error) {
	Patient, err := r.client.Patient.Get(ctx, id)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToPatientDTO(Patient), nil
}

func (r *PatientRepo) List(ctx context.Context) (dto.Patients, error) {
	Patients, err := r.client.Patient.Query().All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToPatientDTOs(Patients), nil
}

func (r *PatientRepo) Create(ctx context.Context, dtm *dto.CreatePatient) (*dto.Patient, error) {
	Patient, err := r.client.Patient.Create().
		SetName(dtm.Name).
		SetHeight(dtm.Height).
		SetPatronymic(dtm.Patronymic).
		SetDegreeOfDanger(dtm.DegreeOfDanger).
		SetSurname(dtm.Surname).
		SetWeight(dtm.Weight).
		SetRoomNumber(dtm.RoomNumber).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToPatientDTO(Patient), nil
}

func (r *PatientRepo) Update(ctx context.Context, id int, dtm *dto.UpdatePatient) (*dto.Patient, error) {
	Patient, err := r.client.Patient.UpdateOneID(id).
		SetName(dtm.Name).
		SetHeight(dtm.Height).
		SetPatronymic(dtm.Patronymic).
		SetDegreeOfDanger(dtm.DegreeOfDanger).
		SetSurname(dtm.Surname).
		SetWeight(dtm.Weight).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToPatientDTO(Patient), nil
}

func (r *PatientRepo) Delete(ctx context.Context, id int) error {
	err := r.client.Patient.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return db.WrapError(err)
	}

	return nil
}

func (r *PatientRepo) Restore(ctx context.Context, id int) (*dto.Patient, error) {
	Patient, err := r.client.Patient.UpdateOneID(id).Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToPatientDTO(Patient), nil
}

func ToPatientDTO(model *ent.Patient) *dto.Patient {
	if model == nil {
		return nil
	}
	return &dto.Patient{
		Id:             model.ID,
		Name:           model.Name,
		Surname:        model.Surname,
		Patronymic:     model.Patronymic,
		Height:         model.Height,
		DegreeOfDanger: model.DegreeOfDanger,
		Weight:         model.Weight,
		RoomNumber:     model.RoomNumber,
	}
}

func ToPatientDTOs(models ent.Patients) dto.Patients {
	if models == nil {
		return nil
	}
	dtms := make(dto.Patients, len(models))
	for i := range models {
		dtms[i] = ToPatientDTO(models[i])
	}
	return dtms
}
