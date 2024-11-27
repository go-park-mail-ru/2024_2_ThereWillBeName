package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"bytes"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateTrip(t *testing.T) {
	tests := []struct {
		name          string
		trip          models.Trip
		mockBehavior  func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "Success",
			trip: models.Trip{
				UserID:      1,
				Name:        "Trip to Paris",
				Description: "A great trip",
				CityID:      2,
				StartDate:   "2024-12-01",
				EndDate:     "2024-12-10",
				Private:     false,
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO trip`).
					WithArgs(1, "Trip to Paris", "A great trip", 2, "2024-12-01", "2024-12-10", false).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Insert Failed",
			trip: models.Trip{
				UserID:      1,
				Name:        "Trip to Paris",
				Description: "A great trip",
				CityID:      2,
				StartDate:   "2024-12-01",
				EndDate:     "2024-12-10",
				Private:     false,
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO trip`).
					WithArgs(1, "Trip to Paris", "A great trip", 2, "2024-12-01", "2024-12-10", false).
					WillReturnError(errors.New("insert failed"))
			},
			expectedError: models.ErrInternal,
		},
		{
			name: "No Rows Affected",
			trip: models.Trip{
				UserID:      1,
				Name:        "Trip to Paris",
				Description: "A great trip",
				CityID:      2,
				StartDate:   "2024-12-01",
				EndDate:     "2024-12-10",
				Private:     false,
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO trip`).
					WithArgs(1, "Trip to Paris", "A great trip", 2, "2024-12-01", "2024-12-10", false).
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: models.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()
			var logBuffer bytes.Buffer

			handler := slog.NewTextHandler(&logBuffer, nil)

			logger := slog.New(handler)
			loggerDB := dblogger.NewDB(db, logger)

			repo := NewTripRepository(loggerDB)

			tt.mockBehavior(mock)

			err := repo.CreateTrip(context.Background(), tt.trip)
			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateTrip(t *testing.T) {
	tests := []struct {
		name          string
		trip          models.Trip
		mockBehavior  func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "Success",
			trip: models.Trip{
				ID:          1,
				Name:        "Updated Trip",
				Description: "Updated Description",
				CityID:      2,
				StartDate:   "2024-12-01",
				EndDate:     "2024-12-10",
				Private:     false,
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE trip`).
					WithArgs("Updated Trip", "Updated Description", 2, "2024-12-01", "2024-12-10", false, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Update Failed",
			trip: models.Trip{
				ID:          1,
				Name:        "Updated Trip",
				Description: "Updated Description",
				CityID:      2,
				StartDate:   "2024-12-01",
				EndDate:     "2024-12-10",
				Private:     false,
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE trip`).
					WithArgs("Updated Trip", "Updated Description", 2, "2024-12-01", "2024-12-10", false, 1).
					WillReturnError(errors.New("update failed"))
			},
			expectedError: models.ErrInternal,
		},
		{
			name: "No Rows Affected",
			trip: models.Trip{
				ID:          1,
				Name:        "Updated Trip",
				Description: "Updated Description",
				CityID:      2,
				StartDate:   "2024-12-01",
				EndDate:     "2024-12-10",
				Private:     false,
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE trip`).
					WithArgs("Updated Trip", "Updated Description", 2, "2024-12-01", "2024-12-10", false, 1).
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: models.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			var logBuffer bytes.Buffer

			handler := slog.NewTextHandler(&logBuffer, nil)

			logger := slog.New(handler)
			loggerDB := dblogger.NewDB(db, logger)

			repo := NewTripRepository(loggerDB)

			tt.mockBehavior(mock)

			err := repo.UpdateTrip(context.Background(), tt.trip)
			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteTrip(t *testing.T) {
	tests := []struct {
		name          string
		tripID        uint
		mockBehavior  func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name:   "Success",
			tripID: 1,
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM trip`).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name:   "Delete Failed",
			tripID: 1,
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM trip`).
					WithArgs(1).
					WillReturnError(errors.New("delete failed"))
			},
			expectedError: models.ErrInternal,
		},
		{
			name:   "No Rows Affected",
			tripID: 1,
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM trip`).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: models.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()
			var logBuffer bytes.Buffer

			handler := slog.NewTextHandler(&logBuffer, nil)

			logger := slog.New(handler)
			loggerDB := dblogger.NewDB(db, logger)

			repo := NewTripRepository(loggerDB)

			tt.mockBehavior(mock)

			err := repo.DeleteTrip(context.Background(), tt.tripID)
			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
func TestGetTripsByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening stub database connection: %v", err)
	}
	defer db.Close()

	var logBuffer bytes.Buffer

	handler := slog.NewTextHandler(&logBuffer, nil)

	logger := slog.New(handler)
	loggerDB := dblogger.NewDB(db, logger)

	repo := NewTripRepository(loggerDB)

	t.Run("Success", func(t *testing.T) {
		createdAt := time.Date(2024, time.November, 25, 0, 0, 0, 0, time.UTC)
		rows := sqlmock.NewRows([]string{
			"id", "user_id", "name", "description", "city_id",
			"start_date", "end_date", "private", "created_at", "photos",
		}).AddRow(
			1, 1, "Trip 1", "Description 1", 2,
			"2024-12-01", "2024-12-10", false, createdAt, pq.Array([]string{"photo1.jpg", "photo2.jpg"}),
		)

		mock.ExpectQuery(`SELECT (.+) FROM trip t LEFT JOIN trip_photo tp`).
			WithArgs(1, 10, 0).
			WillReturnRows(rows)

		expected := []models.Trip{
			{
				ID:          1,
				UserID:      1,
				Name:        "Trip 1",
				Description: "Description 1",
				CityID:      2,
				StartDate:   "2024-12-01",
				EndDate:     "2024-12-10",
				Private:     false,
				Photos:      []string{"photo1.jpg", "photo2.jpg"},
				CreatedAt:   createdAt,
			},
		}

		result, err := repo.GetTripsByUserID(context.Background(), 1, 10, 0)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("No Trips Found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "user_id", "name", "description", "city_id",
			"start_date", "end_date", "private", "created_at", "photos",
		})

		mock.ExpectQuery(`SELECT (.+) FROM trip t LEFT JOIN trip_photo tp`).
			WithArgs(1, 10, 0).
			WillReturnRows(rows)

		result, err := repo.GetTripsByUserID(context.Background(), 1, 10, 0)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no trips found")
		assert.Nil(t, result)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT (.+) FROM trip t LEFT JOIN trip_photo tp`).
			WithArgs(1, 10, 0).
			WillReturnError(fmt.Errorf("database error"))

		result, err := repo.GetTripsByUserID(context.Background(), 1, 10, 0)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to retrieve trips")
		assert.Nil(t, result)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestAddPlaceToTrip(t *testing.T) {
	tests := []struct {
		name          string
		tripID        uint
		placeID       uint
		mockBehavior  func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name:    "Success",
			tripID:  1,
			placeID: 2,
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO trip_place`).
					WithArgs(1, 2).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name:    "Insert Failed",
			tripID:  1,
			placeID: 2,
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO trip_place`).
					WithArgs(1, 2).
					WillReturnError(errors.New("insert failed"))
			},
			expectedError: models.ErrInternal,
		},
		{
			name:    "No Rows Affected",
			tripID:  1,
			placeID: 2,
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO trip_place`).
					WithArgs(1, 2).
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: models.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			var logBuffer bytes.Buffer

			handler := slog.NewTextHandler(&logBuffer, nil)

			logger := slog.New(handler)
			loggerDB := dblogger.NewDB(db, logger)

			repo := NewTripRepository(loggerDB)
			tt.mockBehavior(mock)

			err := repo.AddPlaceToTrip(context.Background(), tt.tripID, tt.placeID)
			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
func TestAddPhotoToTrip(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening stub database connection: %v", err)
	}
	defer db.Close()

	var logBuffer bytes.Buffer

	handler := slog.NewTextHandler(&logBuffer, nil)

	logger := slog.New(handler)
	loggerDB := dblogger.NewDB(db, logger)

	repo := NewTripRepository(loggerDB)
	t.Run("Success", func(t *testing.T) {
		tripID := uint(1)
		photoPath := "photo1.jpg"

		query := `
        INSERT INTO trip_photo \(trip_id, photo_path\)
        VALUES \(\$1, \$2\)`

		mock.ExpectExec(query).
			WithArgs(tripID, photoPath).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.AddPhotoToTrip(context.Background(), tripID, photoPath)
		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		tripID := uint(1)
		photoPath := "photo1.jpg"

		query := `
        INSERT INTO trip_photo \(trip_id, photo_path\)
        VALUES \(\$1, \$2\)`

		mock.ExpectExec(query).
			WithArgs(tripID, photoPath).
			WillReturnError(fmt.Errorf("database error"))

		err := repo.AddPhotoToTrip(context.Background(), tripID, photoPath)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to insert photo into database")

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestDeletePhotoFromTrip(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening stub database connection: %v", err)
	}
	defer db.Close()

	var logBuffer bytes.Buffer

	handler := slog.NewTextHandler(&logBuffer, nil)

	logger := slog.New(handler)
	loggerDB := dblogger.NewDB(db, logger)

	repo := NewTripRepository(loggerDB)
	t.Run("Success", func(t *testing.T) {
		tripID := uint(1)
		photoPath := "photo1.jpg"

		query := `DELETE FROM trip_photo WHERE trip_id = \$1 AND photo_path = \$2`

		mock.ExpectExec(query).
			WithArgs(tripID, photoPath).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.DeletePhotoFromTrip(context.Background(), tripID, photoPath)
		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Photo Not Found", func(t *testing.T) {
		tripID := uint(1)
		photoPath := "photo1.jpg"

		query := `DELETE FROM trip_photo WHERE trip_id = \$1 AND photo_path = \$2`

		mock.ExpectExec(query).
			WithArgs(tripID, photoPath).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.DeletePhotoFromTrip(context.Background(), tripID, photoPath)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "photo not found in trip")

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		tripID := uint(1)
		photoPath := "photo1.jpg"

		query := `DELETE FROM trip_photo WHERE trip_id = \$1 AND photo_path = \$2`

		mock.ExpectExec(query).
			WithArgs(tripID, photoPath).
			WillReturnError(fmt.Errorf("database error"))

		err := repo.DeletePhotoFromTrip(context.Background(), tripID, photoPath)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to delete photo from database")

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
