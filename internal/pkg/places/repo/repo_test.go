package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestPlaceRepository_GetPlaces_ErrorUnmarshal(t *testing.T) {
	badJsonFileData := []byte(`{"invalid": "data"}`)

	originalJsonFileData := jsonFileData
	defer func() { jsonFileData = originalJsonFileData }()
	jsonFileData = badJsonFileData

	r := NewPLaceRepository()
	got, err := r.GetPlaces(context.Background())

	assert.Error(t, err)
	assert.Nil(t, got)
}

func TestPlaceRepository_GetPlaces_EmptyFile(t *testing.T) {
	emptyJsonFileData := []byte(`[]`)

	originalJsonFileData := jsonFileData
	defer func() { jsonFileData = originalJsonFileData }()
	jsonFileData = emptyJsonFileData

	r := NewPLaceRepository()
	got, err := r.GetPlaces(context.Background())

	assert.NoError(t, err)
	assert.Empty(t, got)
}
