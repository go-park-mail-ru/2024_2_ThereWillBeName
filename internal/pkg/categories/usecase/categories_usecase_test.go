package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	mock_categories "2024_2_ThereWillBeName/internal/pkg/categories/mocks"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoriesUsecase_GetCategories(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mock_categories.NewMockCategoriesRepository(ctrl)

	tests := []struct {
		name          string
		mockReturn    []models.Category
		mockError     error
		expectedCode  []models.Category
		expectedError error
	}{
		{
			name: "Success",
			mockReturn: []models.Category{
				{ID: 1, Name: "театр"},
				{ID: 2, Name: "собор"},
			},
			mockError: nil,
			expectedCode: []models.Category{
				{ID: 1, Name: "театр"},
				{ID: 2, Name: "собор"},
			},
			expectedError: nil,
		},
		{
			name:          "Error",
			mockReturn:    nil,
			mockError:     errors.New("error"),
			expectedCode:  []models.Category(nil),
			expectedError: errors.New("error"),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			repo.EXPECT().GetCategories(context.Background(), gomock.Any(), gomock.Any()).Return(testCase.mockReturn, testCase.mockError)
			u := NewCategoriesUsecase(repo)
			got, err := u.GetCategories(context.Background(), 10, 0)
			assert.Equal(t, testCase.expectedError, err)
			assert.Equal(t, testCase.expectedCode, got)
		})
	}
}
