package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"hospital/internal/modules/domain/patient/dto"

	"reflect"
	"testing"
)

func TestNewPatientService(t *testing.T) {
	type args struct {
		repo IPatientRepo
	}
	mockPatient := new(MockIPatientRepo)

	tests := []struct {
		name string
		args args
		want *PatientService
	}{
		{
			name: "Simple positive test",
			args: args{
				repo: mockPatient,
			},
			want: &PatientService{
				repo: mockPatient,
			},
		},
	}
	for _, tt := range tests {
		runner.Run(t, tt.name, func(t provider.T) {
			if got := NewPatientService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPatientService() = %v, want %v", got, tt.want)
			}
		})

	}
}

func TestPatientService_Create(t *testing.T) {
	type fields struct {
		repo IPatientRepo
	}

	type args struct {
		ctx context.Context
		dtm *dto.CreatePatient
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIPatientRepo(ctrl)

	// Test case 1: Successful create
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Patient
		wantErr bool
	}{
		name: "Successful create",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			dtm: &dto.CreatePatient{
				Surname:        "Doe",
				Name:           "John",
				Patronymic:     "Smith",
				Height:         180,
				Weight:         75.5,
				RoomNumber:     101,
				DegreeOfDanger: 2,
			},
		},
		want: &dto.Patient{
			Surname:        "Doe",
			Name:           "John",
			Patronymic:     "Smith",
			Height:         180,
			Weight:         75.5,
			RoomNumber:     101,
			DegreeOfDanger: 2,
		},
		wantErr: false,
	}

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&dto.Patient{
		Surname:        "Doe",
		Name:           "John",
		Patronymic:     "Smith",
		Height:         180,
		Weight:         75.5,
		RoomNumber:     101,
		DegreeOfDanger: 2,
	}, nil)

	// Test case 2: Error while creating patient
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Patient
		wantErr bool
	}{
		name: "Error while creating patient",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			dtm: &dto.CreatePatient{
				Surname:        "Doe",
				Name:           "John",
				Patronymic:     "Smith",
				Height:         180,
				Weight:         75.5,
				RoomNumber:     101,
				DegreeOfDanger: 2,
			},
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("error while creating patient"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Patient
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		//runner.Run(t, tt.name, func(t provider.T) {
		//	r := &PatientService{
		//		repo: tt.fields.repo,
		//	}
		//	got, err := r.Create(tt.args.ctx, tt.args.dtm)
		//	if (err != nil) != tt.wantErr {
		//		t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
		//		return
		//	}
		//	if !reflect.DeepEqual(got, tt.want) {
		//		t.Errorf("Create() got = %v, want %v", got, tt.want)
		//	}
		//})

		runner.Run(t, tt.name, func(t provider.T) {
			r := &PatientService{
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

func TestPatientService_Delete(t *testing.T) {
	type fields struct {
		repo IPatientRepo
	}
	type args struct {
		ctx context.Context
		id  int
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIPatientRepo(ctrl)

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

	// Test case 2: Error while deleting patient
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		name: "Error while deleting patient",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  2,
		},
		wantErr: true,
	}

	mockRepo.EXPECT().Delete(gomock.Any(), testCase2.args.id).Return(errors.New("error while deleting patient"))

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
			r := &PatientService{
				repo: tt.fields.repo,
			}
			if err := r.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPatientService_GetById(t *testing.T) {
	type fields struct {
		repo IPatientRepo
	}
	type args struct {
		ctx context.Context
		id  int
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIPatientRepo(ctrl)

	// Test case 1: Successful get
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Patient
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
		want: &dto.Patient{
			Id:             1,
			Surname:        "Doe",
			Name:           "John",
			Patronymic:     "Smith",
			Height:         180,
			Weight:         75.5,
			RoomNumber:     101,
			DegreeOfDanger: 2,
		},
		wantErr: false,
	}

	mockRepo.EXPECT().GetById(gomock.Any(), testCase1.args.id).Return(testCase1.want, nil)

	// Test case 2: Error while getting patient
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Patient
		wantErr bool
	}{
		name: "Error while getting patient",
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

	mockRepo.EXPECT().GetById(gomock.Any(), testCase2.args.id).Return(nil, errors.New("error while getting patient"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Patient
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		runner.Run(t, tt.name, func(t provider.T) {
			r := &PatientService{
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

func TestPatientService_List(t *testing.T) {
	type fields struct {
		repo IPatientRepo
	}
	type args struct {
		ctx context.Context
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIPatientRepo(ctrl)

	// Test case 1: Successful list
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    dto.Patients
		wantErr bool
	}{
		name: "Successful list",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
		},
		want: dto.Patients{
			{
				Id:             1,
				Surname:        "Doe",
				Name:           "John",
				Patronymic:     "Smith",
				Height:         180,
				Weight:         75.5,
				RoomNumber:     101,
				DegreeOfDanger: 2,
			},
			{
				Id:             2,
				Surname:        "Doe",
				Name:           "Jane",
				Patronymic:     "Smith",
				Height:         170,
				Weight:         60.0,
				RoomNumber:     102,
				DegreeOfDanger: 1,
			},
		},
		wantErr: false,
	}

	mockRepo.EXPECT().List(gomock.Any()).Return(testCase1.want, nil)

	// Test case 2: Error while listing patients
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    dto.Patients
		wantErr bool
	}{
		name: "Error while listing patients",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().List(gomock.Any()).Return(nil, errors.New("error while listing patients"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		want    dto.Patients
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		runner.Run(t, tt.name, func(t provider.T) {
			r := &PatientService{
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

func TestPatientService_Update(t *testing.T) {
	type fields struct {
		repo IPatientRepo
	}
	type args struct {
		ctx context.Context
		id  int
		dtm *dto.UpdatePatient
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIPatientRepo(ctrl)

	// Test case 1: Successful update
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Patient
		wantErr bool
	}{
		name: "Successful update",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  1,
			dtm: &dto.UpdatePatient{
				Surname:        "Doe",
				Name:           "John",
				Patronymic:     "Smith",
				Height:         180,
				Weight:         75.5,
				RoomNumber:     101,
				DegreeOfDanger: 2,
			},
		},
		want: &dto.Patient{
			Id:             1,
			Surname:        "Doe",
			Name:           "John",
			Patronymic:     "Smith",
			Height:         180,
			Weight:         75.5,
			RoomNumber:     101,
			DegreeOfDanger: 2,
		},
		wantErr: false,
	}

	mockRepo.EXPECT().Update(gomock.Any(), testCase1.args.id, testCase1.args.dtm).Return(testCase1.want, nil)

	// Test case 2: Patient not found
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Patient
		wantErr bool
	}{
		name: "Patient not found",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  1,
			dtm: &dto.UpdatePatient{
				Surname:        "Doe",
				Name:           "John",
				Patronymic:     "Smith",
				Height:         180,
				Weight:         75.5,
				RoomNumber:     101,
				DegreeOfDanger: 2,
			},
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().Update(gomock.Any(), testCase2.args.id, testCase2.args.dtm).Return(nil, errors.New("patient not found"))

	// Test case 3: Invalid update data
	testCase3 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Patient
		wantErr bool
	}{
		name: "Invalid update data",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  1,
			dtm: &dto.UpdatePatient{
				Surname: "",
				Name:    "",
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
		want    *dto.Patient
		wantErr bool
	}{
		testCase1,
		testCase2,
		testCase3,
	} {
		runner.Run(t, tt.name, func(t provider.T) {
			r := &PatientService{
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
