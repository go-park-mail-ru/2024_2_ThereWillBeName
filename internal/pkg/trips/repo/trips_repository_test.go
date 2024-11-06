package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTrip(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	repo := NewTripRepository(db)

	tests := []struct {
		name        string
		trip        models.Trip
		mockSetup   func()
		expectedErr error
	}{
		{
			name: "successful creation",
			trip: models.Trip{UserID: 1, Name: "Test trip", Description: "A trip for testing", CityID: 1},
			mockSetup: func() {
				mock.ExpectExec(`INSERT INTO trips`).WithArgs(1, "Test trip", "A trip for testing", 1, sqlmock.AnyArg(), sqlmock.AnyArg(), false).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedErr: nil,
		},
		{
			name: "error on exec",
			trip: models.Trip{UserID: 1, Name: "Test trip", Description: "A trip for testing", CityID: 1},
			mockSetup: func() {
				mock.ExpectExec(`INSERT INTO trips`).WithArgs(1, "Test trip", "A trip for testing", 1, sqlmock.AnyArg(), sqlmock.AnyArg(), false).
					WillReturnError(errors.New("exec error"))
			},
			expectedErr: fmt.Errorf("failed to create a trip: %w", models.ErrInternal),
		},
		{
			name: "no rows created",
			trip: models.Trip{UserID: 1, Name: "Test trip", Description: "A trip for testing", CityID: 1},
			mockSetup: func() {
				mock.ExpectExec(`INSERT INTO trips`).WithArgs(1, "Test trip", "A trip for testing", 1, sqlmock.AnyArg(), sqlmock.AnyArg(), false).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedErr: fmt.Errorf("no rows were created: %w", models.ErrNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			err := repo.CreateTrip(context.Background(), tt.trip)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUpdateTrip(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	repo := NewTripRepository(db)

	tests := []struct {
		name        string
		trip        models.Trip
		mockSetup   func()
		expectedErr error
	}{
		{
			name: "successful update",
			trip: models.Trip{ID: 1, Name: "Updated Trip", Description: "Updated description", CityID: 2},
			mockSetup: func() {
				mock.ExpectExec(`UPDATE trips`).WithArgs("Updated Trip", "Updated description", 2, sqlmock.AnyArg(), sqlmock.AnyArg(), false, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedErr: nil,
		},
		{
			name: "error on exec",
			trip: models.Trip{ID: 1, Name: "Updated Trip", Description: "Updated description", CityID: 2},
			mockSetup: func() {
				mock.ExpectExec(`UPDATE trips`).WithArgs("Updated Trip", "Updated description", 2, sqlmock.AnyArg(), sqlmock.AnyArg(), false, 1).
					WillReturnError(errors.New("exec error"))
			},
			expectedErr: fmt.Errorf("failed to execute update query: %w", models.ErrInternal),
		},
		{
			name: "no rows updated",
			trip: models.Trip{ID: 1, Name: "Updated Trip", Description: "Updated description", CityID: 2},
			mockSetup: func() {
				mock.ExpectExec(`UPDATE trips`).WithArgs("Updated Trip", "Updated description", 2, sqlmock.AnyArg(), sqlmock.AnyArg(), false, 1).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedErr: fmt.Errorf("no rows were updated: %w", models.ErrNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			err := repo.UpdateTrip(context.Background(), tt.trip)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestDeleteTrip(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	repo := NewTripRepository(db)

	tests := []struct {
		name        string
		tripID      uint
		mockSetup   func()
		expectedErr error
	}{
		{
			name:   "successful deletion",
			tripID: 1,
			mockSetup: func() {
				mock.ExpectExec(`DELETE FROM trips`).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedErr: nil,
		},
		{
			name:   "error on exec",
			tripID: 1,
			mockSetup: func() {
				mock.ExpectExec(`DELETE FROM trips`).WithArgs(1).WillReturnError(errors.New("exec error"))
			},
			expectedErr: fmt.Errorf("failed to delete trip: %w", models.ErrInternal),
		},
		{
			name:   "no rows deleted",
			tripID: 1,
			mockSetup: func() {
				mock.ExpectExec(`DELETE FROM trips`).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedErr: fmt.Errorf("no rows were deleted: %w", models.ErrNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			err := repo.DeleteTrip(context.Background(), tt.tripID)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetTripsByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()
	ctx := context.Background()
	repo := NewTripRepository(db)
	userID := uint(1)
	limit, offset := 10, 0
	createdAt := time.Now()

	tests := []struct {
		name          string
		mockSetup     func()
		expectedTrips []models.Trip
		expectedError error
	}{
		{
			name: "successful retrieval",
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "name", "description", "city_id", "start_date", "end_date", "private", "created_at"}).
					AddRow(1, userID, "Test trip 1", "A trip for testing", 1, "2024-01-01", "2024-01-05", false, createdAt).
					AddRow(2, userID, "Test trip 2", "A trip for testing", 2, "2024-02-01", "2024-02-10", true, createdAt)

				mock.ExpectQuery(`SELECT id, user_id, name, description, city_id, start_date, end_date, private, created_at FROM trips WHERE user_id = \$1 ORDER BY created_at DESC LIMIT \$2 OFFSET \$3`).
					WithArgs(userID, limit, offset).
					WillReturnRows(rows)
			},
			expectedTrips: []models.Trip{
				{ID: 1, UserID: userID, Name: "Test trip 1", Description: "A trip for testing", CityID: 1, StartDate: "2024-01-01", EndDate: "2024-01-05", Private: false, CreatedAt: createdAt},
				{ID: 2, UserID: userID, Name: "Test trip 2", Description: "A trip for testing", CityID: 2, StartDate: "2024-02-01", EndDate: "2024-02-10", Private: true, CreatedAt: createdAt},
			},
			expectedError: nil,
		},
		{
			name: "query execution error",
			mockSetup: func() {
				mock.ExpectQuery(`SELECT id, user_id, name, description, city_id, start_date, end_date, private, created_at FROM trips WHERE user_id = \$1 ORDER BY created_at DESC LIMIT \$2 OFFSET \$3`).
					WithArgs(userID, limit, offset).
					WillReturnError(models.ErrInternal)
			},
			expectedTrips: nil,
			expectedError: fmt.Errorf("failed to retrieve trips: %w", models.ErrInternal),
		},
		{
			name: "no trips found",
			mockSetup: func() {
				rows := sqlmock.NewRows([]string{"id", "user_id", "name", "description", "city_id", "start_date", "end_date", "private", "created_at"})
				mock.ExpectQuery(`SELECT id, user_id, name, description, city_id, start_date, end_date, private, created_at FROM trips WHERE user_id = \$1 ORDER BY created_at DESC LIMIT \$2 OFFSET \$3`).
					WithArgs(userID, limit, offset).
					WillReturnRows(rows)
			},
			expectedTrips: nil,
			expectedError: fmt.Errorf("no trips found: %w", models.ErrNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			trips, err := repo.GetTripsByUserID(ctx, userID, limit, offset)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedTrips, trips)
		})
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetTrip(t *testing.T) {
	tests := []struct {
		name          string
		tripID        uint
		mockSetup     func(sqlmock.Sqlmock)
		expectedTrip  models.Trip
		expectedError error
	}{
		{
			name:   "successful retrieval",
			tripID: 1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "user_id", "name", "description", "city_id", "start_date", "end_date", "private", "created_at"}).
					AddRow(1, 1, "Test trip", "A trip for testing", 1, "2024-01-01", "2024-01-05", false, time.Now())
				mock.ExpectQuery(`SELECT id, user_id, name, description, city_id, start_date, end_date, private, created_at FROM trips WHERE id = \$1`).
					WithArgs(1).
					WillReturnRows(rows)
			},
			expectedTrip: models.Trip{
				ID:          1,
				UserID:      1,
				Name:        "Test trip",
				Description: "A trip for testing",
				CityID:      1,
				StartDate:   "2024-01-01",
				EndDate:     "2024-01-05",
				Private:     false,
			},
			expectedError: nil,
		},
		{
			name:   "trip not found",
			tripID: 2,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT id, user_id, name, description, city_id, start_date, end_date, private, created_at FROM trips WHERE id = \$1`).
					WithArgs(2).
					WillReturnError(sql.ErrNoRows)
			},
			expectedTrip:  models.Trip{},
			expectedError: fmt.Errorf("trip not found: %w", models.ErrNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)
			defer db.Close()

			repo := NewTripRepository(db)
			tt.mockSetup(mock)

			trip, err := repo.GetTrip(context.Background(), tt.tripID)
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedTrip.ID, trip.ID)
				assert.Equal(t, tt.expectedTrip.UserID, trip.UserID)
				assert.Equal(t, tt.expectedTrip.Name, trip.Name)
				assert.Equal(t, tt.expectedTrip.Description, trip.Description)
				assert.Equal(t, tt.expectedTrip.CityID, trip.CityID)
				assert.Equal(t, tt.expectedTrip.StartDate, trip.StartDate)
				assert.Equal(t, tt.expectedTrip.EndDate, trip.EndDate)
				assert.Equal(t, tt.expectedTrip.Private, trip.Private)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
