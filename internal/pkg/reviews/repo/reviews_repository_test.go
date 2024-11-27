package reviews

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"reflect"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetReview(t *testing.T) {
	tests := []struct {
		name           string
		reviewID       uint
		mockBehavior   func(mock sqlmock.Sqlmock)
		expectedError  error
		expectedReview models.GetReview
	}{
		{
			name:     "Success",
			reviewID: 1,
			mockBehavior: func(mock sqlmock.Sqlmock) {
				// Мокаем запрос для получения данных отзыва
				mock.ExpectQuery(`SELECT r.id, u.login, u.avatar_path, r.rating, r.review_text`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_login", "avatar_path", "rating", "review_text"}).
						AddRow(1, "test_user", "/path/to/avatar", 5, "Great place!"))
			},
			expectedError: nil,
			expectedReview: models.GetReview{
				ID:         1,
				UserLogin:  "test_user",
				AvatarPath: "/path/to/avatar",
				Rating:     5,
				ReviewText: "Great place!",
			},
		},
		{
			name:     "Review Not Found",
			reviewID: 99,
			mockBehavior: func(mock sqlmock.Sqlmock) {
				// Мокаем запрос, который не находит отзыва
				mock.ExpectQuery(`SELECT r.id, u.login, u.avatar_path, r.rating, r.review_text`).
					WithArgs(99).
					WillReturnError(sql.ErrNoRows) // Нет таких данных
			},
			expectedError:  models.ErrNotFound,
			expectedReview: models.GetReview{}, // Пустой объект, так как отзыва нет
		},
		// Добавьте другие тесты, если нужно
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()
			var logBuffer bytes.Buffer

			handler := slog.NewTextHandler(&logBuffer, nil)
			logger := slog.New(handler)
			loggerDB := dblogger.NewDB(db, logger)

			repo := NewReviewRepository(loggerDB)

			tt.mockBehavior(mock)

			createdReview, err := repo.GetReview(context.Background(), tt.reviewID)
			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedReview, createdReview)
			}
		})
	}
}

func TestCreateReview(t *testing.T) {
	tests := []struct {
		name           string
		review         models.Review
		mockBehavior   func(mock sqlmock.Sqlmock)
		expectedError  error
		expectedReview models.GetReview
	}{
		{
			name: "Success",
			review: models.Review{
				UserID:     1,
				PlaceID:    2,
				Rating:     5,
				ReviewText: "Great place!",
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`INSERT INTO review`).
					WithArgs(1, 2, 5, "Great place!").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1)) // ID 1 возвращается при успешной вставке

				mock.ExpectQuery(`SELECT r.id, u.login, u.avatar_path, r.rating, r.review_text`).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_login", "avatar_path", "rating", "review_text"}).
						AddRow(1, "test_user", "/path/to/avatar", 5, "Great place!"))
			},
			expectedError: nil,
			expectedReview: models.GetReview{
				ID:         1,
				UserLogin:  "test_user",
				AvatarPath: "/path/to/avatar",
				Rating:     5,
				ReviewText: "Great place!",
			},
		},
		{
			name: "Failed to Insert Review",
			review: models.Review{
				UserID:     1,
				PlaceID:    2,
				Rating:     5,
				ReviewText: "Great place!",
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`INSERT INTO review`).
					WithArgs(1, 2, 5, "Great place!").
					WillReturnError(fmt.Errorf("internal repository error"))
			},
			expectedError: fmt.Errorf("failed to create review: %w", models.ErrInternal),
		},
		{
			name: "Failed to Retrieve Created Review",
			review: models.Review{
				UserID:     1,
				PlaceID:    2,
				Rating:     5,
				ReviewText: "Great place!",
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`INSERT INTO review`).
					WithArgs(1, 2, 5, "Great place!").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectQuery(`SELECT r.id, u.login, u.avatar_path`).
					WithArgs(1).
					WillReturnError(fmt.Errorf("failed to scan review"))
			},
			expectedError: fmt.Errorf("failed to retrieve created review details: failed to scan review"),
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

			repo := NewReviewRepository(loggerDB)

			tt.mockBehavior(mock)

			createdReview, err := repo.CreateReview(context.Background(), tt.review)

			if tt.expectedError != nil {
				if !strings.Contains(err.Error(), tt.expectedError.Error()) {
					t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if !reflect.DeepEqual(createdReview, tt.expectedReview) {
					t.Errorf("expected review: %+v, got: %+v", tt.expectedReview, createdReview)
				}
			}
		})
	}
}

func TestUpdateReview(t *testing.T) {
	tests := []struct {
		name          string
		review        models.Review
		mockBehavior  func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "Success",
			review: models.Review{
				ID:         1,
				Rating:     5,
				ReviewText: "Updated review text",
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE review`).
					WithArgs(5, "Updated review text", 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Failed to Update Review - Query Error",
			review: models.Review{
				ID:         1,
				Rating:     5,
				ReviewText: "Updated review text",
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE review`).
					WithArgs(5, "Updated review text", 1).
					WillReturnError(errors.New("update failed"))
			},
			expectedError: models.ErrInternal,
		},
		{
			name: "No Rows Affected",
			review: models.Review{
				ID:         1,
				Rating:     5,
				ReviewText: "Updated review text",
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE review`).
					WithArgs(5, "Updated review text", 1).
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

			repo := NewReviewRepository(loggerDB)

			tt.mockBehavior(mock)

			err := repo.UpdateReview(context.Background(), tt.review)

			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteReview(t *testing.T) {
	tests := []struct {
		name          string
		reviewID      uint
		mockBehavior  func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name:     "Success",
			reviewID: 1,
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM review`).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name:     "Delete Failed",
			reviewID: 1,
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM review`).
					WithArgs(1).
					WillReturnError(errors.New("delete failed"))
			},
			expectedError: models.ErrInternal,
		},
		{
			name:     "No Rows Affected",
			reviewID: 1,
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM review`).
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

			repo := NewReviewRepository(loggerDB)

			tt.mockBehavior(mock)

			err := repo.DeleteReview(context.Background(), tt.reviewID)

			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetReviewsByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening stub database connection: %v", err)
	}
	defer db.Close()

	var logBuffer bytes.Buffer
	handler := slog.NewTextHandler(&logBuffer, nil)
	logger := slog.New(handler)
	loggerDB := dblogger.NewDB(db, logger)

	repo := NewReviewRepository(loggerDB)

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "rating", "review_text", "name"}).
			AddRow(1, 5, "Great place!", "Place 1").
			AddRow(2, 4, "Nice food", "Place 2")

		mock.ExpectQuery(`SELECT r.id, r.rating, r.review_text, p.name`).
			WithArgs(1, 10, 0).
			WillReturnRows(rows)

		expected := []models.GetReviewByUserID{
			{ID: 1, Rating: 5, ReviewText: "Great place!", PlaceName: "Place 1"},
			{ID: 2, Rating: 4, ReviewText: "Nice food", PlaceName: "Place 2"},
		}

		result, err := repo.GetReviewsByUserID(context.Background(), 1, 10, 0)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("No Reviews Found", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "rating", "review_text", "name"})

		mock.ExpectQuery(`SELECT r.id, r.rating, r.review_text, p.name`).
			WithArgs(1, 10, 0).
			WillReturnRows(rows)

		result, err := repo.GetReviewsByUserID(context.Background(), 1, 10, 0)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no reviews found")
		assert.Nil(t, result)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT r.id, r.rating, r.review_text, p.name`).
			WithArgs(1, 10, 0).
			WillReturnError(fmt.Errorf("database error"))

		result, err := repo.GetReviewsByUserID(context.Background(), 1, 10, 0)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to retrieve reviews")
		assert.Nil(t, result)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

}
