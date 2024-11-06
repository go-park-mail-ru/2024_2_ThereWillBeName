package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"

	"strconv"
	"testing"

	"2024_2_ThereWillBeName/internal/models"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	mocks "2024_2_ThereWillBeName/internal/pkg/trips/mocks"

	"github.com/gorilla/mux"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTripHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handl := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handl)

	mockUsecase := mocks.NewMockTripsUsecase(ctrl)
	handler := NewTripHandler(mockUsecase, logger)

	tests := []struct {
		name           string
		inputTrip      models.Trip
		usecaseErr     error
		expectedStatus int
		expectedBody   httpresponse.ErrorResponse
	}{
		{
			name: "successful creation",
			inputTrip: models.Trip{

				ID:          0,
				UserID:      100,
				Name:        "Test Trip",
				Description: "A trip for testing",
				CityID:      1,
				StartDate:   "2024-12-01",
				EndDate:     "2024-12-15",
				Private:     false,
			},
			usecaseErr:     nil,
			expectedStatus: http.StatusCreated,
		},
		{
			name: "invalid input data",
			inputTrip: models.Trip{

				ID:        0,
				UserID:    101,
				StartDate: "invalid-date",
				EndDate:   "2024-12-15",
			},
			usecaseErr:     errors.New("validation error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   httpresponse.ErrorResponse{Message: "Failed to create trip"},
		},
		{
			name: "internal server error",
			inputTrip: models.Trip{

				ID:          0,
				UserID:      102,
				Name:        "Error Trip",
				Description: "This trip causes an error",
				CityID:      1,
				StartDate:   "2024-12-01",
				EndDate:     "2024-12-15",
				Private:     true,
			},
			usecaseErr:     models.ErrNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   httpresponse.ErrorResponse{Message: "Invalid request"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase.EXPECT().CreateTrip(gomock.Any(), tt.inputTrip).Return(tt.usecaseErr)

			reqBody, _ := json.Marshal(tt.inputTrip)
			req := httptest.NewRequest("POST", "/trips", bytes.NewReader(reqBody))
			rec := httptest.NewRecorder()

			handler.CreateTripHandler(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			if tt.expectedStatus != http.StatusCreated {
				var response httpresponse.ErrorResponse
				_ = json.NewDecoder(rec.Body).Decode(&response)
				assert.Equal(t, tt.expectedBody.Message, response.Message)
			}
		})
	}
}

func TestUpdateTripHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handl := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handl)
	mockUsecase := mocks.NewMockTripsUsecase(ctrl)
	handler := NewTripHandler(mockUsecase, logger)

	tests := []struct {
		name           string
		inputTrip      models.Trip
		usecaseErr     error
		expectedStatus int
		expectedBody   httpresponse.ErrorResponse
	}{
		{
			name: "successful update",
			inputTrip: models.Trip{
				ID:          1,
				Name:        "Updated Trip",
				UserID:      100,
				Description: "Updated description",
				CityID:      1,
				StartDate:   "2024-12-01",
				EndDate:     "2024-12-15"},
			usecaseErr:     nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid request",
			inputTrip:      models.Trip{ID: 10000},
			usecaseErr:     models.ErrNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   httpresponse.ErrorResponse{Message: "Invalid request"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase.EXPECT().UpdateTrip(gomock.Any(), tt.inputTrip).Return(tt.usecaseErr)

			reqBody, _ := json.Marshal(tt.inputTrip)
			req := httptest.NewRequest("PUT", "/trips/"+strconv.Itoa(int(tt.inputTrip.ID)), bytes.NewReader(reqBody))
			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/trips/{id}", handler.UpdateTripHandler).Methods("PUT")
			r.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedStatus != http.StatusOK {
				var response httpresponse.ErrorResponse
				_ = json.NewDecoder(rec.Body).Decode(&response)
				assert.Equal(t, tt.expectedBody.Message, response.Message)
			}
		})
	}
}

func TestDeleteTripHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handl := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handl)

	mockUsecase := mocks.NewMockTripsUsecase(ctrl)
	handler := NewTripHandler(mockUsecase, logger)

	tests := []struct {
		name           string
		tripID         uint
		usecaseErr     error
		expectedStatus int
		expectedBody   httpresponse.ErrorResponse
	}{
		{
			name:           "successful deletion",
			tripID:         1,
			usecaseErr:     nil,
			expectedStatus: http.StatusNoContent,
		},
		{
			name:           "invalid request",
			tripID:         10000,
			usecaseErr:     models.ErrNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   httpresponse.ErrorResponse{Message: "Invalid request"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase.EXPECT().DeleteTrip(gomock.Any(), tt.tripID).Return(tt.usecaseErr)

			req := httptest.NewRequest("DELETE", "/trips/"+strconv.Itoa(int(tt.tripID)), nil)
			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/trips/{id}", handler.DeleteTripHandler).Methods("DELETE")
			r.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedStatus != http.StatusNoContent {
				var response httpresponse.ErrorResponse
				_ = json.NewDecoder(rec.Body).Decode(&response)
				assert.Equal(t, tt.expectedBody.Message, response.Message)
			}
		})
	}
}

func TestGetTripsByUserIDHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handl := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handl)
	mockUsecase := mocks.NewMockTripsUsecase(ctrl)
	handler := NewTripHandler(mockUsecase, logger)

	tests := []struct {
		name           string
		userID         uint
		expectedTrips  []models.Trip
		usecaseErr     error
		expectedStatus int
		expectedBody   httpresponse.ErrorResponse
	}{
		{
			name:   "successful retrieval",
			userID: 100,
			expectedTrips: []models.Trip{{
				ID:          1,
				UserID:      100,
				Name:        "Test Trip",
				Description: "A trip for testing"}},
			usecaseErr:     nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid request",
			userID:         100,
			expectedTrips:  nil,
			usecaseErr:     models.ErrNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   httpresponse.ErrorResponse{Message: "Invalid request"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase.EXPECT().GetTripsByUserID(gomock.Any(), tt.userID, gomock.Any(), gomock.Any()).Return(tt.expectedTrips, tt.usecaseErr)

			req := httptest.NewRequest("GET", "/users/"+strconv.Itoa(int(tt.userID))+"/trips", nil)
			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/users/{userID}/trips", handler.GetTripsByUserIDHandler).Methods("GET")
			r.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedStatus == http.StatusOK {
				var trips []models.Trip
				_ = json.NewDecoder(rec.Body).Decode(&trips)
				assert.Equal(t, tt.expectedTrips, trips)
			} else {
				var response httpresponse.ErrorResponse
				_ = json.NewDecoder(rec.Body).Decode(&response)
				assert.Equal(t, tt.expectedBody.Message, response.Message)
			}
		})
	}
}

func TestGetTripHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handl := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handl)
	mockUsecase := mocks.NewMockTripsUsecase(ctrl)
	handler := NewTripHandler(mockUsecase, logger)

	tests := []struct {
		name           string
		tripID         uint
		expectedTrip   models.Trip
		usecaseErr     error
		expectedStatus int
		expectedBody   httpresponse.ErrorResponse
	}{
		{
			name:   "successful retrieval",
			tripID: 1,
			expectedTrip: models.Trip{
				ID:          1,
				UserID:      100,
				Name:        "Test Trip",
				Description: "A trip for testing"},
			usecaseErr:     nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "invalid request",
			tripID:         10000,
			usecaseErr:     models.ErrNotFound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   httpresponse.ErrorResponse{Message: "Invalid request"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUsecase.EXPECT().GetTrip(gomock.Any(), tt.tripID).Return(tt.expectedTrip, tt.usecaseErr)

			req := httptest.NewRequest("GET", "/trips/"+strconv.Itoa(int(tt.tripID)), nil)
			rec := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/trips/{id}", handler.GetTripHandler).Methods("GET")
			r.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
			if tt.expectedStatus == http.StatusOK {
				var trip models.Trip
				_ = json.NewDecoder(rec.Body).Decode(&trip)
				assert.Equal(t, tt.expectedTrip, trip)
			} else {
				var response httpresponse.ErrorResponse
				_ = json.NewDecoder(rec.Body).Decode(&response)
				assert.Equal(t, tt.expectedBody.Message, response.Message)
			}
		})
	}
}

// func TestAddPlaceToTripHandler(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}

// 	handl := slog.NewJSONHandler(os.Stdout, opts)

// 	logger := slog.New(handl)

// 	mockUsecase := mocks.NewMockTripsUsecase(ctrl)
// 	handler := NewTripHandler(mockUsecase, logger)

// 	tests := []struct {
// 		name           string
// 		ID             uint
// 		requestBody    string
// 		usecaseErr     error
// 		expectedStatus int
// 		expectedBody   httpresponse.ErrorResponse
// 	}{
// 		{
// 			name:           "successful addition of place",
// 			ID:             1,
// 			requestBody:    `{"place_id": 2}`,
// 			usecaseErr:     nil,
// 			expectedStatus: http.StatusCreated,
// 		},
// 		{
// 			name:           "invalid request body",
// 			ID:             2,
// 			requestBody:    `{"place_id": "invalid"}`,
// 			usecaseErr:     nil,
// 			expectedStatus: http.StatusBadRequest,
// 			expectedBody:   httpresponse.ErrorResponse{Message: "Invalid place ID"},
// 		},
// 		{
// 			name:           "error from usecase",
// 			ID:             3,
// 			requestBody:    `{"place_id": 2}`,
// 			usecaseErr:     errors.New("usecase error"),
// 			expectedStatus: http.StatusBadRequest,
// 			expectedBody:   httpresponse.ErrorResponse{Message: "Invalid trip ID"},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockUsecase.EXPECT().AddPlaceToTrip(gomock.Any(), gomock.Any(), gomock.Any()).Return(tt.usecaseErr)
// 			req := httptest.NewRequest("POST", "/trips/"+strconv.Itoa(int(tt.ID)), bytes.NewReader([]byte(tt.requestBody)))
// 			rec := httptest.NewRecorder()

// 			handler.AddPlaceToTripHandler(rec, req)

// 			assert.Equal(t, tt.expectedStatus, rec.Code)

// 			if tt.expectedStatus != http.StatusCreated {
// 				var response httpresponse.ErrorResponse
// 				_ = json.NewDecoder(rec.Body).Decode(&response)
// 				fmt.Println(response)
// 				assert.Equal(t, tt.expectedBody.Message, response.Message)
// 			}
// 		})
// 	}
// }
