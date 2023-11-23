package repo

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"hospital/internal/modules/config"
	"hospital/internal/modules/db"
	"hospital/internal/modules/db/ent"
	"hospital/internal/modules/domain/disease/dto"
	"hospital/internal/modules/logger"
	"reflect"
	"testing"
)

func TestNewDiseaseRepo(t *testing.T) {
	type args struct {
		client *ent.Client
	}
	client := ent.NewClient()
	// Run the test cases
	for _, tt := range []struct {
		name string
		args args
		want *DiseaseRepo
	}{
		{
			name: "NewDiseaseRepo",
			args: args{
				client: client,
			},
			want: &DiseaseRepo{
				client: client,
			},
		},
	} {
		runner.Run(t, tt.name, func(t provider.T) {
			if got := NewDiseaseRepo(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDiseaseRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiseaseRepo_Create(t *testing.T) {
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
	// create a new disease
	disease, err := client.Disease.Create().
		SetName("Kovel").
		SetThreat("Kovel").
		SetDegreeOfDanger(1).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}

	// Create a new disease repository
	repo := NewDiseaseRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DiseaseRepo
		id      int
		want    *dto.Disease
		wantErr bool
	}{
		name: "Successful get by ID",
		repo: repo,
		id:   disease.ID,
		want: &dto.Disease{
			Id:             disease.ID,
			Name:           "Kovel",
			Threat:         "Kovel",
			DegreeOfDanger: 1,
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

	// Test case 2: Get by ID with non-existent disease
	testCase2 := struct {
		name    string
		repo    *DiseaseRepo
		id      int
		want    *dto.Disease
		wantErr bool
	}{
		name:    "Get by ID with non-existent disease",
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

func TestDiseaseRepo_Delete(t *testing.T) {
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
	// create a new disease
	disease, err := client.Disease.Create().
		SetName("Kovel").
		SetThreat("Kovel").
		SetDegreeOfDanger(1).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}
	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to update disease: %v", err)
	}

	// Create a new disease repository
	repo := NewDiseaseRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DiseaseRepo
		id      int
		want    *dto.Disease
		wantErr bool
	}{
		name: "Successful Deleted",
		repo: repo,
		id:   disease.ID,
		want: &dto.Disease{
			Id:             disease.ID,
			Name:           "Kovel",
			Threat:         "Kovel",
			DegreeOfDanger: 1,
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

	// Test case 2: Get by ID with non-existent disease
	testCase2 := struct {
		name    string
		repo    *DiseaseRepo
		id      int
		want    *dto.Disease
		wantErr bool
	}{
		name:    "Delete by ID with non-existent disease",
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

func TestDiseaseRepo_GetById(t *testing.T) {
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
	// create a new disease
	disease, err := client.Disease.Create().
		SetName("Kovel").
		SetThreat("Kovel").
		SetDegreeOfDanger(1).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}
	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}

	// Create a new disease repository
	repo := NewDiseaseRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DiseaseRepo
		id      int
		want    *dto.Disease
		wantErr bool
	}{
		name: "Successful get by ID",
		repo: repo,
		id:   disease.ID,
		want: &dto.Disease{
			Id:             disease.ID,
			Name:           "Kovel",
			Threat:         "Kovel",
			DegreeOfDanger: 1,
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

	// Test case 2: Get by ID with non-existent disease
	testCase2 := struct {
		name    string
		repo    *DiseaseRepo
		id      int
		want    *dto.Disease
		wantErr bool
	}{
		name:    "Get by ID with non-existent disease",
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

func TestDiseaseRepo_List(t *testing.T) {
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
	// create a new disease
	disease, err := client.Disease.Create().
		SetName("Kovel").
		SetThreat("Kovel").
		SetDegreeOfDanger(1).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}

	// Create a new disease repository
	repo := NewDiseaseRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DiseaseRepo
		id      int
		want    dto.Diseases
		wantErr bool
	}{
		name: "Successful list",
		repo: repo,
		id:   disease.ID,
		want: dto.Diseases{
			&dto.Disease{
				Id:             disease.ID,
				Name:           "Kovel",
				Threat:         "Kovel",
				DegreeOfDanger: 1,
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

func TestDiseaseRepo_Restore(t *testing.T) {
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
	// create a new disease
	disease, err := client.Disease.Create().
		SetName("Kovel").
		SetThreat("Kovel").
		SetDegreeOfDanger(1).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to update disease: %v", err)
	}

	// Create a new disease repository
	repo := NewDiseaseRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DiseaseRepo
		id      int
		want    *dto.Disease
		wantErr bool
	}{
		name: "Successful Restored",
		repo: repo,
		id:   disease.ID,
		want: &dto.Disease{
			Id:             disease.ID,
			Name:           "Kovel",
			Threat:         "Kovel",
			DegreeOfDanger: 1,
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

	// Test case 2: Get by ID with non-existent disease
	testCase2 := struct {
		name    string
		repo    *DiseaseRepo
		id      int
		want    *dto.Disease
		wantErr bool
	}{
		name:    "Resotre by ID with non-existent disease",
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

func TestDiseaseRepo_Update(t *testing.T) {
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
	// create a new disease
	disease, err := client.Disease.Create().
		SetName("Kovel").
		SetThreat("Kovel").
		SetDegreeOfDanger(1).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}

	_, err = client.Disease.UpdateOneID(disease.ID).
		SetName("Kovel").
		SetThreat("Kovel").
		SetDegreeOfDanger(1).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to update disease: %v", err)
	}

	// Create a new disease repository
	repo := NewDiseaseRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DiseaseRepo
		id      int
		want    *dto.Disease
		wantErr bool
	}{
		name: "Successful Updated",
		repo: repo,
		id:   disease.ID,
		want: &dto.Disease{
			Id:             disease.ID,
			Name:           "Kovel",
			Threat:         "Kovel",
			DegreeOfDanger: 1,
		},
		wantErr: false,
	}
	upd_disease := dto.UpdateDisease{
		Name:           "Kovel",
		Threat:         "Kovel",
		DegreeOfDanger: 1,
	}
	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got, err := testCase1.repo.Update(context.Background(), testCase1.id, &upd_disease)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("Update() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("Update() got = %v, want %v", got, testCase1.want)
		}
	})

	// Test case 2: Get by ID with non-existent disease
	testCase2 := struct {
		name    string
		repo    *DiseaseRepo
		id      int
		want    *dto.Disease
		wantErr bool
	}{
		name:    "Get by ID with non-existent disease",
		repo:    repo,
		id:      100,
		want:    nil,
		wantErr: true,
	}

	// Run the test case 2
	runner.Run(t, testCase2.name, func(t provider.T) {
		got, err := testCase2.repo.Update(context.Background(), testCase2.id, &upd_disease)
		if (err != nil) != testCase2.wantErr {
			t.Errorf("Update() error = %v, wantErr %v", err, testCase2.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase2.want) {
			t.Errorf("Update() got = %v, want %v", got, testCase2.want)
		}
	})
}

func TestToDiseaseDTO(t *testing.T) {
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
	// create a new disease
	disease, err := client.Disease.Create().
		SetName("Kovel").
		SetThreat("Kovel").
		SetDegreeOfDanger(1).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}
	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}

	// Create a new disease repository
	repo := NewDiseaseRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DiseaseRepo
		id      int
		want    *dto.Disease
		wantErr bool
	}{
		name: "Successful get by ID",
		repo: repo,
		id:   disease.ID,
		want: &dto.Disease{
			Id:             disease.ID,
			Name:           "Kovel",
			Threat:         "Kovel",
			DegreeOfDanger: 1,
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got := ToDiseaseDTO(disease)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("ToDiseaseDTO() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("ToDiseaseDTO() got = %v, want %v", got, testCase1.want)
		}
	})

}

func TestToDiseaseDTOs(t *testing.T) {
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
	// create a new disease
	disease, err := client.Disease.Create().
		SetName("Kovel").
		SetThreat("Kovel").
		SetDegreeOfDanger(1).
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to create disease: %v", err)
	}

	// Create a new disease repository
	repo := NewDiseaseRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *DiseaseRepo
		id      int
		want    dto.Diseases
		wantErr bool
	}{
		name: "Successful To disease dtos",
		repo: repo,
		id:   disease.ID,
		want: dto.Diseases{
			&dto.Disease{
				Id:             disease.ID,
				Name:           "Kovel",
				Threat:         "Kovel",
				DegreeOfDanger: 1,
			},
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		diseases, err := client.Disease.Query().All(context.Background())
		got := ToDiseaseDTOs(diseases)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("ToDiseaseDTOs() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("ToDiseaseDTOs() got = %v, want %v", got, testCase1.want)
		}
	})
}
