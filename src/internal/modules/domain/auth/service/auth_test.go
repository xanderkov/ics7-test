package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"hospital/internal/modules/config"
	"hospital/internal/modules/domain/auth/dto"
	doctor_dto "hospital/internal/modules/domain/doctor/dto"
	"reflect"
	"testing"
)

func TestAuthService_SignUp(t *testing.T) {
	type fields struct {
		repo IDoctorRepo
	}

	type args struct {
		ctx context.Context
		dtm *dto.NewDoctor
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIDoctorRepo(ctrl)

	// Test case 1: Successful create
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    *doctor_dto.Doctor
		wantErr bool
	}{
		name: "Successful create",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			dtm: &dto.NewDoctor{
				Surname:    "Doe",
				TokenId:    "1",
				Speciality: "Doctor",
				Role:       "Role",
			},
		},
		want: &doctor_dto.Doctor{
			Surname:    "Doe",
			TokenId:    "1",
			Speciality: "Doctor",
			Role:       "Role",
		},
		wantErr: false,
	}

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&doctor_dto.Doctor{
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
		want    *doctor_dto.Doctor
		wantErr bool
	}{
		name: "Error while creating doctor",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			dtm: &dto.NewDoctor{
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
		want    *doctor_dto.Doctor
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		runner.Run(t, tt.name, func(t provider.T) {
			r := &AuthService{
				repo: tt.fields.repo,
			}
			got, err := r.SignUp(tt.args.ctx, tt.args.dtm)
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

func TestNewAuthService(t *testing.T) {
	type args struct {
		repo   IDoctorRepo
		config config.Config
	}
	mockDoctor := new(MockIDoctorRepo)

	tests := []struct {
		name string
		args args
		want *AuthService
	}{
		{
			name: "Simple positive test",
			args: args{
				repo: mockDoctor,
			},
			want: &AuthService{
				repo: mockDoctor,
			},
		},
	}
	for _, tt := range tests {
		runner.Run(t, tt.name, func(t provider.T) {
			if got := NewAuthService(tt.args.repo, tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthService() = %v, want %v", got, tt.want)
			}
		})
	}
}
