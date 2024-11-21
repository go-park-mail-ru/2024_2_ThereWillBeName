// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_reviews is a generated GoMock package.
package mock_reviews

import (
	models "2024_2_ThereWillBeName/internal/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReviewsUsecase is a mock of ReviewsUsecase interface.
type MockReviewsUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockReviewsUsecaseMockRecorder
}

// MockReviewsUsecaseMockRecorder is the mock recorder for MockReviewsUsecase.
type MockReviewsUsecaseMockRecorder struct {
	mock *MockReviewsUsecase
}

// NewMockReviewsUsecase creates a new mock instance.
func NewMockReviewsUsecase(ctrl *gomock.Controller) *MockReviewsUsecase {
	mock := &MockReviewsUsecase{ctrl: ctrl}
	mock.recorder = &MockReviewsUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReviewsUsecase) EXPECT() *MockReviewsUsecaseMockRecorder {
	return m.recorder
}

// CreateReview mocks base method.
func (m *MockReviewsUsecase) CreateReview(ctx context.Context, review models.Review) (models.GetReview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReview", ctx, review)
	ret0, _ := ret[0].(models.GetReview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReview indicates an expected call of CreateReview.
func (mr *MockReviewsUsecaseMockRecorder) CreateReview(ctx, review interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReview", reflect.TypeOf((*MockReviewsUsecase)(nil).CreateReview), ctx, review)
}

// DeleteReview mocks base method.
func (m *MockReviewsUsecase) DeleteReview(ctx context.Context, reviewID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReview", ctx, reviewID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReview indicates an expected call of DeleteReview.
func (mr *MockReviewsUsecaseMockRecorder) DeleteReview(ctx, reviewID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReview", reflect.TypeOf((*MockReviewsUsecase)(nil).DeleteReview), ctx, reviewID)
}

// GetReview mocks base method.
func (m *MockReviewsUsecase) GetReview(ctx context.Context, reviewID uint) (models.GetReview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReview", ctx, reviewID)
	ret0, _ := ret[0].(models.GetReview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReview indicates an expected call of GetReview.
func (mr *MockReviewsUsecaseMockRecorder) GetReview(ctx, reviewID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReview", reflect.TypeOf((*MockReviewsUsecase)(nil).GetReview), ctx, reviewID)
}

// GetReviewsByPlaceID mocks base method.
func (m *MockReviewsUsecase) GetReviewsByPlaceID(ctx context.Context, placeID uint, limit, offset int) ([]models.GetReview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviewsByPlaceID", ctx, placeID, limit, offset)
	ret0, _ := ret[0].([]models.GetReview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReviewsByPlaceID indicates an expected call of GetReviewsByPlaceID.
func (mr *MockReviewsUsecaseMockRecorder) GetReviewsByPlaceID(ctx, placeID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviewsByPlaceID", reflect.TypeOf((*MockReviewsUsecase)(nil).GetReviewsByPlaceID), ctx, placeID, limit, offset)
}

// GetReviewsByUserID mocks base method.
func (m *MockReviewsUsecase) GetReviewsByUserID(ctx context.Context, userID uint, limit, offset int) ([]models.GetReviewByUserID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviewsByUserID", ctx, userID, limit, offset)
	ret0, _ := ret[0].([]models.GetReviewByUserID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReviewsByUserID indicates an expected call of GetReviewsByUserID.
func (mr *MockReviewsUsecaseMockRecorder) GetReviewsByUserID(ctx, userID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviewsByUserID", reflect.TypeOf((*MockReviewsUsecase)(nil).GetReviewsByUserID), ctx, userID, limit, offset)
}

// UpdateReview mocks base method.
func (m *MockReviewsUsecase) UpdateReview(ctx context.Context, review models.Review) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReview", ctx, review)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateReview indicates an expected call of UpdateReview.
func (mr *MockReviewsUsecaseMockRecorder) UpdateReview(ctx, review interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReview", reflect.TypeOf((*MockReviewsUsecase)(nil).UpdateReview), ctx, review)
}

// MockReviewsRepo is a mock of ReviewsRepo interface.
type MockReviewsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockReviewsRepoMockRecorder
}

// MockReviewsRepoMockRecorder is the mock recorder for MockReviewsRepo.
type MockReviewsRepoMockRecorder struct {
	mock *MockReviewsRepo
}

// NewMockReviewsRepo creates a new mock instance.
func NewMockReviewsRepo(ctrl *gomock.Controller) *MockReviewsRepo {
	mock := &MockReviewsRepo{ctrl: ctrl}
	mock.recorder = &MockReviewsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReviewsRepo) EXPECT() *MockReviewsRepoMockRecorder {
	return m.recorder
}

// CreateReview mocks base method.
func (m *MockReviewsRepo) CreateReview(ctx context.Context, review models.Review) (models.GetReview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReview", ctx, review)
	ret0, _ := ret[0].(models.GetReview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReview indicates an expected call of CreateReview.
func (mr *MockReviewsRepoMockRecorder) CreateReview(ctx, review interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReview", reflect.TypeOf((*MockReviewsRepo)(nil).CreateReview), ctx, review)
}

// DeleteReview mocks base method.
func (m *MockReviewsRepo) DeleteReview(ctx context.Context, reviewID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReview", ctx, reviewID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReview indicates an expected call of DeleteReview.
func (mr *MockReviewsRepoMockRecorder) DeleteReview(ctx, reviewID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReview", reflect.TypeOf((*MockReviewsRepo)(nil).DeleteReview), ctx, reviewID)
}

// GetReview mocks base method.
func (m *MockReviewsRepo) GetReview(ctx context.Context, reviewID uint) (models.GetReview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReview", ctx, reviewID)
	ret0, _ := ret[0].(models.GetReview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReview indicates an expected call of GetReview.
func (mr *MockReviewsRepoMockRecorder) GetReview(ctx, reviewID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReview", reflect.TypeOf((*MockReviewsRepo)(nil).GetReview), ctx, reviewID)
}

// GetReviewsByPlaceID mocks base method.
func (m *MockReviewsRepo) GetReviewsByPlaceID(ctx context.Context, placeID uint, limit, offset int) ([]models.GetReview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviewsByPlaceID", ctx, placeID, limit, offset)
	ret0, _ := ret[0].([]models.GetReview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReviewsByPlaceID indicates an expected call of GetReviewsByPlaceID.
func (mr *MockReviewsRepoMockRecorder) GetReviewsByPlaceID(ctx, placeID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviewsByPlaceID", reflect.TypeOf((*MockReviewsRepo)(nil).GetReviewsByPlaceID), ctx, placeID, limit, offset)
}

// GetReviewsByUserID mocks base method.
func (m *MockReviewsRepo) GetReviewsByUserID(ctx context.Context, userID uint, limit, offset int) ([]models.GetReviewByUserID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviewsByUserID", ctx, userID, limit, offset)
	ret0, _ := ret[0].([]models.GetReviewByUserID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReviewsByUserID indicates an expected call of GetReviewsByUserID.
func (mr *MockReviewsRepoMockRecorder) GetReviewsByUserID(ctx, userID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviewsByUserID", reflect.TypeOf((*MockReviewsRepo)(nil).GetReviewsByUserID), ctx, userID, limit, offset)
}

// UpdateReview mocks base method.
func (m *MockReviewsRepo) UpdateReview(ctx context.Context, review models.Review) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReview", ctx, review)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateReview indicates an expected call of UpdateReview.
func (mr *MockReviewsRepoMockRecorder) UpdateReview(ctx, review interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReview", reflect.TypeOf((*MockReviewsRepo)(nil).UpdateReview), ctx, review)
}