package http

import (
	"2024_2_ThereWillBeName/internal/models"
	mockplaces "2024_2_ThereWillBeName/internal/pkg/places/mocks"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetPlacesHandler(t *testing.T) {
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

	jsonPlaces, _ := json.Marshal(places)
	stringPlaces := string(jsonPlaces)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mockplaces.NewMockPlaceUsecase(ctrl)
	handler := NewPlacesHandler(mockUsecase)

	tests := []struct {
		name         string
		mockReturn   []models.GetPlace
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Success",
			mockReturn:   places,
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: stringPlaces + "\n",
		},
		{
			name:         "Error",
			mockReturn:   nil,
			mockError:    assert.AnError,
			expectedCode: http.StatusInternalServerError,
			expectedBody: "",
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			mockUsecase.EXPECT().GetPlaces(gomock.Any(), gomock.Any(), gomock.Any()).Return(testcase.mockReturn, testcase.mockError)

			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/places", bytes.NewBufferString(`{"limit": 10, "offset": 0}`))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.GetPlacesHandler(rr, req)

			assert.Equal(t, testcase.expectedCode, rr.Code)
			assert.Equal(t, testcase.expectedBody, rr.Body.String())
		})
	}
}
