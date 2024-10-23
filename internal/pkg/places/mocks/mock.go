// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_places is a generated GoMock package.
package mock_places

import (
	models "2024_2_ThereWillBeName/internal/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPlaceRepo is a mock of PlaceRepo interface.
type MockPlaceRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPlaceRepoMockRecorder
}

// MockPlaceRepoMockRecorder is the mock recorder for MockPlaceRepo.
type MockPlaceRepoMockRecorder struct {
	mock *MockPlaceRepo
}

// NewMockPlaceRepo creates a new mock instance.
func NewMockPlaceRepo(ctrl *gomock.Controller) *MockPlaceRepo {
	mock := &MockPlaceRepo{ctrl: ctrl}
	mock.recorder = &MockPlaceRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlaceRepo) EXPECT() *MockPlaceRepoMockRecorder {
	return m.recorder
}

// CreatePlace mocks base method.
func (m *MockPlaceRepo) CreatePlace(ctx context.Context, place models.CreatePlace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlace", ctx, place)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePlace indicates an expected call of CreatePlace.
func (mr *MockPlaceRepoMockRecorder) CreatePlace(ctx, place interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlace", reflect.TypeOf((*MockPlaceRepo)(nil).CreatePlace), ctx, place)
}

// DeletePlace mocks base method.
func (m *MockPlaceRepo) DeletePlace(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlace", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePlace indicates an expected call of DeletePlace.
func (mr *MockPlaceRepoMockRecorder) DeletePlace(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlace", reflect.TypeOf((*MockPlaceRepo)(nil).DeletePlace), ctx, id)
}

// GetPlace mocks base method.
func (m *MockPlaceRepo) GetPlace(ctx context.Context, id int) (models.GetPlace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlace", ctx, id)
	ret0, _ := ret[0].(models.GetPlace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlace indicates an expected call of GetPlace.
func (mr *MockPlaceRepoMockRecorder) GetPlace(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlace", reflect.TypeOf((*MockPlaceRepo)(nil).GetPlace), ctx, id)
}

// GetPlaces mocks base method.
func (m *MockPlaceRepo) GetPlaces(ctx context.Context, limit, offset int) ([]models.GetPlace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaces", ctx, limit, offset)
	ret0, _ := ret[0].([]models.GetPlace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaces indicates an expected call of GetPlaces.
func (mr *MockPlaceRepoMockRecorder) GetPlaces(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaces", reflect.TypeOf((*MockPlaceRepo)(nil).GetPlaces), ctx, limit, offset)
}

// SearchPlaces mocks base method.
func (m *MockPlaceRepo) SearchPlaces(ctx context.Context, name string, limit, offset int) ([]models.GetPlace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPlaces", ctx, name, limit, offset)
	ret0, _ := ret[0].([]models.GetPlace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPlaces indicates an expected call of SearchPlaces.
func (mr *MockPlaceRepoMockRecorder) SearchPlaces(ctx, name, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPlaces", reflect.TypeOf((*MockPlaceRepo)(nil).SearchPlaces), ctx, name, limit, offset)
}

// UpdatePlace mocks base method.
func (m *MockPlaceRepo) UpdatePlace(ctx context.Context, place models.UpdatePlace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePlace", ctx, place)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePlace indicates an expected call of UpdatePlace.
func (mr *MockPlaceRepoMockRecorder) UpdatePlace(ctx, place interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlace", reflect.TypeOf((*MockPlaceRepo)(nil).UpdatePlace), ctx, place)
}

// MockPlaceUsecase is a mock of PlaceUsecase interface.
type MockPlaceUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockPlaceUsecaseMockRecorder
}

// MockPlaceUsecaseMockRecorder is the mock recorder for MockPlaceUsecase.
type MockPlaceUsecaseMockRecorder struct {
	mock *MockPlaceUsecase
}

// NewMockPlaceUsecase creates a new mock instance.
func NewMockPlaceUsecase(ctrl *gomock.Controller) *MockPlaceUsecase {
	mock := &MockPlaceUsecase{ctrl: ctrl}
	mock.recorder = &MockPlaceUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlaceUsecase) EXPECT() *MockPlaceUsecaseMockRecorder {
	return m.recorder
}

// CreatePlace mocks base method.
func (m *MockPlaceUsecase) CreatePlace(ctx context.Context, place models.CreatePlace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePlace", ctx, place)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePlace indicates an expected call of CreatePlace.
func (mr *MockPlaceUsecaseMockRecorder) CreatePlace(ctx, place interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePlace", reflect.TypeOf((*MockPlaceUsecase)(nil).CreatePlace), ctx, place)
}

// DeletePlace mocks base method.
func (m *MockPlaceUsecase) DeletePlace(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlace", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePlace indicates an expected call of DeletePlace.
func (mr *MockPlaceUsecaseMockRecorder) DeletePlace(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlace", reflect.TypeOf((*MockPlaceUsecase)(nil).DeletePlace), ctx, id)
}

// GetPlace mocks base method.
func (m *MockPlaceUsecase) GetPlace(ctx context.Context, id int) (models.GetPlace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlace", ctx, id)
	ret0, _ := ret[0].(models.GetPlace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlace indicates an expected call of GetPlace.
func (mr *MockPlaceUsecaseMockRecorder) GetPlace(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlace", reflect.TypeOf((*MockPlaceUsecase)(nil).GetPlace), ctx, id)
}

// GetPlaces mocks base method.
func (m *MockPlaceUsecase) GetPlaces(ctx context.Context, limit, offset int) ([]models.GetPlace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaces", ctx, limit, offset)
	ret0, _ := ret[0].([]models.GetPlace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaces indicates an expected call of GetPlaces.
func (mr *MockPlaceUsecaseMockRecorder) GetPlaces(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaces", reflect.TypeOf((*MockPlaceUsecase)(nil).GetPlaces), ctx, limit, offset)
}

// SearchPlaces mocks base method.
func (m *MockPlaceUsecase) SearchPlaces(ctx context.Context, name string, limit, offset int) ([]models.GetPlace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPlaces", ctx, name, limit, offset)
	ret0, _ := ret[0].([]models.GetPlace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPlaces indicates an expected call of SearchPlaces.
func (mr *MockPlaceUsecaseMockRecorder) SearchPlaces(ctx, name, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPlaces", reflect.TypeOf((*MockPlaceUsecase)(nil).SearchPlaces), ctx, name, limit, offset)
}

// UpdatePlace mocks base method.
func (m *MockPlaceUsecase) UpdatePlace(ctx context.Context, place models.UpdatePlace) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePlace", ctx, place)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePlace indicates an expected call of UpdatePlace.
func (mr *MockPlaceUsecaseMockRecorder) UpdatePlace(ctx, place interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePlace", reflect.TypeOf((*MockPlaceUsecase)(nil).UpdatePlace), ctx, place)
}
