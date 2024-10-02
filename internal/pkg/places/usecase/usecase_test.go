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

	places := []models.Place{
		{ID: 1, Name: "Place 1", Image: "/image1.png", Description: "1"},
		{ID: 2, Name: "Place 2", Image: "/image2.png", Description: "2"},
	}

	tests := []struct {
		name          string
		mockReturn    []models.Place
		mockError     error
		expectedCode  []models.Place
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
			expectedCode:  []models.Place(nil),
			expectedError: errors.New("error"),
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			repo.EXPECT().GetPlaces(context.Background()).Return(testCase.mockReturn, testCase.mockError)
			u := NewPlaceUsecase(repo)
			got, err := u.GetPlaces(context.Background())
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedCode, got)
		})
	}
}
