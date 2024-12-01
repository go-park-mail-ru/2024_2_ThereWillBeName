package grpc

import (
	"2024_2_ThereWillBeName/internal/models"
	tripsGen "2024_2_ThereWillBeName/internal/pkg/trips/delivery/grpc/gen"
	mock "2024_2_ThereWillBeName/internal/pkg/trips/mocks"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTrip(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock.NewMockTripsUsecase(ctrl)
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handl := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handl)

	handler := NewGrpcTripHandler(mockUsecase, logger)

	tests := []struct {
		name         string
		input        *tripsGen.CreateTripRequest
		mockBehavior func()
		expectedErr  error
	}{
		{
			name: "Success",
			input: &tripsGen.CreateTripRequest{
				Trip: &tripsGen.Trip{
					UserId:      1,
					Name:        "Trip to Paris",
					Description: "A nice trip",
					CityId:      2,
					StartDate:   "2024-12-01",
					EndDate:     "2024-12-10",
					Private:     false,
				},
			},
			mockBehavior: func() {
				mockUsecase.EXPECT().CreateTrip(gomock.Any(), models.Trip{
					UserID:      1,
					Name:        "Trip to Paris",
					Description: "A nice trip",
					CityID:      2,
					StartDate:   "2024-12-01",
					EndDate:     "2024-12-10",
					Private:     false,
				}).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "Usecase Error",
			input: &tripsGen.CreateTripRequest{
				Trip: &tripsGen.Trip{
					UserId:      1,
					Name:        "Trip to Paris",
					Description: "A nice trip",
					CityId:      2,
					StartDate:   "2024-12-01",
					EndDate:     "2024-12-10",
					Private:     false,
				},
			},
			mockBehavior: func() {
				mockUsecase.EXPECT().CreateTrip(gomock.Any(), models.Trip{
					UserID:      1,
					Name:        "Trip to Paris",
					Description: "A nice trip",
					CityID:      2,
					StartDate:   "2024-12-01",
					EndDate:     "2024-12-10",
					Private:     false,
				}).Return(errors.New("usecase error"))
			},
			expectedErr: errors.New("usecase error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			_, err := handler.CreateTrip(context.Background(), tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateTrip(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock.NewMockTripsUsecase(ctrl)
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handl := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handl)

	handler := NewGrpcTripHandler(mockUsecase, logger)

	tests := []struct {
		name         string
		input        *tripsGen.UpdateTripRequest
		mockBehavior func()
		expectedErr  error
	}{
		{
			name: "Success",
			input: &tripsGen.UpdateTripRequest{
				Trip: &tripsGen.Trip{
					UserId:      1,
					Name:        "Updated Trip",
					Description: "Updated description",
					CityId:      2,
					StartDate:   "2024-12-01",
					EndDate:     "2024-12-10",
					Private:     false,
				},
			},
			mockBehavior: func() {
				mockUsecase.EXPECT().UpdateTrip(gomock.Any(), models.Trip{
					UserID:      1,
					Name:        "Updated Trip",
					Description: "Updated description",
					CityID:      2,
					StartDate:   "2024-12-01",
					EndDate:     "2024-12-10",
					Private:     false,
				}).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "Usecase Error",
			input: &tripsGen.UpdateTripRequest{
				Trip: &tripsGen.Trip{
					UserId:      1,
					Name:        "Updated Trip",
					Description: "Updated description",
					CityId:      2,
					StartDate:   "2024-12-01",
					EndDate:     "2024-12-10",
					Private:     false,
				},
			},
			mockBehavior: func() {
				mockUsecase.EXPECT().UpdateTrip(gomock.Any(), models.Trip{
					UserID:      1,
					Name:        "Updated Trip",
					Description: "Updated description",
					CityID:      2,
					StartDate:   "2024-12-01",
					EndDate:     "2024-12-10",
					Private:     false,
				}).Return(errors.New("usecase error"))
			},
			expectedErr: errors.New("usecase error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			_, err := handler.UpdateTrip(context.Background(), tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
func TestDeleteTrip(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock.NewMockTripsUsecase(ctrl)
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handl := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handl)

	handler := NewGrpcTripHandler(mockUsecase, logger)

	tests := []struct {
		name         string
		input        *tripsGen.DeleteTripRequest
		mockBehavior func()
		expectedErr  error
	}{
		{
			name: "Success",
			input: &tripsGen.DeleteTripRequest{
				Id: 123,
			},
			mockBehavior: func() {
				mockUsecase.EXPECT().DeleteTrip(gomock.Any(), uint(123)).Return(nil)
			},
			expectedErr: nil,
		},
		{
			name: "Usecase Error",
			input: &tripsGen.DeleteTripRequest{
				Id: 123,
			},
			mockBehavior: func() {
				mockUsecase.EXPECT().DeleteTrip(gomock.Any(), uint(123)).Return(errors.New("usecase error"))
			},
			expectedErr: errors.New("usecase error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			_, err := handler.DeleteTrip(context.Background(), tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetTripsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock.NewMockTripsUsecase(ctrl)
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handl := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handl)

	handler := NewGrpcTripHandler(mockUsecase, logger)

	tests := []struct {
		name         string
		input        *tripsGen.GetTripsByUserIDRequest
		mockBehavior func()
		expectedResp *tripsGen.GetTripsByUserIDResponse
		expectedErr  error
	}{
		{
			name: "Success",
			input: &tripsGen.GetTripsByUserIDRequest{
				UserId: 1,
				Limit:  10,
				Offset: 0,
			},
			mockBehavior: func() {
				mockUsecase.EXPECT().GetTripsByUserID(gomock.Any(), uint(1), 10, 0).Return([]models.Trip{
					{
						ID:          1,
						UserID:      1,
						Name:        "Trip 1",
						Description: "Description 1",
						CityID:      2,
						StartDate:   "2024-12-01",
						EndDate:     "2024-12-10",
						Photos:      []string{"photo1.jpg", "photo2.jpg"},
						Private:     false,
					},
				}, nil)
			},
			expectedResp: &tripsGen.GetTripsByUserIDResponse{
				Trips: []*tripsGen.Trip{
					{
						Id:          1,
						UserId:      1,
						Name:        "Trip 1",
						Description: "Description 1",
						CityId:      2,
						StartDate:   "2024-12-01",
						EndDate:     "2024-12-10",
						Photos:      []string{"photo1.jpg", "photo2.jpg"},
						Private:     false,
					},
				},
			},
			expectedErr: nil,
		},
		{
			name: "Usecase Error",
			input: &tripsGen.GetTripsByUserIDRequest{
				UserId: 1,
				Limit:  10,
				Offset: 0,
			},
			mockBehavior: func() {
				mockUsecase.EXPECT().GetTripsByUserID(gomock.Any(), uint(1), 10, 0).Return(nil, errors.New("usecase error"))
			},
			expectedResp: nil,
			expectedErr:  errors.New("usecase error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			resp, err := handler.GetTripsByUserID(context.Background(), tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}
func TestGetTrip(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock.NewMockTripsUsecase(ctrl)
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handl := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handl)

	handler := NewGrpcTripHandler(mockUsecase, logger)

	tests := []struct {
		name         string
		input        *tripsGen.GetTripRequest
		mockBehavior func()
		expectedResp *tripsGen.GetTripResponse
		expectedErr  error
	}{
		{
			name: "Success",
			input: &tripsGen.GetTripRequest{
				TripId: 1,
			},
			mockBehavior: func() {
				mockUsecase.EXPECT().GetTrip(gomock.Any(), uint(1)).Return(models.Trip{
					ID:          1,
					UserID:      1,
					Name:        "Trip to Paris",
					Description: "A nice trip",
					CityID:      2,
					StartDate:   "2024-12-01",
					EndDate:     "2024-12-10",
					Private:     false,
					Photos:      []string{"photo1.jpg", "photo2.jpg"},
				}, nil)
			},
			expectedResp: &tripsGen.GetTripResponse{
				Trip: &tripsGen.Trip{
					Id:          1,
					UserId:      1,
					Name:        "Trip to Paris",
					Description: "A nice trip",
					CityId:      2,
					StartDate:   "2024-12-01",
					EndDate:     "2024-12-10",
					Private:     false,
					Photos:      []string{"photo1.jpg", "photo2.jpg"},
				},
			},
			expectedErr: nil,
		},
		{
			name: "Usecase Error",
			input: &tripsGen.GetTripRequest{
				TripId: 1,
			},
			mockBehavior: func() {
				mockUsecase.EXPECT().GetTrip(gomock.Any(), uint(1)).Return(models.Trip{}, errors.New("usecase error"))
			},
			expectedResp: nil,
			expectedErr:  errors.New("usecase error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			resp, err := handler.GetTrip(context.Background(), tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}

func setupTestStorage() (string, func()) {
	dir, err := os.MkdirTemp("", "photo_storage_test")
	if err != nil {
		panic(fmt.Sprintf("Failed to create temp dir: %v", err))
	}

	os.Setenv("PHOTO_STORAGE_PATH", dir)

	return dir, func() {
		os.RemoveAll(dir)
	}
}

func TestAddPlaceToTrip(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mock.NewMockTripsUsecase(ctrl)
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handl := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handl)

	handler := NewGrpcTripHandler(mockUsecase, logger)

	tests := []struct {
		name         string
		input        *tripsGen.AddPlaceToTripRequest
		mockBehavior func()
		expectedResp *tripsGen.EmptyResponse
		expectedErr  error
	}{
		{
			name: "Success",
			input: &tripsGen.AddPlaceToTripRequest{
				TripId:  1,
				PlaceId: 2,
			},
			mockBehavior: func() {
				mockUsecase.EXPECT().AddPlaceToTrip(gomock.Any(), uint(1), uint(2)).Return(nil)
			},
			expectedResp: &tripsGen.EmptyResponse{},
			expectedErr:  nil,
		},
		{
			name: "Usecase Error",
			input: &tripsGen.AddPlaceToTripRequest{
				TripId:  1,
				PlaceId: 2,
			},
			mockBehavior: func() {
				mockUsecase.EXPECT().AddPlaceToTrip(gomock.Any(), uint(1), uint(2)).Return(errors.New("usecase error"))
			},
			expectedResp: nil,
			expectedErr:  errors.New("usecase error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			resp, err := handler.AddPlaceToTrip(context.Background(), tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, resp)
			}
		})
	}
}
func TestAddPhotosToTrip(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dir, cleanup := setupTestStorage()
	defer cleanup()

	mockUsecase := mock.NewMockTripsUsecase(ctrl)
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handl := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handl)

	handler := NewGrpcTripHandler(mockUsecase, logger)

	tests := []struct {
		name         string
		input        *tripsGen.AddPhotosToTripRequest
		mockBehavior func()
		expectedResp *tripsGen.AddPhotosToTripResponse
		expectedErr  error
	}{
		{
			name: "Success",
			input: &tripsGen.AddPhotosToTripRequest{
				TripId: 1,
				Photos: []string{"data:image/jpeg;base64,/9j/4AAQSkZJRgABAQEA"},
			},
			mockBehavior: func() {
				mockUsecase.EXPECT().AddPhotosToTrip(gomock.Any(), uint(1), gomock.Any()).Return(nil)
			},
			expectedResp: &tripsGen.AddPhotosToTripResponse{
				Photos: []*tripsGen.Photo{
					{PhotoPath: filepath.Join(dir, "trip_1")},
				},
			},
			expectedErr: nil,
		},
		{
			name: "Invalid Base64 Data",
			input: &tripsGen.AddPhotosToTripRequest{
				TripId: 1,
				Photos: []string{"invalid-base64"},
			},
			mockBehavior: func() {
			},
			expectedResp: nil,
			expectedErr:  errors.New("invalid base64 data"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			resp, err := handler.AddPhotosToTrip(context.Background(), tt.input)

			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}
