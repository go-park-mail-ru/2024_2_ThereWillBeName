package reviews

import (
	"2024_2_ThereWillBeName/internal/models"
	mock "2024_2_ThereWillBeName/internal/pkg/reviews/mocks"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestReviewsUsecaseImpl_CreateReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockReviewsRepo(ctrl)
	usecase := NewReviewsUsecase(mockRepo)

	tests := []struct {
		name         string
		inputReview  models.Review
		mockBehavior func()
		expectedErr  string
		expectedResp models.GetReview
	}{
		{
			name: "Success",
			inputReview: models.Review{
				UserID:     1,
				PlaceID:    2,
				Rating:     5,
				ReviewText: "Great place!",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().CreateReview(gomock.Any(), gomock.Any()).Return(models.GetReview{
					ID:         1,
					UserLogin:  "john_doe",
					AvatarPath: "/images/john_doe.jpg",
					Rating:     5,
					ReviewText: "Great place!",
				}, nil)
			},
			expectedErr:  "",
			expectedResp: models.GetReview{ID: 1, UserLogin: "john_doe", AvatarPath: "/images/john_doe.jpg", Rating: 5, ReviewText: "Great place!"},
		},
		{
			name: "Repository Error - Not Found",
			inputReview: models.Review{
				UserID:     1,
				PlaceID:    2,
				Rating:     5,
				ReviewText: "Great place!",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().CreateReview(gomock.Any(), gomock.Any()).Return(models.GetReview{}, models.ErrNotFound)
			},
			expectedErr:  "invalid request: not found",
			expectedResp: models.GetReview{},
		},
		{
			name: "Repository Error - Internal",
			inputReview: models.Review{
				UserID:     1,
				PlaceID:    2,
				Rating:     5,
				ReviewText: "Great place!",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().CreateReview(gomock.Any(), gomock.Any()).Return(models.GetReview{}, errors.New("internal repository error"))
			},
			expectedErr:  "internal error: internal repository error",
			expectedResp: models.GetReview{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			createdReview, err := usecase.CreateReview(context.Background(), tt.inputReview)

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
				assert.Equal(t, tt.expectedResp, createdReview)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResp, createdReview)
			}
		})
	}
}

func TestReviewsUsecaseImpl_UpdateReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockReviewsRepo(ctrl)
	usecase := NewReviewsUsecase(mockRepo)

	tests := []struct {
		name         string
		inputReview  models.Review
		mockBehavior func()
		expectedErr  string
	}{
		{
			name: "Success",
			inputReview: models.Review{
				ID:         1,
				UserID:     1,
				PlaceID:    2,
				Rating:     5,
				ReviewText: "Great place!",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().UpdateReview(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedErr: "",
		},
		{
			name: "Repository Error - Not Found",
			inputReview: models.Review{
				ID:         1,
				UserID:     1,
				PlaceID:    2,
				Rating:     5,
				ReviewText: "Great place!",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().UpdateReview(gomock.Any(), gomock.Any()).Return(models.ErrNotFound)
			},
			expectedErr: "invalid request: not found",
		},
		{
			name: "Repository Error - Internal",
			inputReview: models.Review{
				ID:         1,
				UserID:     1,
				PlaceID:    2,
				Rating:     5,
				ReviewText: "Great place!",
			},
			mockBehavior: func() {
				mockRepo.EXPECT().UpdateReview(gomock.Any(), gomock.Any()).Return(errors.New("internal repository error"))
			},
			expectedErr: "internal error: internal repository error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			err := usecase.UpdateReview(context.Background(), tt.inputReview)

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestReviewsUsecaseImpl_DeleteReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockReviewsRepo(ctrl)
	usecase := NewReviewsUsecase(mockRepo)

	tests := []struct {
		name         string
		reviewID     uint
		mockBehavior func()
		expectedErr  string
	}{
		{
			name:     "Success",
			reviewID: 1,
			mockBehavior: func() {
				mockRepo.EXPECT().DeleteReview(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedErr: "",
		},
		{
			name:     "Repository Error - Not Found",
			reviewID: 1,
			mockBehavior: func() {
				mockRepo.EXPECT().DeleteReview(gomock.Any(), gomock.Any()).Return(models.ErrNotFound)
			},
			expectedErr: "invalid request: not found",
		},
		{
			name:     "Repository Error - Internal",
			reviewID: 1,
			mockBehavior: func() {
				mockRepo.EXPECT().DeleteReview(gomock.Any(), gomock.Any()).Return(errors.New("internal repository error"))
			},
			expectedErr: "internal error: internal repository error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			err := usecase.DeleteReview(context.Background(), tt.reviewID)

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestReviewsUsecaseImpl_GetReviewsByPlaceID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockReviewsRepo(ctrl)
	usecase := NewReviewsUsecase(mockRepo)

	placeID := uint(1)
	limit := 10
	offset := 0
	expectedReviews := []models.GetReview{
		{
			ID:         1,
			UserLogin:  "user1",
			AvatarPath: "/path/to/avatar1",
			Rating:     5,
			ReviewText: "Great place!",
		},
		{
			ID:         2,
			UserLogin:  "user2",
			AvatarPath: "/path/to/avatar2",
			Rating:     4,
			ReviewText: "Nice place, could be better.",
		},
	}

	tests := []struct {
		name            string
		mockBehavior    func()
		expectedReviews []models.GetReview
		expectedErr     string
	}{
		{
			name: "Success",
			mockBehavior: func() {
				mockRepo.EXPECT().GetReviewsByPlaceID(gomock.Any(), placeID, limit, offset).Return(expectedReviews, nil)
			},
			expectedReviews: expectedReviews,
			expectedErr:     "",
		},
		{
			name: "Repository Error - Not Found",
			mockBehavior: func() {
				mockRepo.EXPECT().GetReviewsByPlaceID(gomock.Any(), placeID, limit, offset).Return(nil, models.ErrNotFound)
			},
			expectedReviews: nil,
			expectedErr:     "invalid request: not found",
		},
		{
			name: "Repository Error - Internal",
			mockBehavior: func() {
				mockRepo.EXPECT().GetReviewsByPlaceID(gomock.Any(), placeID, limit, offset).Return(nil, errors.New("internal repository error"))
			},
			expectedReviews: nil,
			expectedErr:     "internal error: internal repository error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			reviews, err := usecase.GetReviewsByPlaceID(context.Background(), placeID, limit, offset)

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
				assert.Nil(t, reviews)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedReviews, reviews)
			}
		})
	}
}

func TestReviewsUsecaseImpl_GetReviewsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockReviewsRepo(ctrl)
	usecase := NewReviewsUsecase(mockRepo)

	userID := uint(1)
	limit := 10
	offset := 0
	expectedReviews := []models.GetReviewByUserID{
		{
			ID:         1,
			PlaceName:  "Central Park",
			Rating:     5,
			ReviewText: "Great place!",
		},
		{
			ID:         2,
			PlaceName:  "Empire State Building",
			Rating:     4,
			ReviewText: "Nice place, could be better.",
		},
	}

	tests := []struct {
		name            string
		mockBehavior    func()
		expectedReviews []models.GetReviewByUserID
		expectedErr     string
	}{
		{
			name: "Success",
			mockBehavior: func() {
				mockRepo.EXPECT().GetReviewsByUserID(gomock.Any(), userID, limit, offset).Return(expectedReviews, nil)
			},
			expectedReviews: expectedReviews,
			expectedErr:     "",
		},
		{
			name: "Repository Error - Not Found",
			mockBehavior: func() {
				mockRepo.EXPECT().GetReviewsByUserID(gomock.Any(), userID, limit, offset).Return(nil, models.ErrNotFound)
			},
			expectedReviews: nil,
			expectedErr:     "invalid request: not found",
		},
		{
			name: "Repository Error - Internal",
			mockBehavior: func() {
				mockRepo.EXPECT().GetReviewsByUserID(gomock.Any(), userID, limit, offset).Return(nil, errors.New("internal repository error"))
			},
			expectedReviews: nil,
			expectedErr:     "internal error: internal repository error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			reviews, err := usecase.GetReviewsByUserID(context.Background(), userID, limit, offset)

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
				assert.Nil(t, reviews)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedReviews, reviews)
			}
		})
	}
}

func TestReviewsUsecaseImpl_GetReview(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockReviewsRepo(ctrl)
	usecase := NewReviewsUsecase(mockRepo)

	reviewID := uint(1)
	expectedReview := models.GetReview{
		ID:         1,
		UserLogin:  "user1",
		AvatarPath: "/avatars/user1.jpg",
		Rating:     5,
		ReviewText: "Excellent place!",
	}

	tests := []struct {
		name           string
		mockBehavior   func()
		expectedReview models.GetReview
		expectedErr    string
	}{
		{
			name: "Success",
			mockBehavior: func() {
				mockRepo.EXPECT().GetReview(gomock.Any(), reviewID).Return(expectedReview, nil)
			},
			expectedReview: expectedReview,
			expectedErr:    "",
		},
		{
			name: "Repository Error - Not Found",
			mockBehavior: func() {
				mockRepo.EXPECT().GetReview(gomock.Any(), reviewID).Return(models.GetReview{}, models.ErrNotFound)
			},
			expectedReview: models.GetReview{},
			expectedErr:    "invalid request: not found",
		},
		{
			name: "Repository Error - Internal",
			mockBehavior: func() {
				mockRepo.EXPECT().GetReview(gomock.Any(), reviewID).Return(models.GetReview{}, errors.New("internal repository error"))
			},
			expectedReview: models.GetReview{},
			expectedErr:    "internal error: internal repository error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior()

			review, err := usecase.GetReview(context.Background(), reviewID)

			if tt.expectedErr != "" {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
				assert.Equal(t, tt.expectedReview, review)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedReview, review)
			}
		})
	}
}
