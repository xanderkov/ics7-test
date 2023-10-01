package repo

import (
	"context"
	"hospital/internal/modules/db"
	"hospital/internal/modules/db/ent"
	"hospital/internal/modules/db/ent/doctor"
	"hospital/internal/modules/domain/doctor/dto"
)

type DoctorRepo struct {
	client *ent.Client
}

func NewDoctorRepo(client *ent.Client) *DoctorRepo {
	return &DoctorRepo{
		client: client,
	}
}

func (r *DoctorRepo) GetById(ctx context.Context, id int) (*dto.Doctor, error) {
	Doctor, err := r.client.Doctor.Get(ctx, id)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToDoctorDTO(Doctor), nil
}

func (r *DoctorRepo) GetByTokenId(ctx context.Context, token string) (*dto.Doctor, error) {
	Doctor, err := r.client.Doctor.Query().Where(doctor.TokenIdEQ(token)).Only(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToDoctorDTO(Doctor), nil
}

func (r *DoctorRepo) List(ctx context.Context) (dto.Doctors, error) {
	Doctors, err := r.client.Doctor.Query().All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToDoctorDTOs(Doctors), nil
}

func (r *DoctorRepo) Create(ctx context.Context, dtm *dto.CreateDoctor) (*dto.Doctor, error) {
	Doctor, err := r.client.Doctor.Create().
		SetSurname(dtm.Surname).
		SetTokenId(dtm.TokenId).
		SetRole(dtm.Role).
		SetSpeciality(dtm.Speciality).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToDoctorDTO(Doctor), nil
}

func (r *DoctorRepo) Update(ctx context.Context, id int, dtm *dto.UpdateDoctor) (*dto.Doctor, error) {
	Doctor, err := r.client.Doctor.UpdateOneID(id).
		SetTokenId(dtm.TokenId).
		SetSurname(dtm.Surname).
		SetRole(dtm.Role).
		SetSpeciality(dtm.Speciality).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToDoctorDTO(Doctor), nil
}

func (r *DoctorRepo) Delete(ctx context.Context, id int) error {
	err := r.client.Doctor.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return db.WrapError(err)
	}

	return nil
}

func (r *DoctorRepo) Restore(ctx context.Context, id int) (*dto.Doctor, error) {
	Doctor, err := r.client.Doctor.UpdateOneID(id).Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToDoctorDTO(Doctor), nil
}

func ToDoctorDTO(model *ent.Doctor) *dto.Doctor {
	if model == nil {
		return nil
	}
	return &dto.Doctor{
		Id:         model.ID,
		TokenId:    model.TokenId,
		Surname:    model.Surname,
		Speciality: model.Speciality,
		Role:       model.Role,
	}
}

func ToDoctorDTOs(models ent.Doctors) dto.Doctors {
	if models == nil {
		return nil
	}
	dtms := make(dto.Doctors, len(models))
	for i := range models {
		dtms[i] = ToDoctorDTO(models[i])
	}
	return dtms
}
