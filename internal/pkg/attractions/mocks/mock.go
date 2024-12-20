// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/attractions/interfaces.go

// Package mocks is a generated GoMock package.
package mocks

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

// GetPlace mocks base method.
func (m *MockPlaceRepo) GetPlace(ctx context.Context, id uint) (models.GetPlace, error) {
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

// GetPlacesByCategory mocks base method.
func (m *MockPlaceRepo) GetPlacesByCategory(ctx context.Context, category string, limit, offset int) ([]models.GetPlace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlacesByCategory", ctx, category, limit, offset)
	ret0, _ := ret[0].([]models.GetPlace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlacesByCategory indicates an expected call of GetPlacesByCategory.
func (mr *MockPlaceRepoMockRecorder) GetPlacesByCategory(ctx, category, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlacesByCategory", reflect.TypeOf((*MockPlaceRepo)(nil).GetPlacesByCategory), ctx, category, limit, offset)
}

// SearchPlaces mocks base method.
func (m *MockPlaceRepo) SearchPlaces(ctx context.Context, name string, category, city, filterType, limit, offset int) ([]models.GetPlace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPlaces", ctx, name, category, city, filterType, limit, offset)
	ret0, _ := ret[0].([]models.GetPlace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPlaces indicates an expected call of SearchPlaces.
func (mr *MockPlaceRepoMockRecorder) SearchPlaces(ctx, name, category, city, filterType, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPlaces", reflect.TypeOf((*MockPlaceRepo)(nil).SearchPlaces), ctx, name, category, city, filterType, limit, offset)
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

// GetPlace mocks base method.
func (m *MockPlaceUsecase) GetPlace(ctx context.Context, id uint) (models.GetPlace, error) {
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

// GetPlacesByCategory mocks base method.
func (m *MockPlaceUsecase) GetPlacesByCategory(ctx context.Context, category string, limit, offset int) ([]models.GetPlace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlacesByCategory", ctx, category, limit, offset)
	ret0, _ := ret[0].([]models.GetPlace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlacesByCategory indicates an expected call of GetPlacesByCategory.
func (mr *MockPlaceUsecaseMockRecorder) GetPlacesByCategory(ctx, category, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlacesByCategory", reflect.TypeOf((*MockPlaceUsecase)(nil).GetPlacesByCategory), ctx, category, limit, offset)
}

// SearchPlaces mocks base method.
func (m *MockPlaceUsecase) SearchPlaces(ctx context.Context, name string, category, city, filterType, limit, offset int) ([]models.GetPlace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPlaces", ctx, name, category, city, filterType, limit, offset)
	ret0, _ := ret[0].([]models.GetPlace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPlaces indicates an expected call of SearchPlaces.
func (mr *MockPlaceUsecaseMockRecorder) SearchPlaces(ctx, name, category, city, filterType, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPlaces", reflect.TypeOf((*MockPlaceUsecase)(nil).SearchPlaces), ctx, name, category, city, filterType, limit, offset)
}
