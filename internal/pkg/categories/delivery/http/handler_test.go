package http

import (
	"2024_2_ThereWillBeName/internal/models"
	mock_categories "2024_2_ThereWillBeName/internal/pkg/categories/mocks"
	"context"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"
)

func TestCategoriesHandler_GetCategoriesHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(h)

	mockUsecase := mock_categories.NewMockCategoriesRepository(ctrl)
	handler := NewCategoriesHandler(mockUsecase, logger)

	categories := []models.Category{
		{ID: 1, Name: "театр"},
		{ID: 2, Name: "собор"},
	}
	jsonCategories, _ := json.Marshal(categories)
	stringCategories := string(jsonCategories)

	tests := []struct {
		name         string
		offset       string
		limit        string
		mockPlaces   []models.Category
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid request",
			offset:       "0",
			limit:        "10",
			mockPlaces:   categories,
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: stringCategories + "\n",
		},
		{
			name:         "Invalid offset",
			offset:       "invalid",
			limit:        "10",
			mockPlaces:   nil,
			mockError:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: ``,
		},
		{
			name:         "Invalid limit",
			offset:       "0",
			limit:        "invalid",
			mockPlaces:   nil,
			mockError:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: ``,
		},
		{
			name:         "Internal server error",
			offset:       "0",
			limit:        "10",
			mockPlaces:   nil,
			mockError:    errors.New("internal server error"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: ``,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name != "Invalid offset" && tt.name != "Invalid limit" {
				offset, _ := strconv.Atoi(tt.offset)
				limit, _ := strconv.Atoi(tt.limit)
				mockUsecase.EXPECT().GetCategories(gomock.Any(), limit, offset).Return(tt.mockPlaces, tt.mockError)
			}

			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/categories", nil)
			assert.NoError(t, err)

			query := url.Values{}
			query.Add("offset", tt.offset)
			query.Add("limit", tt.limit)
			req.URL.RawQuery = query.Encode()

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/categories", handler.GetCategoriesHandler)
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())
		})
	}

}
