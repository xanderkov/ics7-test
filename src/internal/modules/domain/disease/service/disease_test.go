package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"hospital/internal/modules/domain/disease/dto"
	"reflect"
	"testing"
)

func TestNewDiseaseService(t *testing.T) {
	type args struct {
		repo IDiseaseRepo
	}
	mockDisease := new(MockIDiseaseRepo)

	tests := []struct {
		name string
		args args
		want *DiseaseService
	}{
		{
			name: "Simple positive test",
			args: args{
				repo: mockDisease,
			},
			want: &DiseaseService{
				repo: mockDisease,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDiseaseService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDiseaseService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiseaseService_Create(t *testing.T) {
	type fields struct {
		repo IDiseaseRepo
	}

	type args struct {
		ctx context.Context
		dtm *dto.CreateDisease
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIDiseaseRepo(ctrl)

	// Test case 1: Successful create
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Disease
		wantErr bool
	}{
		name: "Successful create",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			dtm: &dto.CreateDisease{
				Threat:         "Doe",
				Name:           "John",
				DegreeOfDanger: 2,
			},
		},
		want: &dto.Disease{
			Threat:         "Doe",
			Name:           "John",
			DegreeOfDanger: 2,
		},
		wantErr: false,
	}

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&dto.Disease{
		Threat:         "Doe",
		Name:           "John",
		DegreeOfDanger: 2,
	}, nil)

	// Test case 2: Error while creating disease
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Disease
		wantErr bool
	}{
		name: "Error while creating disease",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			dtm: &dto.CreateDisease{
				Threat:         "Doe",
				Name:           "John",
				DegreeOfDanger: 2,
			},
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("error while creating disease"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Disease
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := &DiseaseService{
				repo: tt.fields.repo,
			}
			got, err := r.Create(tt.args.ctx, tt.args.dtm)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiseaseService_Delete(t *testing.T) {
	type fields struct {
		repo IDiseaseRepo
	}
	type args struct {
		ctx context.Context
		id  int
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIDiseaseRepo(ctrl)

	// Test case 1: Successful delete
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		name: "Successful delete",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  1,
		},
		wantErr: false,
	}

	mockRepo.EXPECT().Delete(gomock.Any(), testCase1.args.id).Return(nil)

	// Test case 2: Error while deleting disease
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		name: "Error while deleting disease",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  2,
		},
		wantErr: true,
	}

	mockRepo.EXPECT().Delete(gomock.Any(), testCase2.args.id).Return(errors.New("error while deleting disease"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := &DiseaseService{
				repo: tt.fields.repo,
			}
			if err := r.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDiseaseService_GetById(t *testing.T) {
	type fields struct {
		repo IDiseaseRepo
	}
	type args struct {
		ctx context.Context
		id  int
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIDiseaseRepo(ctrl)

	// Test case 1: Successful get
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Disease
		wantErr bool
	}{
		name: "Successful get",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  1,
		},
		want: &dto.Disease{
			Threat:         "Doe",
			Name:           "John",
			DegreeOfDanger: 2,
		},
		wantErr: false,
	}

	mockRepo.EXPECT().GetById(gomock.Any(), testCase1.args.id).Return(testCase1.want, nil)

	// Test case 2: Error while getting disease
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Disease
		wantErr bool
	}{
		name: "Error while getting disease",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  2,
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().GetById(gomock.Any(), testCase2.args.id).Return(nil, errors.New("error while getting disease"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Disease
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := &DiseaseService{
				repo: tt.fields.repo,
			}
			got, err := r.GetById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiseaseService_List(t *testing.T) {
	type fields struct {
		repo IDiseaseRepo
	}
	type args struct {
		ctx context.Context
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIDiseaseRepo(ctrl)

	// Test case 1: Successful list
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    dto.Diseases
		wantErr bool
	}{
		name: "Successful list",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
		},
		want: dto.Diseases{
			{
				Threat:         "Doe",
				Name:           "John",
				DegreeOfDanger: 2,
			},
			{
				Threat:         "Doe",
				Name:           "John",
				DegreeOfDanger: 2,
			},
		},
		wantErr: false,
	}

	mockRepo.EXPECT().List(gomock.Any()).Return(testCase1.want, nil)

	// Test case 2: Error while listing diseases
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    dto.Diseases
		wantErr bool
	}{
		name: "Error while listing diseases",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().List(gomock.Any()).Return(nil, errors.New("error while listing diseases"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		want    dto.Diseases
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := &DiseaseService{
				repo: tt.fields.repo,
			}
			got, err := r.List(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiseaseService_Update(t *testing.T) {
	type fields struct {
		repo IDiseaseRepo
	}
	type args struct {
		ctx context.Context
		id  int
		dtm *dto.UpdateDisease
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIDiseaseRepo(ctrl)

	// Test case 1: Successful update
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Disease
		wantErr bool
	}{
		name: "Successful update",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  1,
			dtm: &dto.UpdateDisease{
				Threat:         "Doe",
				Name:           "John",
				DegreeOfDanger: 2,
			},
		},
		want: &dto.Disease{
			Threat:         "Doe",
			Name:           "John",
			DegreeOfDanger: 2,
		},
		wantErr: false,
	}

	mockRepo.EXPECT().Update(gomock.Any(), testCase1.args.id, testCase1.args.dtm).Return(testCase1.want, nil)

	// Test case 2: Disease not found
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Disease
		wantErr bool
	}{
		name: "Disease not found",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  1,
			dtm: &dto.UpdateDisease{
				Threat:         "Doe",
				Name:           "John",
				DegreeOfDanger: 2,
			},
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().Update(gomock.Any(), testCase2.args.id, testCase2.args.dtm).Return(nil, errors.New("disease not found"))

	// Test case 3: Invalid update data
	testCase3 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Disease
		wantErr bool
	}{
		name: "Invalid update data",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  1,
			dtm: &dto.UpdateDisease{
				Threat:         "Doe",
				Name:           "John",
				DegreeOfDanger: 2,
			},
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().Update(gomock.Any(), testCase3.args.id, testCase3.args.dtm).Return(nil, errors.New("invalid update data"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Disease
		wantErr bool
	}{
		testCase1,
		testCase2,
		testCase3,
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := &DiseaseService{
				repo: tt.fields.repo,
			}
			got, err := r.Update(tt.args.ctx, tt.args.id, tt.args.dtm)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}
