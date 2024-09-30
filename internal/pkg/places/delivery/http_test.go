package delivery

import (
	"2024_2_ThereWillBeName/internal/models"
	mockplaces "2024_2_ThereWillBeName/internal/pkg/places/mocks"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlaceHandler(t *testing.T) {
	places := []models.Place{
		{ID: 1, Name: "Place 1", Image: "/image1.png", Description: "1"},
		{ID: 2, Name: "Place 2", Image: "/image2.png", Description: "2"},
	}
	jsonPlaces, _ := json.Marshal(places)
	stringPlaces := string(jsonPlaces)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUsecase := mockplaces.NewMockPlaceUsecase(ctrl)
	handler := NewPlacesHandler(mockUsecase)

	tests := []struct {
		name         string
		mockReturn   []models.Place
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Success",
			mockReturn:   places,
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: stringPlaces,
		},
		{
			name:         "Error",
			mockReturn:   nil,
			mockError:    assert.AnError,
			expectedCode: http.StatusInternalServerError,
			expectedBody: "Не удалось получить список достопримечательностей\n",
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			mockUsecase.EXPECT().GetPlaces(gomock.Any()).Return(testcase.mockReturn, testcase.mockError)

			req, err := http.NewRequest(http.MethodGet, "/places", nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.GetPlaceHandler(rr, req)

			assert.Equal(t, testcase.expectedCode, rr.Code)
			assert.Equal(t, testcase.expectedBody, rr.Body.String())
		})
	}
}
