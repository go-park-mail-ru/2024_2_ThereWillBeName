package usecase

// import (
// 	"2024_2_ThereWillBeName/internal/models"
// 	mock "2024_2_ThereWillBeName/internal/pkg/trips/mocks"
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestTripsUsecaseImpl_CreateTrip(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mock.NewMockTripsRepo(ctrl)
// 	usecase := NewTripsUsecase(mockRepo)

// 	tests := []struct {
// 		name         string
// 		inputTrip    models.Trip
// 		mockBehavior func()
// 		expectedErr  error
// 	}{
// 		{
// 			name: "Success",
// 			inputTrip: models.Trip{
// 				UserID:      1,
// 				Name:        "Trip to Paris",
// 				Description: "A nice trip",
// 				CityID:      2,
// 				StartDate:   "2024-12-01",
// 				EndDate:     "2024-12-10",
// 			},
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().CreateTrip(gomock.Any(), gomock.Any()).Return(nil)
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name: "Repository Error - Not Found",
// 			inputTrip: models.Trip{
// 				UserID:      1,
// 				Name:        "Trip to Paris",
// 				Description: "A nice trip",
// 				CityID:      2,
// 				StartDate:   "2024-12-01",
// 				EndDate:     "2024-12-10",
// 			},
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().CreateTrip(gomock.Any(), gomock.Any()).Return(models.ErrNotFound)
// 			},
// 			expectedErr: models.ErrNotFound,
// 		},
// 		{
// 			name: "Repository Error - Internal",
// 			inputTrip: models.Trip{
// 				UserID:      1,
// 				Name:        "Trip to Paris",
// 				Description: "A nice trip",
// 				CityID:      2,
// 				StartDate:   "2024-12-01",
// 				EndDate:     "2024-12-10",
// 			},
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().CreateTrip(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
// 			},
// 			expectedErr: models.ErrInternal,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mockBehavior()

// 			err := usecase.CreateTrip(context.Background(), tt.inputTrip)

// 			if tt.expectedErr != nil {
// 				assert.Error(t, err)
// 				assert.True(t, errors.Is(err, tt.expectedErr))
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func TestTripsUsecaseImpl_UpdateTrip(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mock.NewMockTripsRepo(ctrl)
// 	usecase := NewTripsUsecase(mockRepo)

// 	tests := []struct {
// 		name         string
// 		inputTrip    models.Trip
// 		mockBehavior func()
// 		expectedErr  error
// 	}{
// 		{
// 			name: "Success",
// 			inputTrip: models.Trip{
// 				ID:          1,
// 				UserID:      1,
// 				Name:        "Updated Trip",
// 				Description: "Updated Description",
// 				CityID:      2,
// 				StartDate:   "2024-12-01",
// 				EndDate:     "2024-12-10",
// 			},
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().UpdateTrip(gomock.Any(), gomock.Any()).Return(nil)
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name: "Repository Error - Not Found",
// 			inputTrip: models.Trip{
// 				ID:          1,
// 				UserID:      1,
// 				Name:        "Updated Trip",
// 				Description: "Updated Description",
// 				CityID:      2,
// 				StartDate:   "2024-12-01",
// 				EndDate:     "2024-12-10",
// 			},
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().UpdateTrip(gomock.Any(), gomock.Any()).Return(models.ErrNotFound)
// 			},
// 			expectedErr: models.ErrNotFound,
// 		},
// 		{
// 			name: "Repository Error - Internal",
// 			inputTrip: models.Trip{
// 				ID:          1,
// 				UserID:      1,
// 				Name:        "Updated Trip",
// 				Description: "Updated Description",
// 				CityID:      2,
// 				StartDate:   "2024-12-01",
// 				EndDate:     "2024-12-10",
// 			},
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().UpdateTrip(gomock.Any(), gomock.Any()).Return(errors.New("internal error"))
// 			},
// 			expectedErr: models.ErrInternal,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mockBehavior()

// 			err := usecase.UpdateTrip(context.Background(), tt.inputTrip)

// 			if tt.expectedErr != nil {
// 				assert.Error(t, err)
// 				assert.True(t, errors.Is(err, tt.expectedErr))
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func TestTripsUsecaseImpl_DeleteTrip(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mock.NewMockTripsRepo(ctrl)
// 	usecase := NewTripsUsecase(mockRepo)

// 	tests := []struct {
// 		name         string
// 		inputID      uint
// 		mockBehavior func()
// 		expectedErr  error
// 	}{
// 		{
// 			name:    "Success",
// 			inputID: 1,
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().DeleteTrip(gomock.Any(), uint(1)).Return(nil)
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name:    "Repository Error - Not Found",
// 			inputID: 1,
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().DeleteTrip(gomock.Any(), uint(1)).Return(models.ErrNotFound)
// 			},
// 			expectedErr: models.ErrNotFound,
// 		},
// 		{
// 			name:    "Repository Error - Internal",
// 			inputID: 1,
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().DeleteTrip(gomock.Any(), uint(1)).Return(errors.New("internal error"))
// 			},
// 			expectedErr: models.ErrInternal,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mockBehavior()

// 			err := usecase.DeleteTrip(context.Background(), tt.inputID)

// 			if tt.expectedErr != nil {
// 				assert.Error(t, err)
// 				assert.True(t, errors.Is(err, tt.expectedErr))
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func TestTripsUsecaseImpl_GetTripsByUserID(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mock.NewMockTripsRepo(ctrl)
// 	usecase := NewTripsUsecase(mockRepo)

// 	tests := []struct {
// 		name          string
// 		userID        uint
// 		limit         int
// 		offset        int
// 		mockBehavior  func()
// 		expectedTrips []models.Trip
// 		expectedErr   error
// 	}{
// 		{
// 			name:   "Success",
// 			userID: 1,
// 			limit:  10,
// 			offset: 0,
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().GetTripsByUserID(gomock.Any(), uint(1), 10, 0).Return([]models.Trip{
// 					{ID: 1, UserID: 1, Name: "Trip 1"},
// 					{ID: 2, UserID: 1, Name: "Trip 2"},
// 				}, nil)
// 			},
// 			expectedTrips: []models.Trip{
// 				{ID: 1, UserID: 1, Name: "Trip 1"},
// 				{ID: 2, UserID: 1, Name: "Trip 2"},
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name:   "Repository Error - Not Found",
// 			userID: 1,
// 			limit:  10,
// 			offset: 0,
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().GetTripsByUserID(gomock.Any(), uint(1), 10, 0).Return(nil, models.ErrNotFound)
// 			},
// 			expectedTrips: nil,
// 			expectedErr:   models.ErrNotFound,
// 		},
// 		{
// 			name:   "Repository Error - Internal",
// 			userID: 1,
// 			limit:  10,
// 			offset: 0,
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().GetTripsByUserID(gomock.Any(), uint(1), 10, 0).Return(nil, errors.New("internal error"))
// 			},
// 			expectedTrips: nil,
// 			expectedErr:   models.ErrInternal,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mockBehavior()

// 			trips, err := usecase.GetTripsByUserID(context.Background(), tt.userID, tt.limit, tt.offset)

// 			if tt.expectedErr != nil {
// 				assert.Error(t, err)
// 				assert.True(t, errors.Is(err, tt.expectedErr))
// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, tt.expectedTrips, trips)
// 			}
// 		})
// 	}
// }

// func TestTripsUsecaseImpl_GetTrip(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mock.NewMockTripsRepo(ctrl)
// 	usecase := NewTripsUsecase(mockRepo)

// 	tests := []struct {
// 		name         string
// 		tripID       uint
// 		mockBehavior func()
// 		expectedTrip models.Trip
// 		expectedErr  error
// 	}{
// 		{
// 			name:   "Success",
// 			tripID: 1,
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().GetTrip(gomock.Any(), uint(1)).Return(models.Trip{
// 					ID:   1,
// 					Name: "Trip 1",
// 				}, nil)
// 			},
// 			expectedTrip: models.Trip{
// 				ID:   1,
// 				Name: "Trip 1",
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name:   "Repository Error - Not Found",
// 			tripID: 1,
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().GetTrip(gomock.Any(), uint(1)).Return(models.Trip{}, models.ErrNotFound)
// 			},
// 			expectedTrip: models.Trip{},
// 			expectedErr:  models.ErrNotFound,
// 		},
// 		{
// 			name:   "Repository Error - Internal",
// 			tripID: 1,
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().GetTrip(gomock.Any(), uint(1)).Return(models.Trip{}, errors.New("internal error"))
// 			},
// 			expectedTrip: models.Trip{},
// 			expectedErr:  models.ErrInternal,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mockBehavior()

// 			trip, err := usecase.GetTrip(context.Background(), tt.tripID)

// 			if tt.expectedErr != nil {
// 				assert.Error(t, err)
// 				assert.True(t, errors.Is(err, tt.expectedErr))
// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, tt.expectedTrip, trip)
// 			}
// 		})
// 	}
// }
// func TestTripsUsecaseImpl_AddPlaceToTrip(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mock.NewMockTripsRepo(ctrl)
// 	usecase := NewTripsUsecase(mockRepo)

// 	tests := []struct {
// 		name         string
// 		tripID       uint
// 		placeID      uint
// 		mockBehavior func()
// 		expectedErr  error
// 	}{
// 		{
// 			name:    "Success",
// 			tripID:  1,
// 			placeID: 2,
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().AddPlaceToTrip(gomock.Any(), uint(1), uint(2)).Return(nil)
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name:    "Repository Error - Not Found",
// 			tripID:  1,
// 			placeID: 2,
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().AddPlaceToTrip(gomock.Any(), uint(1), uint(2)).Return(models.ErrNotFound)
// 			},
// 			expectedErr: models.ErrNotFound,
// 		},
// 		{
// 			name:    "Repository Error - Internal",
// 			tripID:  1,
// 			placeID: 2,
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().AddPlaceToTrip(gomock.Any(), uint(1), uint(2)).Return(errors.New("internal error"))
// 			},
// 			expectedErr: models.ErrInternal,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mockBehavior()

// 			err := usecase.AddPlaceToTrip(context.Background(), tt.tripID, tt.placeID)

// 			if tt.expectedErr != nil {
// 				assert.Error(t, err)
// 				assert.True(t, errors.Is(err, tt.expectedErr))
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func TestTripsUsecaseImpl_AddPhotosToTrip(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mock.NewMockTripsRepo(ctrl)
// 	usecase := NewTripsUsecase(mockRepo)

// 	tests := []struct {
// 		name         string
// 		tripID       uint
// 		photoPaths   []string
// 		mockBehavior func()
// 		expectedErr  error
// 	}{
// 		{
// 			name:       "Success",
// 			tripID:     1,
// 			photoPaths: []string{"path/to/photo1.jpg", "path/to/photo2.jpg"},
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().AddPhotoToTrip(gomock.Any(), uint(1), "photo1.jpg").Return(nil)
// 				mockRepo.EXPECT().AddPhotoToTrip(gomock.Any(), uint(1), "photo2.jpg").Return(nil)
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name:       "Repository Error - AddPhotoToTrip Fails",
// 			tripID:     1,
// 			photoPaths: []string{"path/to/photo1.jpg", "path/to/photo2.jpg"},
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().AddPhotoToTrip(gomock.Any(), uint(1), "photo1.jpg").Return(errors.New("internal error"))
// 			},
// 			expectedErr: errors.New("failed to add photo to trip: internal error"),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mockBehavior()

// 			err := usecase.AddPhotosToTrip(context.Background(), tt.tripID, tt.photoPaths)

// 			if tt.expectedErr != nil {
// 				assert.Error(t, err)
// 				assert.Contains(t, err.Error(), tt.expectedErr.Error())
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }

// func TestTripsUsecaseImpl_DeletePhotoFromTrip(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mock.NewMockTripsRepo(ctrl)
// 	usecase := NewTripsUsecase(mockRepo)

// 	tests := []struct {
// 		name         string
// 		tripID       uint
// 		photoPath    string
// 		mockBehavior func()
// 		expectedErr  error
// 	}{
// 		{
// 			name:      "Success",
// 			tripID:    1,
// 			photoPath: "photo.jpg",
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().DeletePhotoFromTrip(gomock.Any(), uint(1), "photo.jpg").Return(nil)
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name:      "Repository Error - DeletePhotoFromTrip Fails",
// 			tripID:    1,
// 			photoPath: "photo.jpg",
// 			mockBehavior: func() {
// 				mockRepo.EXPECT().DeletePhotoFromTrip(gomock.Any(), uint(1), "photo.jpg").Return(errors.New("internal error"))
// 			},
// 			expectedErr: errors.New("failed to delete photo from database: internal error"),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.mockBehavior()

// 			err := usecase.DeletePhotoFromTrip(context.Background(), tt.tripID, tt.photoPath)

// 			if tt.expectedErr != nil {
// 				assert.Error(t, err)
// 				assert.Contains(t, err.Error(), tt.expectedErr.Error())
// 			} else {
// 				assert.NoError(t, err)
// 			}
// 		})
// 	}
// }
