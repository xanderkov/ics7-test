package repo

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"hospital/internal/modules/config"
	"hospital/internal/modules/db"
	"hospital/internal/modules/db/ent"
	"hospital/internal/modules/domain/doctor/dto"
	"hospital/internal/modules/logger"
	"reflect"
	"testing"
)

func TestNewDoctorRepo(t *testing.T) {
	type args struct {
		client *ent.Client
	}
	client := ent.NewClient()
	// Run the test cases
	for _, tt := range []struct {
		name string
		args args
		want *DoctorRepo
	}{
		{
			name: "NewDoctorRepo",
			args: args{
				client: client,
			},
			want: &DoctorRepo{
				client: client,
			},
		},
	} {
		runner.Run(t, tt.name, func(t provider.T) {
			if got := NewDoctorRepo(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDoctorRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoctorRepo_Create(t *testing.T) {
	log, levelog, err := logger.NewLogger()

	if err != nil {
		t.Fatalf("failed to init logger: %v", err)
	}
	cfg, err := config.NewConfig(log, levelog)
	if err != nil {
		t.Fatalf("failed to init config: %v", err)
	}
	// Create a new in-memory database for testing
	client, err := db.NewDBClient(cfg, log)
	err = db.TruncateAll(client)
	if err != nil {
		t.Fatalf("failed to truncate all: %v", err)
	}
	// create a new doctor
	doctor, err := client.Doctor.Create().
		SetSurname("Kovel").
		SetSpeciality("Doctor").
		SetRole("Doctor").
		SetTokenId("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}

	// Create a new doctor repository
	repo := NewDoctorRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DoctorRepo
		id      int
		want    *dto.Doctor
		wantErr bool
	}{
		name: "Successful get by ID",
		repo: repo,
		id:   doctor.ID,
		want: &dto.Doctor{
			Id:         doctor.ID,
			Surname:    "Kovel",
			Speciality: "Doctor",
			Role:       "Doctor",
			TokenId:    "1",
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got, err := testCase1.repo.GetById(context.Background(), testCase1.id)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("Create() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("Create() got = %v, want %v", got, testCase1.want)
		}
	})

	// Test case 2: Get by ID with non-existent doctor
	testCase2 := struct {
		name    string
		repo    *DoctorRepo
		id      int
		want    *dto.Doctor
		wantErr bool
	}{
		name:    "Get by ID with non-existent doctor",
		repo:    repo,
		id:      100,
		want:    nil,
		wantErr: true,
	}

	// Run the test case 2
	runner.Run(t, testCase2.name, func(t provider.T) {
		got, err := testCase2.repo.GetById(context.Background(), testCase2.id)
		if (err != nil) != testCase2.wantErr {
			t.Errorf("GetById() error = %v, wantErr %v", err, testCase2.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase2.want) {
			t.Errorf("GetById() got = %v, want %v", got, testCase2.want)
		}
	})
}

func TestDoctorRepo_Delete(t *testing.T) {
	log, levelog, err := logger.NewLogger()

	if err != nil {
		t.Fatalf("failed to init logger: %v", err)
	}
	cfg, err := config.NewConfig(log, levelog)
	if err != nil {
		t.Fatalf("failed to init config: %v", err)
	}
	// Create a new in-memory database for testing
	client, err := db.NewDBClient(cfg, log)
	err = db.TruncateAll(client)
	if err != nil {
		t.Fatalf("failed to truncate all: %v", err)
	}
	// create a new doctor
	doctor, err := client.Doctor.Create().
		SetSurname("Kovel").
		SetSpeciality("Doctor").
		SetRole("Doctor").
		SetTokenId("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}
	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to update doctor: %v", err)
	}

	// Create a new doctor repository
	repo := NewDoctorRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DoctorRepo
		id      int
		want    *dto.Doctor
		wantErr bool
	}{
		name: "Successful Deleted",
		repo: repo,
		id:   doctor.ID,
		want: &dto.Doctor{
			Id:         doctor.ID,
			Surname:    "Kovel",
			Speciality: "Doctor",
			Role:       "Doctor",
			TokenId:    "1",
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		err := testCase1.repo.Delete(context.Background(), testCase1.id)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("Delete() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
	})

	// Test case 2: Get by ID with non-existent doctor
	testCase2 := struct {
		name    string
		repo    *DoctorRepo
		id      int
		want    *dto.Doctor
		wantErr bool
	}{
		name:    "Delete by ID with non-existent doctor",
		repo:    repo,
		id:      100,
		want:    nil,
		wantErr: true,
	}

	// Run the test case 2
	runner.Run(t, testCase2.name, func(t provider.T) {
		err := testCase2.repo.Delete(context.Background(), testCase2.id)
		if (err != nil) != testCase2.wantErr {
			t.Errorf("Delete() error = %v, wantErr %v", err, testCase2.wantErr)
			return
		}
	})
}

func TestDoctorRepo_GetById(t *testing.T) {
	log, levelog, err := logger.NewLogger()

	if err != nil {
		t.Fatalf("failed to init logger: %v", err)
	}
	cfg, err := config.NewConfig(log, levelog)
	if err != nil {
		t.Fatalf("failed to init config: %v", err)
	}
	// Create a new in-memory database for testing
	client, err := db.NewDBClient(cfg, log)
	err = db.TruncateAll(client)
	if err != nil {
		t.Fatalf("failed to truncate all: %v", err)
	}
	// create a new doctor
	doctor, err := client.Doctor.Create().
		SetSurname("Kovel").
		SetSpeciality("Doctor").
		SetRole("Doctor").
		SetTokenId("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}
	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}

	// Create a new doctor repository
	repo := NewDoctorRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DoctorRepo
		id      int
		want    *dto.Doctor
		wantErr bool
	}{
		name: "Successful get by ID",
		repo: repo,
		id:   doctor.ID,
		want: &dto.Doctor{
			Id:         doctor.ID,
			Surname:    "Kovel",
			Speciality: "Doctor",
			Role:       "Doctor",
			TokenId:    "1",
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got, err := testCase1.repo.GetById(context.Background(), testCase1.id)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("GetById() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("GetById() got = %v, want %v", got, testCase1.want)
		}
	})

	// Test case 2: Get by ID with non-existent doctor
	testCase2 := struct {
		name    string
		repo    *DoctorRepo
		id      int
		want    *dto.Doctor
		wantErr bool
	}{
		name:    "Get by ID with non-existent doctor",
		repo:    repo,
		id:      100,
		want:    nil,
		wantErr: true,
	}

	// Run the test case 2
	runner.Run(t, testCase2.name, func(t provider.T) {
		got, err := testCase2.repo.GetById(context.Background(), testCase2.id)
		if (err != nil) != testCase2.wantErr {
			t.Errorf("GetById() error = %v, wantErr %v", err, testCase2.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase2.want) {
			t.Errorf("GetById() got = %v, want %v", got, testCase2.want)
		}
	})
}

func TestDoctorRepo_List(t *testing.T) {
	log, levelog, err := logger.NewLogger()

	if err != nil {
		t.Fatalf("failed to init logger: %v", err)
	}
	cfg, err := config.NewConfig(log, levelog)
	if err != nil {
		t.Fatalf("failed to init config: %v", err)
	}
	// Create a new in-memory database for testing
	client, err := db.NewDBClient(cfg, log)
	err = db.TruncateAll(client)
	if err != nil {
		t.Fatalf("failed to truncate all: %v", err)
	}
	// create a new doctor
	doctor, err := client.Doctor.Create().
		SetSurname("Kovel").
		SetSpeciality("Doctor").
		SetRole("Doctor").
		SetTokenId("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}

	// Create a new doctor repository
	repo := NewDoctorRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DoctorRepo
		id      int
		want    dto.Doctors
		wantErr bool
	}{
		name: "Successful list",
		repo: repo,
		id:   doctor.ID,
		want: dto.Doctors{
			&dto.Doctor{
				Id:         doctor.ID,
				Surname:    "Kovel",
				Speciality: "Doctor",
				Role:       "Doctor",
				TokenId:    "1",
			},
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got, err := testCase1.repo.List(context.Background())
		if (err != nil) != testCase1.wantErr {
			t.Errorf("List() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("List() got = %v, want %v", got, testCase1.want)
		}
	})

}

func TestDoctorRepo_Restore(t *testing.T) {
	log, levelog, err := logger.NewLogger()

	if err != nil {
		t.Fatalf("failed to init logger: %v", err)
	}
	cfg, err := config.NewConfig(log, levelog)
	if err != nil {
		t.Fatalf("failed to init config: %v", err)
	}
	// Create a new in-memory database for testing
	client, err := db.NewDBClient(cfg, log)
	err = db.TruncateAll(client)
	if err != nil {
		t.Fatalf("failed to truncate all: %v", err)
	}
	// create a new doctor
	doctor, err := client.Doctor.Create().
		SetSurname("Kovel").
		SetSpeciality("Doctor").
		SetRole("Doctor").
		SetTokenId("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to update doctor: %v", err)
	}

	// Create a new doctor repository
	repo := NewDoctorRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DoctorRepo
		id      int
		want    *dto.Doctor
		wantErr bool
	}{
		name: "Successful Restored",
		repo: repo,
		id:   doctor.ID,
		want: &dto.Doctor{
			Id:         doctor.ID,
			Surname:    "Kovel",
			Speciality: "Doctor",
			Role:       "Doctor",
			TokenId:    "1",
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got, err := testCase1.repo.Restore(context.Background(), testCase1.id)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("Restore() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("Restore() got = %v, want %v", got, testCase1.want)
		}
	})

	// Test case 2: Get by ID with non-existent doctor
	testCase2 := struct {
		name    string
		repo    *DoctorRepo
		id      int
		want    *dto.Doctor
		wantErr bool
	}{
		name:    "Resotre by ID with non-existent doctor",
		repo:    repo,
		id:      100,
		want:    nil,
		wantErr: true,
	}

	// Run the test case 2
	runner.Run(t, testCase2.name, func(t provider.T) {
		got, err := testCase2.repo.Restore(context.Background(), testCase2.id)
		if (err != nil) != testCase2.wantErr {
			t.Errorf("Restore() error = %v, wantErr %v", err, testCase2.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase2.want) {
			t.Errorf("Restore() got = %v, want %v", got, testCase2.want)
		}
	})
}

func TestDoctorRepo_Update(t *testing.T) {
	log, levelog, err := logger.NewLogger()

	if err != nil {
		t.Fatalf("failed to init logger: %v", err)
	}
	cfg, err := config.NewConfig(log, levelog)
	if err != nil {
		t.Fatalf("failed to init config: %v", err)
	}
	// Create a new in-memory database for testing
	client, err := db.NewDBClient(cfg, log)
	err = db.TruncateAll(client)
	if err != nil {
		t.Fatalf("failed to truncate all: %v", err)
	}
	// create a new doctor
	doctor, err := client.Doctor.Create().
		SetSurname("Kovel").
		SetSpeciality("Doctor").
		SetRole("Doctor").
		SetTokenId("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}

	_, err = client.Doctor.UpdateOneID(doctor.ID).
		SetSurname("Kovel").
		SetSpeciality("Doctor").
		SetRole("Doctor").
		SetTokenId("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to update doctor: %v", err)
	}

	// Create a new doctor repository
	repo := NewDoctorRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DoctorRepo
		id      int
		want    *dto.Doctor
		wantErr bool
	}{
		name: "Successful Updated",
		repo: repo,
		id:   doctor.ID,
		want: &dto.Doctor{
			Id:         doctor.ID,
			Surname:    "Kovel",
			Speciality: "Doctor",
			Role:       "Doctor",
			TokenId:    "1",
		},
		wantErr: false,
	}
	upd_doctor := dto.UpdateDoctor{
		Surname:    "Kovel",
		Speciality: "Doctor",
		Role:       "Doctor",
		TokenId:    "1",
	}
	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got, err := testCase1.repo.Update(context.Background(), testCase1.id, &upd_doctor)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("Update() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("Update() got = %v, want %v", got, testCase1.want)
		}
	})

	// Test case 2: Get by ID with non-existent doctor
	testCase2 := struct {
		name    string
		repo    *DoctorRepo
		id      int
		want    *dto.Doctor
		wantErr bool
	}{
		name:    "Get by ID with non-existent doctor",
		repo:    repo,
		id:      100,
		want:    nil,
		wantErr: true,
	}

	// Run the test case 2
	runner.Run(t, testCase2.name, func(t provider.T) {
		got, err := testCase2.repo.Update(context.Background(), testCase2.id, &upd_doctor)
		if (err != nil) != testCase2.wantErr {
			t.Errorf("Update() error = %v, wantErr %v", err, testCase2.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase2.want) {
			t.Errorf("Update() got = %v, want %v", got, testCase2.want)
		}
	})
}

func TestToDoctorDTO(t *testing.T) {
	log, levelog, err := logger.NewLogger()

	if err != nil {
		t.Fatalf("failed to init logger: %v", err)
	}
	cfg, err := config.NewConfig(log, levelog)
	if err != nil {
		t.Fatalf("failed to init config: %v", err)
	}
	// Create a new in-memory database for testing
	client, err := db.NewDBClient(cfg, log)
	err = db.TruncateAll(client)
	if err != nil {
		t.Fatalf("failed to truncate all: %v", err)
	}
	// create a new doctor
	doctor, err := client.Doctor.Create().
		SetSurname("Kovel").
		SetSpeciality("Doctor").
		SetRole("Doctor").
		SetTokenId("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}
	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}

	// Create a new doctor repository
	repo := NewDoctorRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DoctorRepo
		id      int
		want    *dto.Doctor
		wantErr bool
	}{
		name: "Successful get by ID",
		repo: repo,
		id:   doctor.ID,
		want: &dto.Doctor{
			Id:         doctor.ID,
			Surname:    "Kovel",
			Speciality: "Doctor",
			Role:       "Doctor",
			TokenId:    "1",
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got := ToDoctorDTO(doctor)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("ToDoctorDTO() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("ToDoctorDTO() got = %v, want %v", got, testCase1.want)
		}
	})

}

func TestToDoctorDTOs(t *testing.T) {
	log, levelog, err := logger.NewLogger()

	if err != nil {
		t.Fatalf("failed to init logger: %v", err)
	}
	cfg, err := config.NewConfig(log, levelog)
	if err != nil {
		t.Fatalf("failed to init config: %v", err)
	}
	// Create a new in-memory database for testing
	client, err := db.NewDBClient(cfg, log)
	err = db.TruncateAll(client)
	if err != nil {
		t.Fatalf("failed to truncate all: %v", err)
	}
	// create a new doctor
	doctor, err := client.Doctor.Create().
		SetSurname("Kovel").
		SetSpeciality("Doctor").
		SetRole("Doctor").
		SetTokenId("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}

	// Create a new doctor repository
	repo := NewDoctorRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DoctorRepo
		id      int
		want    dto.Doctors
		wantErr bool
	}{
		name: "Successful To doctor dtos",
		repo: repo,
		id:   doctor.ID,
		want: dto.Doctors{
			&dto.Doctor{
				Id:         doctor.ID,
				Surname:    "Kovel",
				Speciality: "Doctor",
				Role:       "Doctor",
				TokenId:    "1",
			},
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		doctors, err := client.Doctor.Query().All(context.Background())
		got := ToDoctorDTOs(doctors)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("ToDoctorDTOs() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("ToDoctorDTOs() got = %v, want %v", got, testCase1.want)
		}
	})
}

func TestDoctorRepo_GetByTokenId(t *testing.T) {
	log, levelog, err := logger.NewLogger()

	if err != nil {
		t.Fatalf("failed to init logger: %v", err)
	}
	cfg, err := config.NewConfig(log, levelog)
	if err != nil {
		t.Fatalf("failed to init config: %v", err)
	}
	// Create a new in-memory database for testing
	client, err := db.NewDBClient(cfg, log)
	err = db.TruncateAll(client)
	if err != nil {
		t.Fatalf("failed to truncate all: %v", err)
	}
	// create a new doctor
	doctor, err := client.Doctor.Create().
		SetSurname("Kovel").
		SetSpeciality("Doctor").
		SetRole("Doctor").
		SetTokenId("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}
	if err != nil {
		t.Fatalf("failed to create doctor: %v", err)
	}

	// Create a new doctor repository
	repo := NewDoctorRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DoctorRepo
		tokenId string
		want    *dto.Doctor
		wantErr bool
	}{
		name:    "Successful get by ID",
		repo:    repo,
		tokenId: doctor.TokenId,
		want: &dto.Doctor{
			Id:         doctor.ID,
			Surname:    "Kovel",
			Speciality: "Doctor",
			Role:       "Doctor",
			TokenId:    "1",
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got, err := testCase1.repo.GetByTokenId(context.Background(), testCase1.tokenId)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("GetByTokenId() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("GetByTokenId() got = %v, want %v", got, testCase1.want)
		}
	})

	// Test case 2: Get by ID with non-existent doctor
	testCase2 := struct {
		name    string
		repo    *DoctorRepo
		tokenId string
		want    *dto.Doctor
		wantErr bool
	}{
		name:    "Get by ID with non-existent doctor",
		repo:    repo,
		tokenId: "100",
		want:    nil,
		wantErr: true,
	}

	// Run the test case 2
	runner.Run(t, testCase2.name, func(t provider.T) {
		got, err := testCase2.repo.GetByTokenId(context.Background(), testCase2.tokenId)
		if (err != nil) != testCase2.wantErr {
			t.Errorf("GetByTokenId() error = %v, wantErr %v", err, testCase2.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase2.want) {
			t.Errorf("GetByTokenId() got = %v, want %v", got, testCase2.want)
		}
	})
}
