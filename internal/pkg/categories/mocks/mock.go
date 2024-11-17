// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_categories is a generated GoMock package.
package mock_categories

import (
	models "2024_2_ThereWillBeName/internal/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCategoriesUsecase is a mock of CategoriesUsecase interface.
type MockCategoriesUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCategoriesUsecaseMockRecorder
}

// MockCategoriesUsecaseMockRecorder is the mock recorder for MockCategoriesUsecase.
type MockCategoriesUsecaseMockRecorder struct {
	mock *MockCategoriesUsecase
}

// NewMockCategoriesUsecase creates a new mock instance.
func NewMockCategoriesUsecase(ctrl *gomock.Controller) *MockCategoriesUsecase {
	mock := &MockCategoriesUsecase{ctrl: ctrl}
	mock.recorder = &MockCategoriesUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoriesUsecase) EXPECT() *MockCategoriesUsecaseMockRecorder {
	return m.recorder
}

// GetCategories mocks base method.
func (m *MockCategoriesUsecase) GetCategories(ctx context.Context, limit, offset int) ([]models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories", ctx, limit, offset)
	ret0, _ := ret[0].([]models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockCategoriesUsecaseMockRecorder) GetCategories(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockCategoriesUsecase)(nil).GetCategories), ctx, limit, offset)
}

// MockCategoriesRepository is a mock of CategoriesRepository interface.
type MockCategoriesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoriesRepositoryMockRecorder
}

// MockCategoriesRepositoryMockRecorder is the mock recorder for MockCategoriesRepository.
type MockCategoriesRepositoryMockRecorder struct {
	mock *MockCategoriesRepository
}

// NewMockCategoriesRepository creates a new mock instance.
func NewMockCategoriesRepository(ctrl *gomock.Controller) *MockCategoriesRepository {
	mock := &MockCategoriesRepository{ctrl: ctrl}
	mock.recorder = &MockCategoriesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoriesRepository) EXPECT() *MockCategoriesRepositoryMockRecorder {
	return m.recorder
}

// GetCategories mocks base method.
func (m *MockCategoriesRepository) GetCategories(ctx context.Context, limit, offset int) ([]models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories", ctx, limit, offset)
	ret0, _ := ret[0].([]models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockCategoriesRepositoryMockRecorder) GetCategories(ctx, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockCategoriesRepository)(nil).GetCategories), ctx, limit, offset)
}
