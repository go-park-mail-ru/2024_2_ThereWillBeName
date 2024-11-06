package http

import (
	"2024_2_ThereWillBeName/internal/models"
	mockplaces "2024_2_ThereWillBeName/internal/pkg/places/mocks"
	"bytes"
	"context"
	"encoding/json"
	"log"

	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const errorParse = "ErrorParse"

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

	h := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.New(h)

	mockUsecase := mockplaces.NewMockPlaceUsecase(ctrl)
	handler := NewPlacesHandler(mockUsecase, logger)

	tests := []struct {
		name         string
		offset       string
		limit        string
		mockPlaces   []models.GetPlace
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid request",
			offset:       "0",
			limit:        "10",
			mockPlaces:   places,
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: stringPlaces + "\n",
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
				mockUsecase.EXPECT().GetPlaces(gomock.Any(), limit, offset).Return(tt.mockPlaces, tt.mockError)
			}

			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/places", nil)
			assert.NoError(t, err)

			query := url.Values{}
			query.Add("offset", tt.offset)
			query.Add("limit", tt.limit)
			req.URL.RawQuery = query.Encode()

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/places", handler.GetPlacesHandler)
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedCode, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())
		})
	}
}

func TestPostPlacesHandler(t *testing.T) {
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

	jsonPlace, _ := json.Marshal(place)
	//reader := bytes.NewReader(jsonPlace)
	//stringPlace := string(jsonPlace)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.New(h)

	mockUsecase := mockplaces.NewMockPlaceUsecase(ctrl)
	handler := NewPlacesHandler(mockUsecase, logger)

	tests := []struct {
		name         string
		mockReturn   []byte
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Success",
			mockReturn:   jsonPlace,
			mockError:    nil,
			expectedCode: http.StatusCreated,
			expectedBody: "\"Place succesfully created\"\n",
		},
		{
			name:         "ErrorINternal",
			mockReturn:   jsonPlace,
			mockError:    errors.New("dcd"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: "",
		},
		{
			name:         errorParse,
			mockReturn:   []byte(`{ "name": "Test Place", "imagePath": "/path/to/image", "description": "Test Description", "rating": 4.5, "numberOfReviews": 10, "address": "Test Address", "cityId": 1, "phoneNumber": "1234567890", "categoriesId": [1, 2, 3]`),
			mockError:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: "",
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			if testcase.name != errorParse {
				mockUsecase.EXPECT().CreatePlace(context.Background(), gomock.Any()).Return(testcase.mockError)
			}

			req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, "/places", bytes.NewBuffer(testcase.mockReturn))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.PostPlaceHandler(rr, req)

			assert.Equal(t, testcase.expectedCode, rr.Code)
			assert.Equal(t, testcase.expectedBody, rr.Body.String())
		})
	}
}

func TestPutPlacesHandler(t *testing.T) {
	place := models.UpdatePlace{
		ID:              1,
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

	jsonPlace, _ := json.Marshal(place)
	//reader := bytes.NewReader(jsonPlace)
	//stringPlace := string(jsonPlace)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.New(h)
	mockUsecase := mockplaces.NewMockPlaceUsecase(ctrl)
	handler := NewPlacesHandler(mockUsecase, logger)

	tests := []struct {
		name         string
		mockReturn   []byte
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Success",
			mockReturn:   jsonPlace,
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: "\"place successfully updated\"\n",
		},
		{
			name:         "ErrorINternal",
			mockReturn:   jsonPlace,
			mockError:    errors.New("dcd"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: "",
		},
		{
			name:         errorParse,
			mockReturn:   []byte(`{ "name": "Test Place", "imagePath": "/path/to/image", "description": "Test Description", "rating": 4.5, "numberOfReviews": 10, "address": "Test Address", "cityId": 1, "phoneNumber": "1234567890", "categoriesId": [1, 2, 3]`),
			mockError:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: "",
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			if testcase.name != errorParse {
				mockUsecase.EXPECT().UpdatePlace(context.Background(), gomock.Any()).Return(testcase.mockError)
			}

			req, err := http.NewRequestWithContext(context.Background(), http.MethodPut, "/places/1", bytes.NewBuffer(testcase.mockReturn))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.PutPlaceHandler(rr, req)

			assert.Equal(t, testcase.expectedCode, rr.Code)
			assert.Equal(t, testcase.expectedBody, rr.Body.String())
		})
	}
}

func TestDeletePlacesHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.New(h)
	mockUsecase := mockplaces.NewMockPlaceUsecase(ctrl)
	handler := NewPlacesHandler(mockUsecase, logger)

	tests := []struct {
		name         string
		id           int
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Success",
			id:           1,
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: "\"place successfully deleted\"\n",
		},
		{
			name:         "ErrorINternal",
			id:           1,
			mockError:    errors.New("dcd"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: "",
		},
		{
			name:         errorParse,
			id:           -1,
			mockError:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: "",
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			if testcase.name != errorParse {
				mockUsecase.EXPECT().DeletePlace(gomock.Any(), uint(testcase.id)).Return(testcase.mockError)
			}

			req, err := http.NewRequestWithContext(context.Background(), http.MethodDelete, "/places/"+strconv.Itoa(int(testcase.id)), nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/places/{id}", handler.DeletePlaceHandler).Methods(http.MethodDelete)
			router.ServeHTTP(rr, req)

			assert.Equal(t, testcase.expectedCode, rr.Code)
			assert.Equal(t, testcase.expectedBody, rr.Body.String())
		})
	}
}

func TestGetPlaceHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.New(h)
	mockUsecase := mockplaces.NewMockPlaceUsecase(ctrl)
	handler := NewPlacesHandler(mockUsecase, logger)

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

	jsonPlaces, _ := json.Marshal(place)
	stringPlace := string(jsonPlaces)

	tests := []struct {
		name         string
		id           int
		mockPlace    models.GetPlace
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Success",
			id:           1,
			mockPlace:    place,
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: stringPlace + "\n",
		},
		{
			name:         "NotFound",
			id:           2,
			mockPlace:    models.GetPlace{},
			mockError:    models.ErrNotFound,
			expectedCode: http.StatusNotFound,
			expectedBody: `{"message":"place not found"}` + "\n",
		},
		{
			name:         "InternalServerError",
			id:           3,
			mockPlace:    models.GetPlace{},
			mockError:    errors.New("internal server error"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: ``,
		},
		{
			name:         "InvalidID",
			id:           -1,
			mockPlace:    models.GetPlace{},
			mockError:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: "",
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			if testcase.name != "InvalidID" {
				mockUsecase.EXPECT().GetPlace(gomock.Any(), uint(testcase.id)).Return(testcase.mockPlace, testcase.mockError)
			}

			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/places/"+strconv.Itoa(int(testcase.id)), nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/places/{id}", handler.GetPlaceHandler).Methods(http.MethodGet)
			router.ServeHTTP(rr, req)

			assert.Equal(t, testcase.expectedCode, rr.Code)
			assert.Equal(t, testcase.expectedBody, rr.Body.String())
		})
	}
}

func TestSearchPlaceHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	h := slog.NewJSONHandler(os.Stdout, nil)

	logger := slog.New(h)
	mockUsecase := mockplaces.NewMockPlaceUsecase(ctrl)
	handler := NewPlacesHandler(mockUsecase, logger)

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
	stringPlace := string(jsonPlaces)

	tests := []struct {
		name         string
		limit        string
		offset       string
		search       string
		mockPlaces   []models.GetPlace
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Success",
			search:       "test",
			limit:        "10",
			offset:       "0",
			mockPlaces:   places,
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectedBody: stringPlace + "\n",
		},
		{
			name:         "InternalServerError",
			search:       "badSearch",
			limit:        "10",
			offset:       "0",
			mockPlaces:   []models.GetPlace{},
			mockError:    errors.New("internal server error"),
			expectedCode: http.StatusInternalServerError,
			expectedBody: ``,
		},
		{
			name:         "InvalidOffset",
			search:       "test",
			offset:       "invalid",
			limit:        "10",
			mockPlaces:   []models.GetPlace{},
			mockError:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: ``,
		},
		{
			name:         "InvalidLimit",
			search:       "test",
			offset:       "0",
			limit:        "invalid",
			mockPlaces:   []models.GetPlace{},
			mockError:    nil,
			expectedCode: http.StatusBadRequest,
			expectedBody: ``,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			if testcase.name != "InvalidOffset" && testcase.name != "InvalidLimit" && testcase.name != "MissingPlaceName" {
				mockUsecase.EXPECT().SearchPlaces(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(testcase.mockPlaces, testcase.mockError)
			}
			urlStr := "/places/search/" + testcase.search
			log.Println(urlStr)
			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, urlStr, nil)
			assert.NoError(t, err)

			query := url.Values{}
			query.Add("offset", testcase.offset)
			query.Add("limit", testcase.limit)
			req.URL.RawQuery = query.Encode()

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/places/search/{search}", handler.SearchPlacesHandler).Methods(http.MethodGet)
			router.ServeHTTP(rr, req)

			assert.Equal(t, testcase.expectedCode, rr.Code)
			assert.Equal(t, testcase.expectedBody, rr.Body.String())
		})
	}
}
