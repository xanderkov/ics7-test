package repo

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"hospital/internal/modules/config"
	"hospital/internal/modules/db"
	"hospital/internal/modules/db/ent"
	"hospital/internal/modules/domain/patient/dto"
	"hospital/internal/modules/logger"
	"reflect"
	"testing"
)

func TestNewPatientRepo(t *testing.T) {
	type args struct {
		client *ent.Client
	}
	client := ent.NewClient()
	// Run the test cases
	for _, tt := range []struct {
		name string
		args args
		want *PatientRepo
	}{
		{
			name: "NewPatientRepo",
			args: args{
				client: client,
			},
			want: &PatientRepo{
				client: client,
			},
		},
	} {
		runner.Run(t, tt.name, func(t provider.T) {
			if got := NewPatientRepo(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPatientRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPatientRepo_Create(t *testing.T) {
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
	// create a new room
	room, err := client.Room.Create().
		SetNumberPatients(1).
		SetFloor(1).
		SetNumber(1).
		SetNumberBeds(1).
		SetTypeRoom("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}
	// Create a new patient
	patient, err := client.Patient.Create().
		SetName("John").
		SetSurname("Doe").
		SetHeight(180).
		SetDegreeOfDanger(2).
		SetWeight(80).
		SetPatronymic("Abob").
		SetRoomNumber(room.ID).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}

	// Create a new patient repository
	repo := NewPatientRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *PatientRepo
		id      int
		want    *dto.Patient
		wantErr bool
	}{
		name: "Successful get by ID",
		repo: repo,
		id:   patient.ID,
		want: &dto.Patient{
			Id:             patient.ID,
			Name:           "John",
			Surname:        "Doe",
			Patronymic:     "Abob",
			Height:         180,
			DegreeOfDanger: 2,
			Weight:         80,
			RoomNumber:     room.ID,
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

	// Test case 2: Get by ID with non-existent patient
	testCase2 := struct {
		name    string
		repo    *PatientRepo
		id      int
		want    *dto.Patient
		wantErr bool
	}{
		name:    "Get by ID with non-existent patient",
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

func TestPatientRepo_Delete(t *testing.T) {
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
	// create a new room
	room, err := client.Room.Create().
		SetNumberPatients(1).
		SetFloor(1).
		SetNumber(1).
		SetNumberBeds(1).
		SetTypeRoom("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}
	// Create a new patient
	patient, err := client.Patient.Create().
		SetName("John").
		SetSurname("Doe").
		SetHeight(180).
		SetDegreeOfDanger(2).
		SetWeight(80).
		SetPatronymic("Abob").
		SetRoomNumber(room.ID).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to update patient: %v", err)
	}

	// Create a new patient repository
	repo := NewPatientRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *PatientRepo
		id      int
		want    *dto.Patient
		wantErr bool
	}{
		name: "Successful Deleted",
		repo: repo,
		id:   patient.ID,
		want: &dto.Patient{
			Id:             patient.ID,
			Name:           "John",
			Surname:        "Doe",
			Patronymic:     "Abob",
			Height:         180,
			DegreeOfDanger: 2,
			Weight:         80,
			RoomNumber:     room.ID,
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

	// Test case 2: Get by ID with non-existent patient
	testCase2 := struct {
		name    string
		repo    *PatientRepo
		id      int
		want    *dto.Patient
		wantErr bool
	}{
		name:    "Delete by ID with non-existent patient",
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

func TestPatientRepo_GetById(t *testing.T) {
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
	// create a new room
	room, err := client.Room.Create().
		SetNumberPatients(1).
		SetFloor(1).
		SetNumber(1).
		SetNumberBeds(1).
		SetTypeRoom("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}
	// Create a new patient
	patient, err := client.Patient.Create().
		SetName("John").
		SetSurname("Doe").
		SetHeight(180).
		SetDegreeOfDanger(2).
		SetWeight(80).
		SetPatronymic("Abob").
		SetRoomNumber(room.ID).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}

	// Create a new patient repository
	repo := NewPatientRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *PatientRepo
		id      int
		want    *dto.Patient
		wantErr bool
	}{
		name: "Successful get by ID",
		repo: repo,
		id:   patient.ID,
		want: &dto.Patient{
			Id:             patient.ID,
			Name:           "John",
			Surname:        "Doe",
			Patronymic:     "Abob",
			Height:         180,
			DegreeOfDanger: 2,
			Weight:         80,
			RoomNumber:     room.ID,
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

	// Test case 2: Get by ID with non-existent patient
	testCase2 := struct {
		name    string
		repo    *PatientRepo
		id      int
		want    *dto.Patient
		wantErr bool
	}{
		name:    "Get by ID with non-existent patient",
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

func TestPatientRepo_List(t *testing.T) {
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
	// create a new room
	room, err := client.Room.Create().
		SetNumberPatients(1).
		SetFloor(1).
		SetNumber(1).
		SetNumberBeds(1).
		SetTypeRoom("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}
	// Create a new patient
	patient, err := client.Patient.Create().
		SetName("John").
		SetSurname("Doe").
		SetHeight(180).
		SetDegreeOfDanger(2).
		SetWeight(80).
		SetPatronymic("Abob").
		SetRoomNumber(room.ID).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}

	// Create a new patient repository
	repo := NewPatientRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *PatientRepo
		id      int
		want    dto.Patients
		wantErr bool
	}{
		name: "Successful list",
		repo: repo,
		id:   patient.ID,
		want: dto.Patients{
			&dto.Patient{
				Id:             patient.ID,
				Name:           "John",
				Surname:        "Doe",
				Patronymic:     "Abob",
				Height:         180,
				DegreeOfDanger: 2,
				Weight:         80,
				RoomNumber:     room.ID,
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

func TestPatientRepo_Restore(t *testing.T) {
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
	// create a new room
	room, err := client.Room.Create().
		SetNumberPatients(1).
		SetFloor(1).
		SetNumber(1).
		SetNumberBeds(1).
		SetTypeRoom("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}
	// Create a new patient
	patient, err := client.Patient.Create().
		SetName("John").
		SetSurname("Doe").
		SetHeight(180).
		SetDegreeOfDanger(2).
		SetWeight(80).
		SetPatronymic("Abob").
		SetRoomNumber(room.ID).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to update patient: %v", err)
	}

	// Create a new patient repository
	repo := NewPatientRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *PatientRepo
		id      int
		want    *dto.Patient
		wantErr bool
	}{
		name: "Successful Restored",
		repo: repo,
		id:   patient.ID,
		want: &dto.Patient{
			Id:             patient.ID,
			Name:           "John",
			Surname:        "Doe",
			Patronymic:     "Abob",
			Height:         180,
			DegreeOfDanger: 2,
			Weight:         80,
			RoomNumber:     room.ID,
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

	// Test case 2: Get by ID with non-existent patient
	testCase2 := struct {
		name    string
		repo    *PatientRepo
		id      int
		want    *dto.Patient
		wantErr bool
	}{
		name:    "Resotre by ID with non-existent patient",
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

func TestPatientRepo_Update(t *testing.T) {
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
	// create a new room
	room, err := client.Room.Create().
		SetNumberPatients(1).
		SetFloor(1).
		SetNumber(1).
		SetNumberBeds(1).
		SetTypeRoom("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}
	// Create a new patient
	patient, err := client.Patient.Create().
		SetName("John").
		SetSurname("Doe").
		SetHeight(180).
		SetDegreeOfDanger(2).
		SetWeight(80).
		SetPatronymic("Abob").
		SetRoomNumber(room.ID).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}

	_, err = client.Patient.UpdateOneID(patient.ID).
		SetName("John").
		SetSurname("Doe").
		SetHeight(180).
		SetDegreeOfDanger(3).
		SetWeight(80).
		SetPatronymic("Abob").
		SetRoomNumber(room.ID).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to update patient: %v", err)
	}

	// Create a new patient repository
	repo := NewPatientRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *PatientRepo
		id      int
		want    *dto.Patient
		wantErr bool
	}{
		name: "Successful Updated",
		repo: repo,
		id:   patient.ID,
		want: &dto.Patient{
			Id:             patient.ID,
			Name:           "John",
			Surname:        "Doe",
			Patronymic:     "Abob",
			Height:         180,
			DegreeOfDanger: 3,
			Weight:         80,
			RoomNumber:     room.ID,
		},
		wantErr: false,
	}
	upd_patient := dto.UpdatePatient{
		Name:           "John",
		Surname:        "Doe",
		Patronymic:     "Abob",
		Height:         180,
		DegreeOfDanger: 3,
		Weight:         80,
		RoomNumber:     room.ID,
	}
	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got, err := testCase1.repo.Update(context.Background(), testCase1.id, &upd_patient)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("Update() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("Update() got = %v, want %v", got, testCase1.want)
		}
	})

	// Test case 2: Get by ID with non-existent patient
	testCase2 := struct {
		name    string
		repo    *PatientRepo
		id      int
		want    *dto.Patient
		wantErr bool
	}{
		name:    "Get by ID with non-existent patient",
		repo:    repo,
		id:      100,
		want:    nil,
		wantErr: true,
	}

	// Run the test case 2
	runner.Run(t, testCase2.name, func(t provider.T) {
		got, err := testCase2.repo.Update(context.Background(), testCase2.id, &upd_patient)
		if (err != nil) != testCase2.wantErr {
			t.Errorf("Update() error = %v, wantErr %v", err, testCase2.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase2.want) {
			t.Errorf("Update() got = %v, want %v", got, testCase2.want)
		}
	})
}

func TestToPatientDTO(t *testing.T) {
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
	// create a new room
	room, err := client.Room.Create().
		SetNumberPatients(1).
		SetFloor(1).
		SetNumber(1).
		SetNumberBeds(1).
		SetTypeRoom("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}
	// Create a new patient
	patient, err := client.Patient.Create().
		SetName("John").
		SetSurname("Doe").
		SetHeight(180).
		SetDegreeOfDanger(2).
		SetWeight(80).
		SetPatronymic("Abob").
		SetRoomNumber(room.ID).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}

	// Create a new patient repository
	repo := NewPatientRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *PatientRepo
		id      int
		want    *dto.Patient
		wantErr bool
	}{
		name: "Successful get by ID",
		repo: repo,
		id:   patient.ID,
		want: &dto.Patient{
			Id:             patient.ID,
			Name:           "John",
			Surname:        "Doe",
			Patronymic:     "Abob",
			Height:         180,
			DegreeOfDanger: 2,
			Weight:         80,
			RoomNumber:     room.ID,
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got := ToPatientDTO(patient)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("GetById() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("GetById() got = %v, want %v", got, testCase1.want)
		}
	})

}

func TestToPatientDTOs(t *testing.T) {
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
	// create a new room
	room, err := client.Room.Create().
		SetNumberPatients(1).
		SetFloor(1).
		SetNumber(1).
		SetNumberBeds(1).
		SetTypeRoom("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}
	// Create a new patient
	patient, err := client.Patient.Create().
		SetName("John").
		SetSurname("Doe").
		SetHeight(180).
		SetDegreeOfDanger(2).
		SetWeight(80).
		SetPatronymic("Abob").
		SetRoomNumber(room.ID).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create patient: %v", err)
	}

	// Create a new patient repository
	repo := NewPatientRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *PatientRepo
		id      int
		want    dto.Patients
		wantErr bool
	}{
		name: "Successful To patient dtos",
		repo: repo,
		id:   patient.ID,
		want: dto.Patients{
			&dto.Patient{
				Id:             patient.ID,
				Name:           "John",
				Surname:        "Doe",
				Patronymic:     "Abob",
				Height:         180,
				DegreeOfDanger: 2,
				Weight:         80,
				RoomNumber:     room.ID,
			},
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		patients, err := client.Patient.Query().All(context.Background())
		got := ToPatientDTOs(patients)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("ToPatientDTOs() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("ToPatientDTOs() got = %v, want %v", got, testCase1.want)
		}
	})
}
