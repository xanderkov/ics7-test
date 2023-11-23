package repo

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"hospital/internal/modules/config"
	"hospital/internal/modules/db"
	"hospital/internal/modules/db/ent"
	"hospital/internal/modules/domain/room/dto"
	"hospital/internal/modules/logger"
	"reflect"
	"testing"
)

func TestNewRoomRepo(t *testing.T) {
	type args struct {
		client *ent.Client
	}
	client := ent.NewClient()
	// Run the test cases
	for _, tt := range []struct {
		name string
		args args
		want *RoomRepo
	}{
		{
			name: "NewRoomRepo",
			args: args{
				client: client,
			},
			want: &RoomRepo{
				client: client,
			},
		},
	} {
		runner.Run(t, tt.name, func(t provider.T) {
			if got := NewRoomRepo(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoomRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoomRepo_Create(t *testing.T) {
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
		t.Fatalf("failed to create room: %v", err)
	}

	// Create a new room repository
	repo := NewRoomRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *RoomRepo
		id      int
		want    *dto.Room
		wantErr bool
	}{
		name: "Successful get by ID",
		repo: repo,
		id:   room.ID,
		want: &dto.Room{
			Id:             room.ID,
			Floor:          1,
			Num:            1,
			NumberBeds:     1,
			TypeRoom:       "1",
			NumberPatients: 1,
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got, err := testCase1.repo.GetByNum(context.Background(), testCase1.id)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("Create() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("Create() got = %v, want %v", got, testCase1.want)
		}
	})

	// Test case 2: Get by ID with non-existent room
	testCase2 := struct {
		name    string
		repo    *RoomRepo
		id      int
		want    *dto.Room
		wantErr bool
	}{
		name:    "Get by ID with non-existent room",
		repo:    repo,
		id:      100,
		want:    nil,
		wantErr: true,
	}

	// Run the test case 2
	runner.Run(t, testCase2.name, func(t provider.T) {
		got, err := testCase2.repo.GetByNum(context.Background(), testCase2.id)
		if (err != nil) != testCase2.wantErr {
			t.Errorf("GetByNum() error = %v, wantErr %v", err, testCase2.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase2.want) {
			t.Errorf("GetByNum() got = %v, want %v", got, testCase2.want)
		}
	})
}

func TestRoomRepo_Delete(t *testing.T) {
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
		t.Fatalf("failed to create room: %v", err)
	}
	if err != nil {
		t.Fatalf("failed to create room: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to update room: %v", err)
	}

	// Create a new room repository
	repo := NewRoomRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *RoomRepo
		id      int
		want    *dto.Room
		wantErr bool
	}{
		name: "Successful Deleted",
		repo: repo,
		id:   room.ID,
		want: &dto.Room{
			Id:             room.ID,
			Floor:          1,
			Num:            1,
			NumberBeds:     1,
			TypeRoom:       "1",
			NumberPatients: 1,
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

	// Test case 2: Get by ID with non-existent room
	testCase2 := struct {
		name    string
		repo    *RoomRepo
		id      int
		want    *dto.Room
		wantErr bool
	}{
		name:    "Delete by ID with non-existent room",
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

func TestRoomRepo_GetByNum(t *testing.T) {
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
		t.Fatalf("failed to create room: %v", err)
	}
	if err != nil {
		t.Fatalf("failed to create room: %v", err)
	}

	// Create a new room repository
	repo := NewRoomRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *RoomRepo
		id      int
		want    *dto.Room
		wantErr bool
	}{
		name: "Successful get by ID",
		repo: repo,
		id:   room.ID,
		want: &dto.Room{
			Id:             room.ID,
			Floor:          1,
			Num:            1,
			NumberBeds:     1,
			TypeRoom:       "1",
			NumberPatients: 1,
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got, err := testCase1.repo.GetByNum(context.Background(), testCase1.id)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("GetByNum() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("GetByNum() got = %v, want %v", got, testCase1.want)
		}
	})

	// Test case 2: Get by ID with non-existent room
	testCase2 := struct {
		name    string
		repo    *RoomRepo
		id      int
		want    *dto.Room
		wantErr bool
	}{
		name:    "Get by ID with non-existent room",
		repo:    repo,
		id:      100,
		want:    nil,
		wantErr: true,
	}

	// Run the test case 2
	runner.Run(t, testCase2.name, func(t provider.T) {
		got, err := testCase2.repo.GetByNum(context.Background(), testCase2.id)
		if (err != nil) != testCase2.wantErr {
			t.Errorf("GetByNum() error = %v, wantErr %v", err, testCase2.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase2.want) {
			t.Errorf("GetByNum() got = %v, want %v", got, testCase2.want)
		}
	})
}

func TestRoomRepo_List(t *testing.T) {
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
		t.Fatalf("failed to create room: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to create room: %v", err)
	}

	// Create a new room repository
	repo := NewRoomRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *RoomRepo
		id      int
		want    dto.Rooms
		wantErr bool
	}{
		name: "Successful list",
		repo: repo,
		id:   room.ID,
		want: dto.Rooms{
			&dto.Room{
				Id:             room.ID,
				Floor:          1,
				Num:            1,
				NumberBeds:     1,
				TypeRoom:       "1",
				NumberPatients: 1,
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

func TestRoomRepo_Restore(t *testing.T) {
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
		t.Fatalf("failed to create room: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to create room: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to update room: %v", err)
	}

	// Create a new room repository
	repo := NewRoomRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *RoomRepo
		id      int
		want    *dto.Room
		wantErr bool
	}{
		name: "Successful Restored",
		repo: repo,
		id:   room.ID,
		want: &dto.Room{
			Id:             room.ID,
			Floor:          1,
			Num:            1,
			NumberBeds:     1,
			TypeRoom:       "1",
			NumberPatients: 1,
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

	// Test case 2: Get by ID with non-existent room
	testCase2 := struct {
		name    string
		repo    *RoomRepo
		id      int
		want    *dto.Room
		wantErr bool
	}{
		name:    "Resotre by ID with non-existent room",
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

func TestRoomRepo_Update(t *testing.T) {
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
		t.Fatalf("failed to create room: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to create room: %v", err)
	}

	_, err = client.Room.UpdateOneID(room.ID).
		SetFloor(1).
		SetNumber(1).
		SetNumberBeds(1).
		SetTypeRoom("1").
		Save(context.Background())
	if err != nil {
		t.Fatalf("failed to update room: %v", err)
	}

	// Create a new room repository
	repo := NewRoomRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *RoomRepo
		id      int
		want    *dto.Room
		wantErr bool
	}{
		name: "Successful Updated",
		repo: repo,
		id:   room.ID,
		want: &dto.Room{
			Id:             room.ID,
			Floor:          1,
			Num:            1,
			NumberBeds:     1,
			TypeRoom:       "1",
			NumberPatients: 1,
		},
		wantErr: false,
	}
	upd_room := dto.UpdateRoom{
		Floor:          1,
		Num:            1,
		NumberBeds:     1,
		TypeRoom:       "1",
		NumberPatients: 1,
	}
	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got, err := testCase1.repo.Update(context.Background(), testCase1.id, &upd_room)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("Update() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("Update() got = %v, want %v", got, testCase1.want)
		}
	})

	// Test case 2: Get by ID with non-existent room
	testCase2 := struct {
		name    string
		repo    *RoomRepo
		id      int
		want    *dto.Room
		wantErr bool
	}{
		name:    "Get by ID with non-existent room",
		repo:    repo,
		id:      100,
		want:    nil,
		wantErr: true,
	}

	// Run the test case 2
	runner.Run(t, testCase2.name, func(t provider.T) {
		got, err := testCase2.repo.Update(context.Background(), testCase2.id, &upd_room)
		if (err != nil) != testCase2.wantErr {
			t.Errorf("Update() error = %v, wantErr %v", err, testCase2.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase2.want) {
			t.Errorf("Update() got = %v, want %v", got, testCase2.want)
		}
	})
}

func TestToRoomDTO(t *testing.T) {
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
		t.Fatalf("failed to create room: %v", err)
	}
	if err != nil {
		t.Fatalf("failed to create room: %v", err)
	}

	// Create a new room repository
	repo := NewRoomRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *RoomRepo
		id      int
		want    *dto.Room
		wantErr bool
	}{
		name: "Successful get by ID",
		repo: repo,
		id:   room.ID,
		want: &dto.Room{
			Id:             room.ID,
			Floor:          1,
			Num:            1,
			NumberBeds:     1,
			TypeRoom:       "1",
			NumberPatients: 1,
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		got := ToRoomDTO(room)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("ToRoomDTO() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("ToRoomDTO() got = %v, want %v", got, testCase1.want)
		}
	})

}

func TestToRoomDTOs(t *testing.T) {
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
		t.Fatalf("failed to create room: %v", err)
	}

	if err != nil {
		t.Fatalf("failed to create room: %v", err)
	}

	// Create a new room repository
	repo := NewRoomRepo(client)

	// Test case 1: Successful get by ID
	testCase1 := struct {
		name    string
		repo    *RoomRepo
		id      int
		want    dto.Rooms
		wantErr bool
	}{
		name: "Successful To room dtos",
		repo: repo,
		id:   room.ID,
		want: dto.Rooms{
			&dto.Room{
				Id:             room.ID,
				Floor:          1,
				Num:            1,
				NumberBeds:     1,
				TypeRoom:       "1",
				NumberPatients: 1,
			},
		},
		wantErr: false,
	}

	// Run the test case 1
	runner.Run(t, testCase1.name, func(t provider.T) {
		rooms, err := client.Room.Query().All(context.Background())
		got := ToRoomDTOs(rooms)
		if (err != nil) != testCase1.wantErr {
			t.Errorf("ToRoomDTOs() error = %v, wantErr %v", err, testCase1.wantErr)
			return
		}
		if !reflect.DeepEqual(got, testCase1.want) {
			t.Errorf("ToRoomDTOs() got = %v, want %v", got, testCase1.want)
		}
	})
}
