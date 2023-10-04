package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"hospital/internal/modules/domain/doctor/dto"
	"reflect"
	"testing"
)

func TestNewDoctorService(t *testing.T) {
	type args struct {
		repo IDoctorRepo
	}
	mockDoctor := new(MockIDoctorRepo)

	tests := []struct {
		name string
		args args
		want *DoctorService
	}{
		{
			name: "Simple positive test",
			args: args{
				repo: mockDoctor,
			},
			want: &DoctorService{
				repo: mockDoctor,
			},
		},
	}
	for _, tt := range tests {
		runner.Run(t, tt.name, func(t provider.T) {
			if got := NewDoctorService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDoctorService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoctorService_Create(t *testing.T) {
	type fields struct {
		repo IDoctorRepo
	}

	type args struct {
		ctx context.Context
		dtm *dto.CreateDoctor
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIDoctorRepo(ctrl)

	// Test case 1: Successful create
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Doctor
		wantErr bool
	}{
		name: "Successful create",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			dtm: &dto.CreateDoctor{
				Surname:    "Doe",
				TokenId:    "1",
				Speciality: "Doctor",
				Role:       "Role",
			},
		},
		want: &dto.Doctor{
			Surname:    "Doe",
			TokenId:    "1",
			Speciality: "Doctor",
			Role:       "Role",
		},
		wantErr: false,
	}

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&dto.Doctor{
		Surname:    "Doe",
		TokenId:    "1",
		Speciality: "Doctor",
		Role:       "Role",
	}, nil)

	// Test case 2: Error while creating doctor
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Doctor
		wantErr bool
	}{
		name: "Error while creating doctor",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			dtm: &dto.CreateDoctor{
				Surname:    "Doe",
				TokenId:    "1",
				Speciality: "Doctor",
				Role:       "Role",
			},
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("error while creating doctor"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Doctor
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		runner.Run(t, tt.name, func(t provider.T) {
			r := &DoctorService{
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

func TestDoctorService_Delete(t *testing.T) {
	type fields struct {
		repo IDoctorRepo
	}
	type args struct {
		ctx context.Context
		id  int
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIDoctorRepo(ctrl)

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

	// Test case 2: Error while deleting doctor
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		name: "Error while deleting doctor",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  2,
		},
		wantErr: true,
	}

	mockRepo.EXPECT().Delete(gomock.Any(), testCase2.args.id).Return(errors.New("error while deleting doctor"))

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
		runner.Run(t, tt.name, func(t provider.T) {
			r := &DoctorService{
				repo: tt.fields.repo,
			}
			if err := r.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoctorService_GetById(t *testing.T) {
	type fields struct {
		repo IDoctorRepo
	}
	type args struct {
		ctx context.Context
		id  int
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIDoctorRepo(ctrl)

	// Test case 1: Successful get
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Doctor
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
		want: &dto.Doctor{
			Surname:    "Doe",
			TokenId:    "1",
			Speciality: "Doctor",
			Role:       "Role",
		},
		wantErr: false,
	}

	mockRepo.EXPECT().GetById(gomock.Any(), testCase1.args.id).Return(testCase1.want, nil)

	// Test case 2: Error while getting doctor
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Doctor
		wantErr bool
	}{
		name: "Error while getting doctor",
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

	mockRepo.EXPECT().GetById(gomock.Any(), testCase2.args.id).Return(nil, errors.New("error while getting doctor"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Doctor
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		runner.Run(t, tt.name, func(t provider.T) {
			r := &DoctorService{
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

func TestDoctorService_List(t *testing.T) {
	type fields struct {
		repo IDoctorRepo
	}
	type args struct {
		ctx context.Context
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIDoctorRepo(ctrl)

	// Test case 1: Successful list
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    dto.Doctors
		wantErr bool
	}{
		name: "Successful list",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
		},
		want: dto.Doctors{
			{
				Id:         1,
				Surname:    "Doe",
				TokenId:    "1",
				Speciality: "Doctor",
				Role:       "Role",
			},
			{
				Id:         2,
				Surname:    "Doe",
				TokenId:    "1",
				Speciality: "Doctor",
				Role:       "Role",
			},
		},
		wantErr: false,
	}

	mockRepo.EXPECT().List(gomock.Any()).Return(testCase1.want, nil)

	// Test case 2: Error while listing doctors
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    dto.Doctors
		wantErr bool
	}{
		name: "Error while listing doctors",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().List(gomock.Any()).Return(nil, errors.New("error while listing doctors"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		want    dto.Doctors
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		runner.Run(t, tt.name, func(t provider.T) {
			r := &DoctorService{
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

func TestDoctorService_Update(t *testing.T) {
	type fields struct {
		repo IDoctorRepo
	}
	type args struct {
		ctx context.Context
		id  int
		dtm *dto.UpdateDoctor
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIDoctorRepo(ctrl)

	// Test case 1: Successful update
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Doctor
		wantErr bool
	}{
		name: "Successful update",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  1,
			dtm: &dto.UpdateDoctor{
				Surname:    "Doe",
				TokenId:    "1",
				Speciality: "Doctor",
				Role:       "Role",
			},
		},
		want: &dto.Doctor{
			Surname:    "Doe",
			TokenId:    "1",
			Speciality: "Doctor",
			Role:       "Role",
		},
		wantErr: false,
	}

	mockRepo.EXPECT().Update(gomock.Any(), testCase1.args.id, testCase1.args.dtm).Return(testCase1.want, nil)

	// Test case 2: Doctor not found
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Doctor
		wantErr bool
	}{
		name: "Doctor not found",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  1,
			dtm: &dto.UpdateDoctor{
				Surname:    "Doe",
				TokenId:    "1",
				Speciality: "Doctor",
				Role:       "Role",
			},
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().Update(gomock.Any(), testCase2.args.id, testCase2.args.dtm).Return(nil, errors.New("doctor not found"))

	// Test case 3: Invalid update data
	testCase3 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Doctor
		wantErr bool
	}{
		name: "Invalid update data",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  1,
			dtm: &dto.UpdateDoctor{
				Surname:    "Doe",
				TokenId:    "1",
				Speciality: "Doctor",
				Role:       "Role",
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
		want    *dto.Doctor
		wantErr bool
	}{
		testCase1,
		testCase2,
		testCase3,
	} {
		runner.Run(t, tt.name, func(t provider.T) {
			r := &DoctorService{
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

func TestDoctorService_GetByTokenId(t *testing.T) {
	type fields struct {
		repo IDoctorRepo
	}
	type args struct {
		ctx     context.Context
		tokenId string
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIDoctorRepo(ctrl)

	// Test case 1: Successful get
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Doctor
		wantErr bool
	}{
		name: "Successful get",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx:     context.Background(),
			tokenId: "1",
		},
		want: &dto.Doctor{
			Surname:    "Doe",
			TokenId:    "1",
			Speciality: "Doctor",
			Role:       "Role",
		},
		wantErr: false,
	}

	mockRepo.EXPECT().GetByTokenId(gomock.Any(), testCase1.args.tokenId).Return(testCase1.want, nil)

	// Test case 2: Error while getting doctor
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Doctor
		wantErr bool
	}{
		name: "Error while getting doctor",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx:     context.Background(),
			tokenId: "2",
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().GetByTokenId(gomock.Any(), testCase2.args.tokenId).Return(nil, errors.New("error while getting doctor"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Doctor
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		runner.Run(t, tt.name, func(t provider.T) {
			r := &DoctorService{
				repo: tt.fields.repo,
			}
			got, err := r.GetByTokenId(tt.args.ctx, tt.args.tokenId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByTokenId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByTokenId() got = %v, want %v", got, tt.want)
			}
		})
	}
}
