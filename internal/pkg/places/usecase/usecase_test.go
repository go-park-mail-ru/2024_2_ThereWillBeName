package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	mock_places "2024_2_ThereWillBeName/internal/pkg/places/mocks"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlaceUsecase_GetPlaces(t *testing.T) {

	ctrl := gomock.NewController(t)
	repo := mock_places.NewMockPlaceRepo(ctrl)

	places := []models.GetPlace{
		{
			ID:              1,
			Name:            "Central Park",
			ImagePath:       "/images/central_park.jpg",
			Description:     "A large public park in New York City, offering a variety of recreational activities.",
			Rating:          5,
			NumberOfReviews: 2500,
			Address:         "59th St to 110th St, New York, NY 10022",
			City:            "New York",
			PhoneNumber:     "+1 212-310-6600",
			Categories:      []string{"Park", "Recreation", "Nature"},
		},
		{
			ID:              2,
			Name:            "Central Park",
			ImagePath:       "/images/central_park.jpg",
			Description:     "A large public park in New York City, offering a variety of recreational activities.",
			Rating:          5,
			NumberOfReviews: 2500,
			Address:         "59th St to 110th St, New York, NY 10022",
			City:            "New York",
			PhoneNumber:     "+1 212-310-6600",
			Categories:      []string{"Park", "Recreation", "Nature"},
		},
	}

	tests := []struct {
		name          string
		mockReturn    []models.GetPlace
		mockError     error
		expectedCode  []models.GetPlace
		expectedError error
	}{
		{
			name:          "Success",
			mockReturn:    places,
			mockError:     nil,
			expectedCode:  places,
			expectedError: nil,
		},
		{
			name:          "Error",
			mockReturn:    nil,
			mockError:     errors.New("error"),
			expectedCode:  []models.GetPlace(nil),
			expectedError: errors.New("error"),
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			repo.EXPECT().GetPlaces(context.Background(), gomock.Any(), gomock.Any()).Return(testCase.mockReturn, testCase.mockError)
			u := NewPlaceUsecase(repo)
			got, err := u.GetPlaces(context.Background(), 10, 0)
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedCode, got)
		})
	}
}
