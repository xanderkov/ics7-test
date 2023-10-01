package repo

import (
	"context"
	"hospital/internal/modules/db"
	"hospital/internal/modules/db/ent"
	"hospital/internal/modules/domain/disease/dto"
)

type DiseaseRepo struct {
	client *ent.Client
}

func NewDiseaseRepo(client *ent.Client) *DiseaseRepo {
	return &DiseaseRepo{
		client: client,
	}
}

func (r *DiseaseRepo) GetById(ctx context.Context, id int) (*dto.Disease, error) {
	Disease, err := r.client.Disease.Get(ctx, id)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToDiseaseDTO(Disease), nil
}

func (r *DiseaseRepo) List(ctx context.Context) (dto.Diseases, error) {
	Diseases, err := r.client.Disease.Query().All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToDiseaseDTOs(Diseases), nil
}

func (r *DiseaseRepo) Create(ctx context.Context, dtm *dto.CreateDisease) (*dto.Disease, error) {
	Disease, err := r.client.Disease.Create().
		SetName(dtm.Name).
		SetDegreeOfDanger(dtm.DegreeOfDanger).
		SetThreat(dtm.Threat).
		Save(ctx)

	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToDiseaseDTO(Disease), nil
}

func (r *DiseaseRepo) Update(ctx context.Context, id int, dtm *dto.UpdateDisease) (*dto.Disease, error) {
	Disease, err := r.client.Disease.UpdateOneID(id).
		SetName(dtm.Name).
		SetDegreeOfDanger(dtm.DegreeOfDanger).
		SetThreat(dtm.Threat).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToDiseaseDTO(Disease), nil
}

func (r *DiseaseRepo) Delete(ctx context.Context, id int) error {
	err := r.client.Disease.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return db.WrapError(err)
	}

	return nil
}

func (r *DiseaseRepo) Restore(ctx context.Context, id int) (*dto.Disease, error) {
	Disease, err := r.client.Disease.UpdateOneID(id).Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToDiseaseDTO(Disease), nil
}

func ToDiseaseDTO(model *ent.Disease) *dto.Disease {
	if model == nil {
		return nil
	}
	return &dto.Disease{
		Id:             model.ID,
		Name:           model.Name,
		DegreeOfDanger: model.DegreeOfDanger,
		Threat:         model.Threat,
	}
}

func ToDiseaseDTOs(models ent.Diseases) dto.Diseases {
	if models == nil {
		return nil
	}
	dtms := make(dto.Diseases, len(models))
	for i := range models {
		dtms[i] = ToDiseaseDTO(models[i])
	}
	return dtms
}
