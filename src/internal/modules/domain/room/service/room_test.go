package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"hospital/internal/modules/domain/room/dto"
	"reflect"
	"testing"
)

func TestNewRoomService(t *testing.T) {
	type args struct {
		repo IRoomRepo
	}
	mockRoom := new(MockIRoomRepo)

	tests := []struct {
		name string
		args args
		want *RoomService
	}{
		{
			name: "Simple positive test",
			args: args{
				repo: mockRoom,
			},
			want: &RoomService{
				repo: mockRoom,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRoomService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoomService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoomService_Create(t *testing.T) {
	type fields struct {
		repo IRoomRepo
	}

	type args struct {
		ctx context.Context
		dtm *dto.CreateRoom
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIRoomRepo(ctrl)

	// Test case 1: Successful create
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Room
		wantErr bool
	}{
		name: "Successful create",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			dtm: &dto.CreateRoom{
				Num:            1,
				Floor:          1,
				NumberPatients: 1,
				NumberBeds:     1,
				TypeRoom:       "1",
			},
		},
		want: &dto.Room{
			Num:            1,
			Floor:          1,
			NumberPatients: 1,
			NumberBeds:     1,
			TypeRoom:       "1",
		},
		wantErr: false,
	}

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&dto.Room{
		Num:            1,
		Floor:          1,
		NumberPatients: 1,
		NumberBeds:     1,
		TypeRoom:       "1",
	}, nil)

	// Test case 2: Error while creating room
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Room
		wantErr bool
	}{
		name: "Error while creating room",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			dtm: &dto.CreateRoom{
				Num:            1,
				Floor:          1,
				NumberPatients: 1,
				NumberBeds:     1,
				TypeRoom:       "1",
			},
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, errors.New("error while creating room"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Room
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoomService{
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

func TestRoomService_Delete(t *testing.T) {
	type fields struct {
		repo IRoomRepo
	}
	type args struct {
		ctx context.Context
		id  int
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIRoomRepo(ctrl)

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

	// Test case 2: Error while deleting room
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		name: "Error while deleting room",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  2,
		},
		wantErr: true,
	}

	mockRepo.EXPECT().Delete(gomock.Any(), testCase2.args.id).Return(errors.New("error while deleting room"))

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
			r := &RoomService{
				repo: tt.fields.repo,
			}
			if err := r.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRoomService_GetById(t *testing.T) {
	type fields struct {
		repo IRoomRepo
	}
	type args struct {
		ctx context.Context
		id  int
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIRoomRepo(ctrl)

	// Test case 1: Successful get
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Room
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
		want: &dto.Room{
			Id:             1,
			Num:            1,
			Floor:          1,
			NumberPatients: 1,
			NumberBeds:     1,
			TypeRoom:       "1",
		},
		wantErr: false,
	}

	mockRepo.EXPECT().GetByNum(gomock.Any(), testCase1.args.id).Return(testCase1.want, nil)

	// Test case 2: Error while getting room
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Room
		wantErr bool
	}{
		name: "Error while getting room",
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

	mockRepo.EXPECT().GetByNum(gomock.Any(), testCase2.args.id).Return(nil, errors.New("error while getting room"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		want    *dto.Room
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoomService{
				repo: tt.fields.repo,
			}
			got, err := r.GetByNum(tt.args.ctx, tt.args.id)
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

func TestRoomService_List(t *testing.T) {
	type fields struct {
		repo IRoomRepo
	}
	type args struct {
		ctx context.Context
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIRoomRepo(ctrl)

	// Test case 1: Successful list
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    dto.Rooms
		wantErr bool
	}{
		name: "Successful list",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
		},
		want: dto.Rooms{
			{
				Id:             1,
				Num:            1,
				Floor:          1,
				NumberPatients: 1,
				NumberBeds:     1,
				TypeRoom:       "1",
			},
			{
				Id:             2,
				Num:            1,
				Floor:          1,
				NumberPatients: 1,
				NumberBeds:     1,
				TypeRoom:       "1",
			},
		},
		wantErr: false,
	}

	mockRepo.EXPECT().List(gomock.Any()).Return(testCase1.want, nil)

	// Test case 2: Error while listing rooms
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    dto.Rooms
		wantErr bool
	}{
		name: "Error while listing rooms",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().List(gomock.Any()).Return(nil, errors.New("error while listing rooms"))

	// Run the test cases
	for _, tt := range []struct {
		name    string
		fields  fields
		args    args
		want    dto.Rooms
		wantErr bool
	}{
		testCase1,
		testCase2,
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoomService{
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

func TestRoomService_Update(t *testing.T) {
	type fields struct {
		repo IRoomRepo
	}
	type args struct {
		ctx context.Context
		id  int
		dtm *dto.UpdateRoom
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockIRoomRepo(ctrl)

	// Test case 1: Successful update
	testCase1 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Room
		wantErr bool
	}{
		name: "Successful update",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  1,
			dtm: &dto.UpdateRoom{
				Num:            1,
				Floor:          1,
				NumberPatients: 1,
				NumberBeds:     1,
				TypeRoom:       "1",
			},
		},
		want: &dto.Room{
			Id:             1,
			Num:            1,
			Floor:          1,
			NumberPatients: 1,
			NumberBeds:     1,
			TypeRoom:       "1",
		},
		wantErr: false,
	}

	mockRepo.EXPECT().Update(gomock.Any(), testCase1.args.id, testCase1.args.dtm).Return(testCase1.want, nil)

	// Test case 2: Room not found
	testCase2 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Room
		wantErr bool
	}{
		name: "Room not found",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  1,
			dtm: &dto.UpdateRoom{
				Num:            1,
				Floor:          1,
				NumberPatients: 1,
				NumberBeds:     1,
				TypeRoom:       "1",
			},
		},
		want:    nil,
		wantErr: true,
	}

	mockRepo.EXPECT().Update(gomock.Any(), testCase2.args.id, testCase2.args.dtm).Return(nil, errors.New("room not found"))

	// Test case 3: Invalid update data
	testCase3 := struct {
		name    string
		fields  fields
		args    args
		want    *dto.Room
		wantErr bool
	}{
		name: "Invalid update data",
		fields: fields{
			repo: mockRepo,
		},
		args: args{
			ctx: context.Background(),
			id:  1,
			dtm: &dto.UpdateRoom{
				Num:            1,
				Floor:          1,
				NumberPatients: 1,
				NumberBeds:     1,
				TypeRoom:       "1",
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
		want    *dto.Room
		wantErr bool
	}{
		testCase1,
		testCase2,
		testCase3,
	} {
		t.Run(tt.name, func(t *testing.T) {
			r := &RoomService{
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
