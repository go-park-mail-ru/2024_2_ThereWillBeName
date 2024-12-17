package http

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"io"
// 	"log/slog"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"

// 	"2024_2_ThereWillBeName/internal/pkg/middleware"
// 	"2024_2_ThereWillBeName/internal/pkg/trips/delivery/grpc/gen"
// 	mock "2024_2_ThereWillBeName/internal/pkg/trips/mocks"

// 	"github.com/golang/mock/gomock"
// 	"github.com/gorilla/mux"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateTripHandler(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockClient := mock.NewMockTripsClient(ctrl)
// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}
// 	handl := slog.NewJSONHandler(os.Stdout, opts)
// 	logger := slog.New(handl)

// 	handler := NewTripHandler(mockClient, logger)

// 	tests := []struct {
// 		name         string
// 		requestBody  interface{}
// 		mockBehavior func()
// 		expectedCode int
// 		expectedBody string
// 	}{
// 		{
// 			name: "Success",
// 			requestBody: TripData{
// 				UserID:      1,
// 				Name:        "Trip to Paris",
// 				CityID:      2,
// 				Description: "A nice trip",
// 				StartDate:   "2024-12-01",
// 				EndDate:     "2024-12-10",
// 				Private:     false,
// 			},
// 			mockBehavior: func() {
// 				mockClient.EXPECT().CreateTrip(gomock.Any(), gomock.Any()).Return(&gen.EmptyResponse{}, nil)
// 			},
// 			expectedCode: http.StatusCreated,
// 			expectedBody: `"Trip created successfully"`,
// 		},
// 		{
// 			name:         "Invalid JSON",
// 			requestBody:  `invalid-json`,
// 			mockBehavior: func() {},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: `{"message":"Invalid request"}`,
// 		},
// 		{
// 			name: "Failed Usecase",
// 			requestBody: TripData{
// 				UserID:      1,
// 				Name:        "Trip to Paris",
// 				CityID:      2,
// 				Description: "A nice trip",
// 				StartDate:   "2024-12-01",
// 				EndDate:     "2024-12-10",
// 				Private:     false,
// 			},
// 			mockBehavior: func() {
// 				mockClient.EXPECT().CreateTrip(gomock.Any(), gomock.Any()).Return(nil, errors.New("usecase error"))
// 			},
// 			expectedCode: http.StatusInternalServerError,
// 			expectedBody: `{"message":"Failed to create trip"}`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			body, _ := json.Marshal(tt.requestBody)
// 			req := httptest.NewRequest(http.MethodPost, "/trips", bytes.NewReader(body))
// 			req = req.WithContext(context.WithValue(req.Context(), middleware.IdKey, uint(1)))

// 			w := httptest.NewRecorder()

// 			tt.mockBehavior()

// 			handler.CreateTripHandler(w, req)

// 			resp := w.Result()
// 			assert.Equal(t, tt.expectedCode, resp.StatusCode)

// 			respBody, _ := io.ReadAll(resp.Body)
// 			assert.Contains(t, string(respBody), tt.expectedBody)
// 		})
// 	}
// }
// func TestUpdateTripHandler(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockClient := mock.NewMockTripsClient(ctrl)
// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}
// 	handl := slog.NewJSONHandler(os.Stdout, opts)
// 	logger := slog.New(handl)

// 	handler := NewTripHandler(mockClient, logger)

// 	router := mux.NewRouter()
// 	router.HandleFunc("/trips/{id}", handler.UpdateTripHandler).Methods(http.MethodPut)

// 	tests := []struct {
// 		name         string
// 		requestPath  string
// 		requestBody  interface{}
// 		mockBehavior func()
// 		expectedCode int
// 		expectedBody string
// 	}{
// 		{
// 			name:        "Success",
// 			requestPath: "/trips/123",
// 			requestBody: TripData{
// 				UserID:      1,
// 				Name:        "Updated Trip",
// 				CityID:      2,
// 				Description: "Updated description",
// 				StartDate:   "2024-12-01",
// 				EndDate:     "2024-12-10",
// 				Private:     false,
// 			},
// 			mockBehavior: func() {
// 				mockClient.EXPECT().UpdateTrip(gomock.Any(), gomock.Any()).Return(&gen.EmptyResponse{}, nil)
// 			},
// 			expectedCode: http.StatusOK,
// 			expectedBody: `"Trip updated successfully"`,
// 		},
// 		{
// 			name:        "Invalid ID",
// 			requestPath: "/trips/abc",
// 			requestBody: TripData{
// 				UserID:      1,
// 				Name:        "Updated Trip",
// 				CityID:      2,
// 				Description: "Updated description",
// 				StartDate:   "2024-12-01",
// 				EndDate:     "2024-12-10",
// 				Private:     false,
// 			},
// 			mockBehavior: func() {},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: `{"message":"Invalid trip ID"}`,
// 		},
// 		{
// 			name:         "Invalid Body",
// 			requestPath:  "/trips/123",
// 			requestBody:  `invalid-json`,
// 			mockBehavior: func() {},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: `{"message":"Invalid trip data"}`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var body []byte
// 			if b, ok := tt.requestBody.([]byte); ok {
// 				body = b
// 			} else {
// 				body, _ = json.Marshal(tt.requestBody)
// 			}

// 			req := httptest.NewRequest(http.MethodPut, tt.requestPath, bytes.NewReader(body))
// 			req = req.WithContext(context.WithValue(req.Context(), middleware.IdKey, uint(1)))

// 			w := httptest.NewRecorder()

// 			tt.mockBehavior()

// 			router.ServeHTTP(w, req)

// 			resp := w.Result()
// 			assert.Equal(t, tt.expectedCode, resp.StatusCode)

// 			respBody, _ := io.ReadAll(resp.Body)
// 			assert.Contains(t, string(respBody), tt.expectedBody)
// 		})
// 	}
// }

// func TestDeleteTripHandler(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockClient := mock.NewMockTripsClient(ctrl)
// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}
// 	handl := slog.NewJSONHandler(os.Stdout, opts)
// 	logger := slog.New(handl)

// 	handler := NewTripHandler(mockClient, logger)

// 	router := mux.NewRouter()
// 	router.HandleFunc("/trips/{id}", handler.DeleteTripHandler).Methods(http.MethodDelete)

// 	tests := []struct {
// 		name         string
// 		tripID       string
// 		authUserID   uint
// 		mockBehavior func()
// 		expectedCode int
// 		expectedBody string
// 	}{
// 		{
// 			name:       "Success",
// 			tripID:     "123",
// 			authUserID: 1,
// 			mockBehavior: func() {
// 				mockClient.EXPECT().DeleteTrip(gomock.Any(), &gen.DeleteTripRequest{Id: 123}).Return(&gen.EmptyResponse{}, nil)
// 			},
// 			expectedCode: http.StatusNoContent,
// 			expectedBody: `"Trip deleted successfully"`,
// 		},
// 		{
// 			name:         "Invalid Trip ID",
// 			tripID:       "invalid-id",
// 			authUserID:   1,
// 			mockBehavior: func() {},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: `{"message":"Invalid trip ID"}`,
// 		},
// 		{
// 			name:         "Unauthorized",
// 			tripID:       "123",
// 			authUserID:   0,
// 			mockBehavior: func() {},
// 			expectedCode: http.StatusUnauthorized,
// 			expectedBody: `{"message":"User is not authorized"}`,
// 		},
// 		{
// 			name:       "GRPC Error",
// 			tripID:     "123",
// 			authUserID: 1,
// 			mockBehavior: func() {
// 				mockClient.EXPECT().DeleteTrip(gomock.Any(), &gen.DeleteTripRequest{Id: 123}).Return(nil, errors.New("grpc error"))
// 			},
// 			expectedCode: http.StatusInternalServerError,
// 			expectedBody: `{"message":"Failed to delete trip"}`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			req := httptest.NewRequest(http.MethodDelete, "/trips/"+tt.tripID, nil)

// 			if tt.authUserID > 0 {
// 				req = req.WithContext(context.WithValue(req.Context(), middleware.IdKey, tt.authUserID))
// 			}

// 			w := httptest.NewRecorder()

// 			tt.mockBehavior()

// 			router.ServeHTTP(w, req)

// 			resp := w.Result()
// 			assert.Equal(t, tt.expectedCode, resp.StatusCode)

// 			respBody, _ := io.ReadAll(resp.Body)
// 			assert.Contains(t, string(respBody), tt.expectedBody)
// 		})
// 	}
// }

// func TestGetTripsByUserIDHandler(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockClient := mock.NewMockTripsClient(ctrl)
// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}
// 	handl := slog.NewJSONHandler(os.Stdout, opts)
// 	logger := slog.New(handl)
// 	handler := NewTripHandler(mockClient, logger)

// 	tests := []struct {
// 		name         string
// 		authUserID   uint
// 		page         string
// 		mockBehavior func()
// 		expectedCode int
// 		expectedBody string
// 	}{
// 		{
// 			name:       "Success",
// 			authUserID: 1,
// 			page:       "1",
// 			mockBehavior: func() {
// 				mockClient.EXPECT().GetTripsByUserID(gomock.Any(), &gen.GetTripsByUserIDRequest{
// 					UserId: 1,
// 					Limit:  10,
// 					Offset: 0,
// 				}).Return(&gen.GetTripsByUserIDResponse{
// 					Trips: []*gen.Trip{
// 						{Id: 1, UserId: 1, Name: "Trip 1", Description: "Description 1"},
// 					},
// 				}, nil)
// 			},
// 			expectedCode: http.StatusOK,
// 			expectedBody: `[{"id":1,"user_id":1,"name":"Trip 1","description":"Description 1"}]`,
// 		},
// 		{
// 			name:         "Unauthorized",
// 			authUserID:   0,
// 			page:         "1",
// 			mockBehavior: func() {},
// 			expectedCode: http.StatusUnauthorized,
// 			expectedBody: `{"message":"User is not authorized"}`,
// 		},
// 		{
// 			name:         "Invalid Page Number",
// 			authUserID:   1,
// 			page:         "invalid",
// 			mockBehavior: func() {},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: `{"message":"Invalid page number"}`,
// 		},
// 		{
// 			name:       "GRPC Error",
// 			authUserID: 1,
// 			page:       "1",
// 			mockBehavior: func() {
// 				mockClient.EXPECT().GetTripsByUserID(gomock.Any(), gomock.Any()).Return(nil, errors.New("grpc error"))
// 			},
// 			expectedCode: http.StatusInternalServerError,
// 			expectedBody: `{"message":"Failed to retrieve trip"}`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			req := httptest.NewRequest(http.MethodGet, "/trips?page="+tt.page, nil)

// 			if tt.authUserID > 0 {
// 				req = req.WithContext(context.WithValue(req.Context(), middleware.IdKey, tt.authUserID))
// 			}

// 			w := httptest.NewRecorder()

// 			tt.mockBehavior()

// 			handler.GetTripsByUserIDHandler(w, req)

// 			resp := w.Result()
// 			assert.Equal(t, tt.expectedCode, resp.StatusCode)

// 			respBody, _ := io.ReadAll(resp.Body)
// 			assert.JSONEq(t, tt.expectedBody, string(respBody), "Response body does not match")
// 		})
// 	}
// }

// func TestGetTripHandler(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockClient := mock.NewMockTripsClient(ctrl)
// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}
// 	handl := slog.NewJSONHandler(os.Stdout, opts)
// 	logger := slog.New(handl)
// 	handler := NewTripHandler(mockClient, logger)

// 	tests := []struct {
// 		name         string
// 		tripID       string
// 		authUserID   uint
// 		mockBehavior func()
// 		expectedCode int
// 		expectedBody map[string]interface{}
// 	}{
// 		{
// 			name:       "Success",
// 			tripID:     "1",
// 			authUserID: 1,
// 			mockBehavior: func() {
// 				mockClient.EXPECT().GetTrip(gomock.Any(), &gen.GetTripRequest{TripId: 1}).Return(&gen.GetTripResponse{
// 					Trip: &gen.Trip{
// 						Id:          1,
// 						UserId:      1,
// 						Name:        "Trip to Paris",
// 						Description: "A nice trip",
// 						CityId:      2,
// 						StartDate:   "2024-12-01",
// 						EndDate:     "2024-12-10",
// 						Private:     false,
// 					},
// 				}, nil)
// 			},
// 			expectedCode: http.StatusOK,
// 			expectedBody: map[string]interface{}{
// 				"id":          float64(1),
// 				"user_id":     float64(1),
// 				"name":        "Trip to Paris",
// 				"description": "A nice trip",
// 				"city_id":     float64(2),
// 				"start_date":  "2024-12-01",
// 				"end_date":    "2024-12-10",
// 			},
// 		},

// 		{
// 			name:       "Invalid Trip ID",
// 			tripID:     "abc",
// 			authUserID: 1,
// 			mockBehavior: func() {
// 			},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: map[string]interface{}{
// 				"message": "Invalid trip ID",
// 			},
// 		},
// 		{
// 			name:       "Trip Not Found",
// 			tripID:     "1",
// 			authUserID: 1,
// 			mockBehavior: func() {
// 				mockClient.EXPECT().GetTrip(gomock.Any(), &gen.GetTripRequest{TripId: 1}).Return(nil, errors.New("not found"))
// 			},
// 			expectedCode: http.StatusInternalServerError,
// 			expectedBody: map[string]interface{}{
// 				"message": "Failed to retrieve trip",
// 			},
// 		},
// 		{
// 			name:       "Internal Server Error",
// 			tripID:     "1",
// 			authUserID: 1,
// 			mockBehavior: func() {
// 				mockClient.EXPECT().GetTrip(gomock.Any(), &gen.GetTripRequest{TripId: 1}).Return(nil, errors.New("internal error"))
// 			},
// 			expectedCode: http.StatusInternalServerError,
// 			expectedBody: map[string]interface{}{
// 				"message": "Failed to retrieve trip",
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()
// 			tt.mockBehavior()

// 			router := mux.NewRouter()
// 			router.HandleFunc("/trips/{id}", func(w http.ResponseWriter, r *http.Request) {
// 				mux.SetURLVars(r, map[string]string{"id": tt.tripID})
// 				handler.GetTripHandler(w, r)
// 			})

// 			req := httptest.NewRequest(http.MethodGet, "/trips/"+tt.tripID, nil)
// 			req = req.WithContext(context.WithValue(req.Context(), middleware.IdKey, tt.authUserID))

// 			w := httptest.NewRecorder()

// 			router.ServeHTTP(w, req)

// 			resp := w.Result()
// 			assert.Equal(t, tt.expectedCode, resp.StatusCode)

// 			var responseBody map[string]interface{}
// 			json.NewDecoder(resp.Body).Decode(&responseBody)

// 			assert.Equal(t, tt.expectedBody, responseBody)
// 		})
// 	}
// }

// func TestAddPlaceToTripHandler(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockClient := mock.NewMockTripsClient(ctrl)
// 	opts := &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	}
// 	handl := slog.NewJSONHandler(os.Stdout, opts)
// 	logger := slog.New(handl)

// 	handler := NewTripHandler(mockClient, logger)

// 	tests := []struct {
// 		name         string
// 		tripID       string
// 		authUserID   uint
// 		requestBody  interface{}
// 		mockBehavior func()
// 		expectedCode int
// 		expectedBody string
// 	}{
// 		{
// 			name:       "Success",
// 			tripID:     "1",
// 			authUserID: 1,
// 			requestBody: AddPlaceRequest{
// 				PlaceID: 123,
// 			},
// 			mockBehavior: func() {
// 				mockClient.EXPECT().AddPlaceToTrip(gomock.Any(), &gen.AddPlaceToTripRequest{
// 					TripId:  1,
// 					PlaceId: 123,
// 				}).Return(&gen.EmptyResponse{}, nil)
// 			},
// 			expectedCode: http.StatusCreated,
// 			expectedBody: `"Place added to trip successfully"`,
// 		},
// 		{
// 			name:         "Unauthorized",
// 			tripID:       "1",
// 			authUserID:   0,
// 			requestBody:  AddPlaceRequest{PlaceID: 123},
// 			mockBehavior: func() {},
// 			expectedCode: http.StatusUnauthorized,
// 			expectedBody: `{"message":"User is not authorized"}`,
// 		},
// 		{
// 			name:         "Invalid Trip ID Format",
// 			tripID:       "invalid",
// 			authUserID:   1,
// 			requestBody:  AddPlaceRequest{PlaceID: 123},
// 			mockBehavior: func() {},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: `{"message":"Invalid trip ID"}`,
// 		},
// 		{
// 			name:         "Invalid Request Body",
// 			tripID:       "1",
// 			authUserID:   1,
// 			requestBody:  "invalid-json",
// 			mockBehavior: func() {},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: `{"message":"Invalid place ID"}`,
// 		},
// 		{
// 			name:       "GRPC Error",
// 			tripID:     "1",
// 			authUserID: 1,
// 			requestBody: AddPlaceRequest{
// 				PlaceID: 123,
// 			},
// 			mockBehavior: func() {
// 				mockClient.EXPECT().AddPlaceToTrip(gomock.Any(), &gen.AddPlaceToTripRequest{
// 					TripId:  1,
// 					PlaceId: 123,
// 				}).Return(nil, errors.New("grpc error"))
// 			},
// 			expectedCode: http.StatusInternalServerError,
// 			expectedBody: `{"message":"Failed to add place trip"}`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			router := mux.NewRouter()
// 			router.HandleFunc("/trips/{id}/places", handler.AddPlaceToTripHandler).Methods(http.MethodPost)

// 			var body []byte
// 			if b, ok := tt.requestBody.([]byte); ok {
// 				body = b
// 			} else {
// 				body, _ = json.Marshal(tt.requestBody)
// 			}

// 			req := httptest.NewRequest(http.MethodPost, "/trips/"+tt.tripID+"/places", bytes.NewReader(body))
// 			if tt.authUserID > 0 {
// 				req = req.WithContext(context.WithValue(req.Context(), middleware.IdKey, tt.authUserID))
// 			}

// 			w := httptest.NewRecorder()

// 			tt.mockBehavior()

// 			router.ServeHTTP(w, req)

// 			resp := w.Result()
// 			assert.Equal(t, tt.expectedCode, resp.StatusCode)

// 			respBody, _ := io.ReadAll(resp.Body)
// 			assert.Contains(t, string(respBody), tt.expectedBody)
// 		})
// 	}
// }

// func TestAddPhotosToTripHandler(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockClient := mock.NewMockTripsClient(ctrl)
// 	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
// 	handler := NewTripHandler(mockClient, logger)

// 	tests := []struct {
// 		name         string
// 		tripID       string
// 		authUserID   uint
// 		requestBody  interface{}
// 		mockBehavior func()
// 		expectedCode int
// 		expectedBody string
// 	}{
// 		{
// 			name:       "Success",
// 			tripID:     "1",
// 			authUserID: 1,
// 			requestBody: map[string][]string{
// 				"photos": {"photo1_base64", "photo2_base64"},
// 			},
// 			mockBehavior: func() {
// 				mockClient.EXPECT().AddPhotosToTrip(gomock.Any(), &gen.AddPhotosToTripRequest{
// 					TripId: 1,
// 					Photos: []string{"photo1_base64", "photo2_base64"},
// 				}).Return(&gen.AddPhotosToTripResponse{
// 					Photos: []*gen.Photo{
// 						{PhotoPath: "photo1_path"},
// 						{PhotoPath: "photo2_path"},
// 					},
// 				}, nil)
// 			},
// 			expectedCode: http.StatusCreated,
// 			expectedBody: `[{"photoPath":"photo1_path"},{"photoPath":"photo2_path"}]`,
// 		},
// 		{
// 			name:         "Invalid Trip ID",
// 			tripID:       "abc",
// 			authUserID:   1,
// 			requestBody:  map[string][]string{"photos": {"photo1_base64"}},
// 			mockBehavior: func() {},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: `{"message":"Invalid trip ID"}`,
// 		},
// 		{
// 			name:         "Invalid Request Body",
// 			tripID:       "1",
// 			authUserID:   1,
// 			requestBody:  `invalid-body`,
// 			mockBehavior: func() {},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: `{"message":"Invalid request body"}`,
// 		},
// 		{
// 			name:       "GRPC Error",
// 			tripID:     "1",
// 			authUserID: 1,
// 			requestBody: map[string][]string{
// 				"photos": {"photo1_base64"},
// 			},
// 			mockBehavior: func() {
// 				mockClient.EXPECT().AddPhotosToTrip(gomock.Any(), gomock.Any()).Return(nil, errors.New("grpc error"))
// 			},
// 			expectedCode: http.StatusInternalServerError,
// 			expectedBody: `{"message":"Failed to add photos trip"}`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			router := mux.NewRouter()
// 			router.HandleFunc("/trips/{id}/photos", handler.AddPhotosToTripHandler).Methods(http.MethodPost)

// 			var body []byte
// 			if b, ok := tt.requestBody.([]byte); ok {
// 				body = b
// 			} else {
// 				body, _ = json.Marshal(tt.requestBody)
// 			}

// 			req := httptest.NewRequest(http.MethodPost, "/trips/"+tt.tripID+"/photos", bytes.NewReader(body))
// 			req = req.WithContext(context.WithValue(req.Context(), middleware.IdKey, tt.authUserID))

// 			w := httptest.NewRecorder()
// 			tt.mockBehavior()
// 			router.ServeHTTP(w, req)

// 			resp := w.Result()
// 			assert.Equal(t, tt.expectedCode, resp.StatusCode)

// 			respBody, _ := io.ReadAll(resp.Body)
// 			assert.JSONEq(t, tt.expectedBody, string(respBody))
// 		})
// 	}
// }

// func TestDeletePhotoHandler(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockClient := mock.NewMockTripsClient(ctrl)
// 	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
// 	handler := NewTripHandler(mockClient, logger)

// 	tests := []struct {
// 		name         string
// 		tripID       string
// 		authUserID   uint
// 		requestBody  interface{}
// 		mockBehavior func()
// 		expectedCode int
// 		expectedBody string
// 	}{
// 		{
// 			name:       "Success",
// 			tripID:     "1",
// 			authUserID: 1,
// 			requestBody: map[string]string{
// 				"photo_path": "photo_path_1",
// 			},
// 			mockBehavior: func() {
// 				mockClient.EXPECT().DeletePhotoFromTrip(gomock.Any(), &gen.DeletePhotoRequest{
// 					TripId:    1,
// 					PhotoPath: "photo_path_1",
// 				}).Return(&gen.EmptyResponse{}, nil)
// 			},
// 			expectedCode: http.StatusOK,
// 			expectedBody: `{"message":"Photo deleted successfully"}`,
// 		},
// 		{
// 			name:         "Invalid Trip ID",
// 			tripID:       "abc",
// 			authUserID:   1,
// 			requestBody:  map[string]string{"photo_path": "photo_path_1"},
// 			mockBehavior: func() {},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: `{"message":"Invalid trip ID"}`,
// 		},
// 		{
// 			name:         "Invalid Request Body",
// 			tripID:       "1",
// 			authUserID:   1,
// 			requestBody:  `invalid-body`,
// 			mockBehavior: func() {},
// 			expectedCode: http.StatusBadRequest,
// 			expectedBody: `{"message":"Invalid request body"}`,
// 		},
// 		{
// 			name:       "GRPC Error",
// 			tripID:     "1",
// 			authUserID: 1,
// 			requestBody: map[string]string{
// 				"photo_path": "photo_path_1",
// 			},
// 			mockBehavior: func() {
// 				mockClient.EXPECT().DeletePhotoFromTrip(gomock.Any(), gomock.Any()).Return(nil, errors.New("grpc error"))
// 			},
// 			expectedCode: http.StatusInternalServerError,
// 			expectedBody: `{"message":"Failed to delete photo"}`,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			router := mux.NewRouter()
// 			router.HandleFunc("/trips/{id}/photos", handler.DeletePhotoHandler).Methods(http.MethodDelete)

// 			var body []byte
// 			if b, ok := tt.requestBody.([]byte); ok {
// 				body = b
// 			} else {
// 				body, _ = json.Marshal(tt.requestBody)
// 			}

// 			req := httptest.NewRequest(http.MethodDelete, "/trips/"+tt.tripID+"/photos", bytes.NewReader(body))
// 			req = req.WithContext(context.WithValue(req.Context(), middleware.IdKey, tt.authUserID))

// 			w := httptest.NewRecorder()
// 			tt.mockBehavior()
// 			router.ServeHTTP(w, req)

// 			resp := w.Result()
// 			assert.Equal(t, tt.expectedCode, resp.StatusCode)

// 			respBody, _ := io.ReadAll(resp.Body)
// 			assert.JSONEq(t, tt.expectedBody, string(respBody))
// 		})
// 	}
// }
