// Code generated by MockGen. DO NOT EDIT.
// Source: hospital/internal/modules/domain/doctor/service (interfaces: IDoctorRepo)

// Package service is a generated GoMock package.
package service

import (
	context "context"
	dto "hospital/internal/modules/domain/doctor/dto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIDoctorRepo is a mock of IDoctorRepo interface.
type MockIDoctorRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIDoctorRepoMockRecorder
}

// MockIDoctorRepoMockRecorder is the mock recorder for MockIDoctorRepo.
type MockIDoctorRepoMockRecorder struct {
	mock *MockIDoctorRepo
}

// NewMockIDoctorRepo creates a new mock instance.
func NewMockIDoctorRepo(ctrl *gomock.Controller) *MockIDoctorRepo {
	mock := &MockIDoctorRepo{ctrl: ctrl}
	mock.recorder = &MockIDoctorRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDoctorRepo) EXPECT() *MockIDoctorRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIDoctorRepo) Create(arg0 context.Context, arg1 *dto.CreateDoctor) (*dto.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*dto.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIDoctorRepoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIDoctorRepo)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIDoctorRepo) Delete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIDoctorRepoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIDoctorRepo)(nil).Delete), arg0, arg1)
}

// GetById mocks base method.
func (m *MockIDoctorRepo) GetById(arg0 context.Context, arg1 int) (*dto.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(*dto.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIDoctorRepoMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIDoctorRepo)(nil).GetById), arg0, arg1)
}

// List mocks base method.
func (m *MockIDoctorRepo) List(arg0 context.Context) (dto.Doctors, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(dto.Doctors)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIDoctorRepoMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIDoctorRepo)(nil).List), arg0)
}

// Update mocks base method.
func (m *MockIDoctorRepo) Update(arg0 context.Context, arg1 int, arg2 *dto.UpdateDoctor) (*dto.Doctor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*dto.Doctor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockIDoctorRepoMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIDoctorRepo)(nil).Update), arg0, arg1, arg2)
}
