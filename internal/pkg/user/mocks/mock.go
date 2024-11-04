// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	models "2024_2_ThereWillBeName/internal/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthUsecase is a mock of AuthUsecase interface.
type MockAuthUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockAuthUsecaseMockRecorder
}

// MockAuthUsecaseMockRecorder is the mock recorder for MockAuthUsecase.
type MockAuthUsecaseMockRecorder struct {
	mock *MockAuthUsecase
}

// NewMockAuthUsecase creates a new mock instance.
func NewMockAuthUsecase(ctrl *gomock.Controller) *MockAuthUsecase {
	mock := &MockAuthUsecase{ctrl: ctrl}
	mock.recorder = &MockAuthUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthUsecase) EXPECT() *MockAuthUsecaseMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthUsecase) Login(ctx context.Context, login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthUsecaseMockRecorder) Login(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthUsecase)(nil).Login), ctx, login, password)
}

// SignUp mocks base method.
func (m *MockAuthUsecase) SignUp(ctx context.Context, user models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp.
func (mr *MockAuthUsecaseMockRecorder) SignUp(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockAuthUsecase)(nil).SignUp), ctx, user)
}

// MockAuthRepo is a mock of AuthRepo interface.
type MockAuthRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAuthRepoMockRecorder
}

// MockAuthRepoMockRecorder is the mock recorder for MockAuthRepo.
type MockAuthRepoMockRecorder struct {
	mock *MockAuthRepo
}

// NewMockAuthRepo creates a new mock instance.
func NewMockAuthRepo(ctrl *gomock.Controller) *MockAuthRepo {
	mock := &MockAuthRepo{ctrl: ctrl}
	mock.recorder = &MockAuthRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthRepo) EXPECT() *MockAuthRepoMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthRepo) CreateUser(ctx context.Context, user models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthRepoMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthRepo)(nil).CreateUser), ctx, user)
}

// DeleteUser mocks base method.
func (m *MockAuthRepo) DeleteUser(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockAuthRepoMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockAuthRepo)(nil).DeleteUser), ctx, id)
}

// GetUserByLogin mocks base method.
func (m *MockAuthRepo) GetUserByLogin(ctx context.Context, login string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByLogin", ctx, login)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByLogin indicates an expected call of GetUserByLogin.
func (mr *MockAuthRepoMockRecorder) GetUserByLogin(ctx, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByLogin", reflect.TypeOf((*MockAuthRepo)(nil).GetUserByLogin), ctx, login)
}

// GetUsers mocks base method.
func (m *MockAuthRepo) GetUsers(ctx context.Context, count, offset int64) ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", ctx, count, offset)
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockAuthRepoMockRecorder) GetUsers(ctx, count, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockAuthRepo)(nil).GetUsers), ctx, count, offset)
}

// UpdateUser mocks base method.
func (m *MockAuthRepo) UpdateUser(ctx context.Context, user models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockAuthRepoMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockAuthRepo)(nil).UpdateUser), ctx, user)
}