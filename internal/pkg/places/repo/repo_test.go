package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlaceRepository_GetPlaces(t *testing.T) {

	images := make([]models.Place, 0)
	_ = json.Unmarshal(jsonFileData, &images)

	tests := []struct {
		name          string
		mockReturn    []models.Place
		mockError     error
		expectedCode  []models.Place
		expectedError error
	}{
		{
			name:          "Success",
			mockReturn:    images,
			mockError:     nil,
			expectedCode:  images,
			expectedError: nil,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			r := NewPLaceRepository()
			got, err := r.GetPlaces(context.Background())
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedCode, got)
		})
	}
}
