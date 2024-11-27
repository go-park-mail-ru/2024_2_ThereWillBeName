package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCategoryRepoGetCategories(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()

	var logBuffer bytes.Buffer

	handler := slog.NewTextHandler(&logBuffer, nil)

	logger := slog.New(handler)
	loggerDB := dblogger.NewDB(db, logger)

	r := NewCategoriesRepo(loggerDB)

	tests := []struct {
		name        string
		categories  []models.Category
		mockSetup   func()
		expectedErr error
	}{
		{
			name: "successful",
			categories: []models.Category{
				{ID: 1, Name: "театр"},
				{ID: 2, Name: "собор"},
			},
			mockSetup: func() {
				mock.ExpectQuery("SELECT id, name FROM category").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
						AddRow(1, "театр").AddRow(2, "собор"))
			},
			expectedErr: nil,
		},
		{
			name:       "failureDb",
			categories: []models.Category(nil),
			mockSetup: func() {
				mock.ExpectQuery("SELECT id, name FROM category").
					WillReturnError(fmt.Errorf("error"))
			},
			expectedErr: fmt.Errorf("couldn't get categories: error"),
		},
		{
			name:       "failureUnmarshal",
			categories: nil,
			mockSetup: func() {
				mock.ExpectQuery("SELECT id, name FROM category").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "fail"}).
						AddRow(0, "name", "fail"))
			},
			expectedErr: fmt.Errorf("Couldn't unmarshal list of categories: sql: expected 3 destination arguments in Scan, not 2"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			res, err := r.GetCategories(context.Background(), 10, 0)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.categories, res)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
