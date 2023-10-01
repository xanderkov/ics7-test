// Code generated by MockGen. DO NOT EDIT.
// Source: hospital/src/internal/modules/domain/repo/service (interfaces: IRoomRepo)

// Package service is a generated GoMock package.
package service

import (
	context "context"
	dto "hospital/src/internal/modules/domain/repo/dto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRoomRepo is a mock of IRoomRepo interface.
type MockIRoomRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIRoomRepoMockRecorder
}

// MockIRoomRepoMockRecorder is the mock recorder for MockIRoomRepo.
type MockIRoomRepoMockRecorder struct {
	mock *MockIRoomRepo
}

// NewMockIRoomRepo creates a new mock instance.
func NewMockIRoomRepo(ctrl *gomock.Controller) *MockIRoomRepo {
	mock := &MockIRoomRepo{ctrl: ctrl}
	mock.recorder = &MockIRoomRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRoomRepo) EXPECT() *MockIRoomRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIRoomRepo) Create(arg0 context.Context, arg1 *dto.CreateRoom) (*dto.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*dto.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIRoomRepoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIRoomRepo)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIRoomRepo) Delete(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIRoomRepoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIRoomRepo)(nil).Delete), arg0, arg1)
}

// GetByNum mocks base method.
func (m *MockIRoomRepo) GetByNum(arg0 context.Context, arg1 int32) (*dto.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByNum", arg0, arg1)
	ret0, _ := ret[0].(*dto.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByNum indicates an expected call of GetByNum.
func (mr *MockIRoomRepoMockRecorder) GetByNum(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByNum", reflect.TypeOf((*MockIRoomRepo)(nil).GetByNum), arg0, arg1)
}

// List mocks base method.
func (m *MockIRoomRepo) List(arg0 context.Context) (dto.Rooms, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(dto.Rooms)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIRoomRepoMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIRoomRepo)(nil).List), arg0)
}

// Update mocks base method.
func (m *MockIRoomRepo) Update(arg0 context.Context, arg1 int32, arg2 *dto.UpdateRoom) (*dto.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dto.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIRoomRepoMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIRoomRepo)(nil).Update), arg0, arg1, arg2)
}
