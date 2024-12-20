// Code generated by MockGen. DO NOT EDIT.
// Source: internal/pkg/trips/interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "2024_2_ThereWillBeName/internal/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTripsUsecase is a mock of TripsUsecase interface.
type MockTripsUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockTripsUsecaseMockRecorder
}

// MockTripsUsecaseMockRecorder is the mock recorder for MockTripsUsecase.
type MockTripsUsecaseMockRecorder struct {
	mock *MockTripsUsecase
}

// NewMockTripsUsecase creates a new mock instance.
func NewMockTripsUsecase(ctrl *gomock.Controller) *MockTripsUsecase {
	mock := &MockTripsUsecase{ctrl: ctrl}
	mock.recorder = &MockTripsUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTripsUsecase) EXPECT() *MockTripsUsecaseMockRecorder {
	return m.recorder
}

// AddPhotosToTrip mocks base method.
func (m *MockTripsUsecase) AddPhotosToTrip(ctx context.Context, tripID uint, photos []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPhotosToTrip", ctx, tripID, photos)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPhotosToTrip indicates an expected call of AddPhotosToTrip.
func (mr *MockTripsUsecaseMockRecorder) AddPhotosToTrip(ctx, tripID, photos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPhotosToTrip", reflect.TypeOf((*MockTripsUsecase)(nil).AddPhotosToTrip), ctx, tripID, photos)
}

// AddPlaceToTrip mocks base method.
func (m *MockTripsUsecase) AddPlaceToTrip(ctx context.Context, tripID, placeID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPlaceToTrip", ctx, tripID, placeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPlaceToTrip indicates an expected call of AddPlaceToTrip.
func (mr *MockTripsUsecaseMockRecorder) AddPlaceToTrip(ctx, tripID, placeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPlaceToTrip", reflect.TypeOf((*MockTripsUsecase)(nil).AddPlaceToTrip), ctx, tripID, placeID)
}

// AddUserToTrip mocks base method.
func (m *MockTripsUsecase) AddUserToTrip(ctx context.Context, tripId, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserToTrip", ctx, tripId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserToTrip indicates an expected call of AddUserToTrip.
func (mr *MockTripsUsecaseMockRecorder) AddUserToTrip(ctx, tripId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserToTrip", reflect.TypeOf((*MockTripsUsecase)(nil).AddUserToTrip), ctx, tripId, userId)
}

// CreateSharingLink mocks base method.
func (m *MockTripsUsecase) CreateSharingLink(ctx context.Context, tripID uint, token, sharingOption string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSharingLink", ctx, tripID, token, sharingOption)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSharingLink indicates an expected call of CreateSharingLink.
func (mr *MockTripsUsecaseMockRecorder) CreateSharingLink(ctx, tripID, token, sharingOption interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSharingLink", reflect.TypeOf((*MockTripsUsecase)(nil).CreateSharingLink), ctx, tripID, token, sharingOption)
}

// CreateTrip mocks base method.
func (m *MockTripsUsecase) CreateTrip(ctx context.Context, trip models.Trip) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrip", ctx, trip)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTrip indicates an expected call of CreateTrip.
func (mr *MockTripsUsecaseMockRecorder) CreateTrip(ctx, trip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrip", reflect.TypeOf((*MockTripsUsecase)(nil).CreateTrip), ctx, trip)
}

// DeletePhotoFromTrip mocks base method.
func (m *MockTripsUsecase) DeletePhotoFromTrip(ctx context.Context, tripID uint, photoPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePhotoFromTrip", ctx, tripID, photoPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePhotoFromTrip indicates an expected call of DeletePhotoFromTrip.
func (mr *MockTripsUsecaseMockRecorder) DeletePhotoFromTrip(ctx, tripID, photoPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhotoFromTrip", reflect.TypeOf((*MockTripsUsecase)(nil).DeletePhotoFromTrip), ctx, tripID, photoPath)
}

// DeleteTrip mocks base method.
func (m *MockTripsUsecase) DeleteTrip(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrip", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrip indicates an expected call of DeleteTrip.
func (mr *MockTripsUsecaseMockRecorder) DeleteTrip(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrip", reflect.TypeOf((*MockTripsUsecase)(nil).DeleteTrip), ctx, id)
}

// GetSharingOption mocks base method.
func (m *MockTripsUsecase) GetSharingOption(ctx context.Context, userId, tripId uint) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSharingOption", ctx, userId, tripId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSharingOption indicates an expected call of GetSharingOption.
func (mr *MockTripsUsecaseMockRecorder) GetSharingOption(ctx, userId, tripId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSharingOption", reflect.TypeOf((*MockTripsUsecase)(nil).GetSharingOption), ctx, userId, tripId)
}

// GetSharingToken mocks base method.
func (m *MockTripsUsecase) GetSharingToken(ctx context.Context, tripID uint) (models.SharingToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSharingToken", ctx, tripID)
	ret0, _ := ret[0].(models.SharingToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSharingToken indicates an expected call of GetSharingToken.
func (mr *MockTripsUsecaseMockRecorder) GetSharingToken(ctx, tripID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSharingToken", reflect.TypeOf((*MockTripsUsecase)(nil).GetSharingToken), ctx, tripID)
}

// GetTrip mocks base method.
func (m *MockTripsUsecase) GetTrip(ctx context.Context, tripID uint) (models.Trip, []models.UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrip", ctx, tripID)
	ret0, _ := ret[0].(models.Trip)
	ret1, _ := ret[1].([]models.UserProfile)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTrip indicates an expected call of GetTrip.
func (mr *MockTripsUsecaseMockRecorder) GetTrip(ctx, tripID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrip", reflect.TypeOf((*MockTripsUsecase)(nil).GetTrip), ctx, tripID)
}

// GetTripBySharingToken mocks base method.
func (m *MockTripsUsecase) GetTripBySharingToken(ctx context.Context, troken string) (models.Trip, []models.UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTripBySharingToken", ctx, troken)
	ret0, _ := ret[0].(models.Trip)
	ret1, _ := ret[1].([]models.UserProfile)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTripBySharingToken indicates an expected call of GetTripBySharingToken.
func (mr *MockTripsUsecaseMockRecorder) GetTripBySharingToken(ctx, troken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTripBySharingToken", reflect.TypeOf((*MockTripsUsecase)(nil).GetTripBySharingToken), ctx, troken)
}

// GetTripsByUserID mocks base method.
func (m *MockTripsUsecase) GetTripsByUserID(ctx context.Context, userID uint, limit, offset int) ([]models.Trip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTripsByUserID", ctx, userID, limit, offset)
	ret0, _ := ret[0].([]models.Trip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTripsByUserID indicates an expected call of GetTripsByUserID.
func (mr *MockTripsUsecaseMockRecorder) GetTripsByUserID(ctx, userID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTripsByUserID", reflect.TypeOf((*MockTripsUsecase)(nil).GetTripsByUserID), ctx, userID, limit, offset)
}

// UpdateTrip mocks base method.
func (m *MockTripsUsecase) UpdateTrip(ctx context.Context, user models.Trip) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrip", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTrip indicates an expected call of UpdateTrip.
func (mr *MockTripsUsecaseMockRecorder) UpdateTrip(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrip", reflect.TypeOf((*MockTripsUsecase)(nil).UpdateTrip), ctx, user)
}

// MockTripsRepo is a mock of TripsRepo interface.
type MockTripsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTripsRepoMockRecorder
}

// MockTripsRepoMockRecorder is the mock recorder for MockTripsRepo.
type MockTripsRepoMockRecorder struct {
	mock *MockTripsRepo
}

// NewMockTripsRepo creates a new mock instance.
func NewMockTripsRepo(ctrl *gomock.Controller) *MockTripsRepo {
	mock := &MockTripsRepo{ctrl: ctrl}
	mock.recorder = &MockTripsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTripsRepo) EXPECT() *MockTripsRepoMockRecorder {
	return m.recorder
}

// AddPhotoToTrip mocks base method.
func (m *MockTripsRepo) AddPhotoToTrip(ctx context.Context, tripID uint, photoPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPhotoToTrip", ctx, tripID, photoPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPhotoToTrip indicates an expected call of AddPhotoToTrip.
func (mr *MockTripsRepoMockRecorder) AddPhotoToTrip(ctx, tripID, photoPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPhotoToTrip", reflect.TypeOf((*MockTripsRepo)(nil).AddPhotoToTrip), ctx, tripID, photoPath)
}

// AddPlaceToTrip mocks base method.
func (m *MockTripsRepo) AddPlaceToTrip(ctx context.Context, tripID, placeID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPlaceToTrip", ctx, tripID, placeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPlaceToTrip indicates an expected call of AddPlaceToTrip.
func (mr *MockTripsRepoMockRecorder) AddPlaceToTrip(ctx, tripID, placeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPlaceToTrip", reflect.TypeOf((*MockTripsRepo)(nil).AddPlaceToTrip), ctx, tripID, placeID)
}

// AddUserToTrip mocks base method.
func (m *MockTripsRepo) AddUserToTrip(ctx context.Context, tripId, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserToTrip", ctx, tripId, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserToTrip indicates an expected call of AddUserToTrip.
func (mr *MockTripsRepoMockRecorder) AddUserToTrip(ctx, tripId, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserToTrip", reflect.TypeOf((*MockTripsRepo)(nil).AddUserToTrip), ctx, tripId, userId)
}

// CreateSharingLink mocks base method.
func (m *MockTripsRepo) CreateSharingLink(ctx context.Context, tripID uint, token, sharingOption string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSharingLink", ctx, tripID, token, sharingOption)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSharingLink indicates an expected call of CreateSharingLink.
func (mr *MockTripsRepoMockRecorder) CreateSharingLink(ctx, tripID, token, sharingOption interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSharingLink", reflect.TypeOf((*MockTripsRepo)(nil).CreateSharingLink), ctx, tripID, token, sharingOption)
}

// CreateTrip mocks base method.
func (m *MockTripsRepo) CreateTrip(ctx context.Context, user models.Trip) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTrip", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTrip indicates an expected call of CreateTrip.
func (mr *MockTripsRepoMockRecorder) CreateTrip(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrip", reflect.TypeOf((*MockTripsRepo)(nil).CreateTrip), ctx, user)
}

// DeletePhotoFromTrip mocks base method.
func (m *MockTripsRepo) DeletePhotoFromTrip(ctx context.Context, tripID uint, photoPath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePhotoFromTrip", ctx, tripID, photoPath)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePhotoFromTrip indicates an expected call of DeletePhotoFromTrip.
func (mr *MockTripsRepoMockRecorder) DeletePhotoFromTrip(ctx, tripID, photoPath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePhotoFromTrip", reflect.TypeOf((*MockTripsRepo)(nil).DeletePhotoFromTrip), ctx, tripID, photoPath)
}

// DeleteTrip mocks base method.
func (m *MockTripsRepo) DeleteTrip(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrip", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrip indicates an expected call of DeleteTrip.
func (mr *MockTripsRepoMockRecorder) DeleteTrip(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrip", reflect.TypeOf((*MockTripsRepo)(nil).DeleteTrip), ctx, id)
}

// GetSharingOption mocks base method.
func (m *MockTripsRepo) GetSharingOption(ctx context.Context, userId, tripId uint) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSharingOption", ctx, userId, tripId)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSharingOption indicates an expected call of GetSharingOption.
func (mr *MockTripsRepoMockRecorder) GetSharingOption(ctx, userId, tripId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSharingOption", reflect.TypeOf((*MockTripsRepo)(nil).GetSharingOption), ctx, userId, tripId)
}

// GetSharingToken mocks base method.
func (m *MockTripsRepo) GetSharingToken(ctx context.Context, tripID uint) (models.SharingToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSharingToken", ctx, tripID)
	ret0, _ := ret[0].(models.SharingToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSharingToken indicates an expected call of GetSharingToken.
func (mr *MockTripsRepoMockRecorder) GetSharingToken(ctx, tripID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSharingToken", reflect.TypeOf((*MockTripsRepo)(nil).GetSharingToken), ctx, tripID)
}

// GetTrip mocks base method.
func (m *MockTripsRepo) GetTrip(ctx context.Context, tripID uint) (models.Trip, []models.UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrip", ctx, tripID)
	ret0, _ := ret[0].(models.Trip)
	ret1, _ := ret[1].([]models.UserProfile)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTrip indicates an expected call of GetTrip.
func (mr *MockTripsRepoMockRecorder) GetTrip(ctx, tripID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrip", reflect.TypeOf((*MockTripsRepo)(nil).GetTrip), ctx, tripID)
}

// GetTripBySharingToken mocks base method.
func (m *MockTripsRepo) GetTripBySharingToken(ctx context.Context, token string) (models.Trip, []models.UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTripBySharingToken", ctx, token)
	ret0, _ := ret[0].(models.Trip)
	ret1, _ := ret[1].([]models.UserProfile)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTripBySharingToken indicates an expected call of GetTripBySharingToken.
func (mr *MockTripsRepoMockRecorder) GetTripBySharingToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTripBySharingToken", reflect.TypeOf((*MockTripsRepo)(nil).GetTripBySharingToken), ctx, token)
}

// GetTripsByUserID mocks base method.
func (m *MockTripsRepo) GetTripsByUserID(ctx context.Context, userID uint, limit, offset int) ([]models.Trip, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTripsByUserID", ctx, userID, limit, offset)
	ret0, _ := ret[0].([]models.Trip)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTripsByUserID indicates an expected call of GetTripsByUserID.
func (mr *MockTripsRepoMockRecorder) GetTripsByUserID(ctx, userID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTripsByUserID", reflect.TypeOf((*MockTripsRepo)(nil).GetTripsByUserID), ctx, userID, limit, offset)
}

// UpdateTrip mocks base method.
func (m *MockTripsRepo) UpdateTrip(ctx context.Context, user models.Trip) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrip", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTrip indicates an expected call of UpdateTrip.
func (mr *MockTripsRepoMockRecorder) UpdateTrip(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrip", reflect.TypeOf((*MockTripsRepo)(nil).UpdateTrip), ctx, user)
}
