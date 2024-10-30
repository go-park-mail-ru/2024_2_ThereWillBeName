package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	mocks "2024_2_ThereWillBeName/internal/pkg/trips/mocks"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTrip(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTripsRepo(ctrl)
	usecase := NewTripsUsecase(mockRepo)

	tests := []struct {
		name        string
		inputTrip   models.Trip
		usecaseErr  error
		expectedErr error
	}{
		{
			name:        "successful_creation",
			inputTrip:   models.Trip{Name: "Test Trip", UserID: 100},
			usecaseErr:  nil,
			expectedErr: nil,
		},
		{
			name:        "repository error",
			inputTrip:   models.Trip{Name: "Test Trip", UserID: 100},
			usecaseErr:  errors.New("internal error"),
			expectedErr: errors.New("internal error: internal repository error"),
		},
		{
			name:        "not found error",
			inputTrip:   models.Trip{Name: "Test Trip", UserID: 100},
			usecaseErr:  models.ErrNotFound,
			expectedErr: errors.New("invalid request: not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().CreateTrip(gomock.Any(), tt.inputTrip).Return(tt.usecaseErr)

			err := usecase.CreateTrip(context.Background(), tt.inputTrip)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateTrip(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTripsRepo(ctrl)
	usecase := NewTripsUsecase(mockRepo)

	tests := []struct {
		name        string
		inputTrip   models.Trip
		usecaseErr  error
		expectedErr error
	}{
		{
			name:        "successful update",
			inputTrip:   models.Trip{ID: 1, Name: "Updated Trip"},
			usecaseErr:  nil,
			expectedErr: nil,
		},
		{
			name:        "repository error",
			inputTrip:   models.Trip{ID: 1, Name: "Updated Trip"},
			usecaseErr:  errors.New("internal error"),
			expectedErr: errors.New("internal error: internal repository error"),
		},
		{
			name:        "not found error",
			inputTrip:   models.Trip{ID: 1, Name: "Updated Trip"},
			usecaseErr:  models.ErrNotFound,
			expectedErr: errors.New("invalid request: not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().UpdateTrip(gomock.Any(), tt.inputTrip).Return(tt.usecaseErr)

			err := usecase.UpdateTrip(context.Background(), tt.inputTrip)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteTrip(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTripsRepo(ctrl)
	usecase := NewTripsUsecase(mockRepo)

	tests := []struct {
		name        string
		tripID      uint
		usecaseErr  error
		expectedErr error
	}{
		{
			name:        "successful deletion",
			tripID:      1,
			usecaseErr:  nil,
			expectedErr: nil,
		},
		{
			name:        "repository error",
			tripID:      1,
			usecaseErr:  errors.New("internal error"),
			expectedErr: errors.New("internal error: internal repository error"),
		},
		{
			name:        "not found error",
			tripID:      1,
			usecaseErr:  models.ErrNotFound,
			expectedErr: errors.New("invalid request: not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().DeleteTrip(gomock.Any(), tt.tripID).Return(tt.usecaseErr)

			err := usecase.DeleteTrip(context.Background(), tt.tripID)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetTripsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTripsRepo(ctrl)
	usecase := NewTripsUsecase(mockRepo)

	tests := []struct {
		name          string
		userID        uint
		limit         int
		offset        int
		repoTrips     []models.Trip
		usecaseErr    error
		expectedTrips []models.Trip
		expectedErr   error
	}{
		{
			name:          "successful retrieval",
			userID:        1,
			limit:         10,
			offset:        0,
			repoTrips:     []models.Trip{{ID: 1, Name: "Test trip"}},
			usecaseErr:    nil,
			expectedTrips: []models.Trip{{ID: 1, Name: "Test trip"}},
			expectedErr:   nil,
		},
		{
			name:          "repository error",
			userID:        1,
			limit:         10,
			offset:        0,
			repoTrips:     nil,
			usecaseErr:    errors.New("internal error"),
			expectedTrips: nil,
			expectedErr:   errors.New("internal error: internal repository error"),
		},
		{
			name:          "not found error",
			userID:        1,
			limit:         10,
			offset:        0,
			repoTrips:     nil,
			usecaseErr:    models.ErrNotFound,
			expectedTrips: nil,
			expectedErr:   errors.New("invalid request: not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().GetTripsByUserID(gomock.Any(), tt.userID, tt.limit, tt.offset).Return(tt.repoTrips, tt.usecaseErr)

			trips, err := usecase.GetTripsByUserID(context.Background(), tt.userID, tt.limit, tt.offset)

			assert.ElementsMatch(t, tt.expectedTrips, trips)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetTrip(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTripsRepo(ctrl)
	usecase := NewTripsUsecase(mockRepo)

	tests := []struct {
		name         string
		tripID       uint
		repoTrip     models.Trip
		usecaseErr   error
		expectedTrip models.Trip
		expectedErr  error
	}{
		{
			name:         "successful retrieval",
			tripID:       1,
			repoTrip:     models.Trip{ID: 1, Name: "Test trip"},
			usecaseErr:   nil,
			expectedTrip: models.Trip{ID: 1, Name: "Test trip"},
			expectedErr:  nil,
		},
		{
			name:         "repository error",
			tripID:       1,
			repoTrip:     models.Trip{},
			usecaseErr:   errors.New("internal error"),
			expectedTrip: models.Trip{},
			expectedErr:  errors.New("internal error: internal repository error"),
		},
		{
			name:         "not found error",
			tripID:       1,
			repoTrip:     models.Trip{},
			usecaseErr:   models.ErrNotFound,
			expectedTrip: models.Trip{},
			expectedErr:  errors.New("invalid request: not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().GetTrip(gomock.Any(), tt.tripID).Return(tt.repoTrip, tt.usecaseErr)

			trip, err := usecase.GetTrip(context.Background(), tt.tripID)

			assert.Equal(t, tt.expectedTrip, trip)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
