// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SabariGanesh-K/cars69-service.git/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/SabariGanesh-K/cars69-service.git/db/sqlc"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddImagesToCar mocks base method.
func (m *MockStore) AddImagesToCar(arg0 context.Context, arg1 db.AddImagesToCarParams) (db.Cars, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddImagesToCar", arg0, arg1)
	ret0, _ := ret[0].(db.Cars)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddImagesToCar indicates an expected call of AddImagesToCar.
func (mr *MockStoreMockRecorder) AddImagesToCar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddImagesToCar", reflect.TypeOf((*MockStore)(nil).AddImagesToCar), arg0, arg1)
}

// CreateCar mocks base method.
func (m *MockStore) CreateCar(arg0 context.Context, arg1 db.CreateCarParams) (db.Cars, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCar", arg0, arg1)
	ret0, _ := ret[0].(db.Cars)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCar indicates an expected call of CreateCar.
func (mr *MockStoreMockRecorder) CreateCar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCar", reflect.TypeOf((*MockStore)(nil).CreateCar), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Sessions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Sessions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateVerifyEmail mocks base method.
func (m *MockStore) CreateVerifyEmail(arg0 context.Context, arg1 db.CreateVerifyEmailParams) (db.VerifyEmails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVerifyEmail", arg0, arg1)
	ret0, _ := ret[0].(db.VerifyEmails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVerifyEmail indicates an expected call of CreateVerifyEmail.
func (mr *MockStoreMockRecorder) CreateVerifyEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVerifyEmail", reflect.TypeOf((*MockStore)(nil).CreateVerifyEmail), arg0, arg1)
}

// DeleteCar mocks base method.
func (m *MockStore) DeleteCar(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCar", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCar indicates an expected call of DeleteCar.
func (mr *MockStoreMockRecorder) DeleteCar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCar", reflect.TypeOf((*MockStore)(nil).DeleteCar), arg0, arg1)
}

// GetCarById mocks base method.
func (m *MockStore) GetCarById(arg0 context.Context, arg1 int32) (db.Cars, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCarById", arg0, arg1)
	ret0, _ := ret[0].(db.Cars)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCarById indicates an expected call of GetCarById.
func (mr *MockStoreMockRecorder) GetCarById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCarById", reflect.TypeOf((*MockStore)(nil).GetCarById), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Sessions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Sessions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserOwnedCars mocks base method.
func (m *MockStore) GetUserOwnedCars(arg0 context.Context, arg1 int32) ([]db.Cars, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserOwnedCars", arg0, arg1)
	ret0, _ := ret[0].([]db.Cars)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserOwnedCars indicates an expected call of GetUserOwnedCars.
func (mr *MockStoreMockRecorder) GetUserOwnedCars(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserOwnedCars", reflect.TypeOf((*MockStore)(nil).GetUserOwnedCars), arg0, arg1)
}

// SearchCarsFTS mocks base method.
func (m *MockStore) SearchCarsFTS(arg0 context.Context, arg1 db.SearchCarsFTSParams) ([]db.Cars, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchCarsFTS", arg0, arg1)
	ret0, _ := ret[0].([]db.Cars)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCarsFTS indicates an expected call of SearchCarsFTS.
func (mr *MockStoreMockRecorder) SearchCarsFTS(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCarsFTS", reflect.TypeOf((*MockStore)(nil).SearchCarsFTS), arg0, arg1)
}

// UpdateCar mocks base method.
func (m *MockStore) UpdateCar(arg0 context.Context, arg1 db.UpdateCarParams) (db.Cars, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCar", arg0, arg1)
	ret0, _ := ret[0].(db.Cars)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCar indicates an expected call of UpdateCar.
func (mr *MockStoreMockRecorder) UpdateCar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCar", reflect.TypeOf((*MockStore)(nil).UpdateCar), arg0, arg1)
}

// UpdateVerifyEmail mocks base method.
func (m *MockStore) UpdateVerifyEmail(arg0 context.Context, arg1 db.UpdateVerifyEmailParams) (db.VerifyEmails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVerifyEmail", arg0, arg1)
	ret0, _ := ret[0].(db.VerifyEmails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateVerifyEmail indicates an expected call of UpdateVerifyEmail.
func (mr *MockStoreMockRecorder) UpdateVerifyEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVerifyEmail", reflect.TypeOf((*MockStore)(nil).UpdateVerifyEmail), arg0, arg1)
}