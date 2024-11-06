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

func TestPlaceUsecase_CreatePlace(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock_places.NewMockPlaceRepo(ctrl)

	place := models.CreatePlace{
		Name:            "Central Park",
		ImagePath:       "/images/central_park.jpg",
		Description:     "A large public park in New York City, offering a variety of recreational activities.",
		Rating:          5,
		NumberOfReviews: 2500,
		Address:         "59th St to 110th St, New York, NY 10022",
		CityId:          1,
		PhoneNumber:     "+1 212-310-6600",
		CategoriesId:    []int{1, 2},
	}
	tests := []struct {
		name          string
		mockinput     models.CreatePlace
		mockError     error
		expectedError error
	}{
		{
			name:          "Success",
			mockinput:     place,
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:          "Error",
			mockinput:     models.CreatePlace{},
			mockError:     errors.New("error"),
			expectedError: errors.New("error"),
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			repo.EXPECT().CreatePlace(context.Background(), testCase.mockinput).Return(testCase.mockError)
			u := NewPlaceUsecase(repo)
			err := u.CreatePlace(context.Background(), testCase.mockinput)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestPlaceUsecase_DeletePlace(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock_places.NewMockPlaceRepo(ctrl)

	tests := []struct {
		name          string
		mockinput     uint
		mockError     error
		expectedError error
	}{
		{
			name:          "Success",
			mockinput:     1,
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:          "Error",
			mockinput:     1,
			mockError:     errors.New("error"),
			expectedError: errors.New("error"),
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			repo.EXPECT().DeletePlace(context.Background(), testCase.mockinput).Return(testCase.mockError)
			u := NewPlaceUsecase(repo)
			err := u.DeletePlace(context.Background(), testCase.mockinput)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestPlaceUsecase_UpdatePlace(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock_places.NewMockPlaceRepo(ctrl)

	place := models.UpdatePlace{
		Name:            "Central Park",
		ImagePath:       "/images/central_park.jpg",
		Description:     "A large public park in New York City, offering a variety of recreational activities.",
		Rating:          5,
		NumberOfReviews: 2500,
		Address:         "59th St to 110th St, New York, NY 10022",
		CityId:          1,
		PhoneNumber:     "+1 212-310-6600",
		CategoriesId:    []int{1, 2},
	}
	tests := []struct {
		name          string
		mockinput     models.UpdatePlace
		mockError     error
		expectedError error
	}{
		{
			name:          "Success",
			mockinput:     place,
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:          "Error",
			mockinput:     models.UpdatePlace{},
			mockError:     errors.New("error"),
			expectedError: errors.New("error"),
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			repo.EXPECT().UpdatePlace(context.Background(), testCase.mockinput).Return(testCase.mockError)
			u := NewPlaceUsecase(repo)
			err := u.UpdatePlace(context.Background(), testCase.mockinput)
			assert.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestPlaceUsecase_GetPlace(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock_places.NewMockPlaceRepo(ctrl)

	place := models.GetPlace{
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
	}
	tests := []struct {
		name          string
		mockinput     uint
		mockOutput    models.GetPlace
		mockError     error
		expectedError error
	}{
		{
			name:          "Success",
			mockinput:     1,
			mockError:     nil,
			expectedError: nil,
			mockOutput:    place,
		},
		{
			name:          "Error",
			mockinput:     1,
			mockError:     errors.New("error"),
			expectedError: errors.New("error"),
			mockOutput:    models.GetPlace{},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			repo.EXPECT().GetPlace(context.Background(), testCase.mockinput).Return(testCase.mockOutput, testCase.mockError)
			u := NewPlaceUsecase(repo)
			got, err := u.GetPlace(context.Background(), testCase.mockinput)
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.mockOutput, got)
		})
	}
}

func TestPlaceUsecase_SearchPlaces(t *testing.T) {
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
		mockinput     string
		mockError     error
		expectedError error
		mockOutput    []models.GetPlace
	}{
		{
			name:          "Success",
			mockinput:     "search",
			mockError:     nil,
			expectedError: nil,
			mockOutput:    places,
		},
		{
			name:          "Error",
			mockinput:     "search",
			mockError:     errors.New("error"),
			expectedError: errors.New("error"),
			mockOutput:    []models.GetPlace{},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			repo.EXPECT().SearchPlaces(context.Background(), testCase.mockinput, gomock.Any(), gomock.Any()).Return(testCase.mockOutput, testCase.mockError)
			u := NewPlaceUsecase(repo)
			got, err := u.SearchPlaces(context.Background(), testCase.mockinput, 10, 0)
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.mockOutput, got)
		})
	}
}
