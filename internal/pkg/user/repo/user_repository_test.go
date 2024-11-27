package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening stub database connection: %v", err)
	}
	defer db.Close()

	var logBuffer bytes.Buffer

	handler := slog.NewTextHandler(&logBuffer, nil)

	logger := slog.New(handler)
	loggerDB := dblogger.NewDB(db, logger)

	repo := NewAuthRepository(loggerDB)

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

		mock.ExpectQuery(`INSERT INTO "user" \(login, email, password_hash, created_at\)`).
			WithArgs("testuser", "test@example.com", "hashedpassword").
			WillReturnRows(rows)

		user := models.User{
			Login:    "testuser",
			Email:    "test@example.com",
			Password: "hashedpassword",
		}
		userID, err := repo.CreateUser(context.Background(), user)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), userID)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Already Exists", func(t *testing.T) {
		mock.ExpectQuery(`INSERT INTO "user" \(login, email, password_hash, created_at\)`).
			WithArgs("testuser", "test@example.com", "hashedpassword").
			WillReturnError(&pq.Error{Code: "23505"})

		user := models.User{
			Login:    "testuser",
			Email:    "test@example.com",
			Password: "hashedpassword",
		}
		userID, err := repo.CreateUser(context.Background(), user)

		assert.Error(t, err)
		assert.Equal(t, models.ErrAlreadyExists, err)
		assert.Equal(t, uint(0), userID)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Query Execution Error", func(t *testing.T) {
		mock.ExpectQuery(`INSERT INTO "user" \(login, email, password_hash, created_at\)`).
			WithArgs("testuser", "test@example.com", "hashedpassword").
			WillReturnError(fmt.Errorf("query execution error"))

		user := models.User{
			Login:    "testuser",
			Email:    "test@example.com",
			Password: "hashedpassword",
		}
		userID, err := repo.CreateUser(context.Background(), user)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "query execution error")
		assert.Equal(t, uint(0), userID)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening stub database connection: %v", err)
	}
	defer db.Close()

	var logBuffer bytes.Buffer
	handler := slog.NewTextHandler(&logBuffer, nil)
	logger := slog.New(handler)
	loggerDB := dblogger.NewDB(db, logger)

	repo := NewAuthRepository(loggerDB)

	t.Run("Success", func(t *testing.T) {

		rows := sqlmock.NewRows([]string{"id", "login", "email", "password_hash"}).
			AddRow(1, "testuser", "test@example.com", "hashedpassword")

		mock.ExpectQuery(`^SELECT id, login, email, password_hash FROM "user" WHERE email = \$1$`).
			WithArgs("test@example.com").
			WillReturnRows(rows)

		expected := models.User{
			ID:       1,
			Login:    "testuser",
			Email:    "test@example.com",
			Password: "hashedpassword",
		}

		result, err := repo.GetUserByEmail(context.Background(), "test@example.com")

		assert.NoError(t, err)
		assert.Equal(t, expected, result)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("User not found", func(t *testing.T) {
		mock.ExpectQuery(`^SELECT id, login, email, password_hash FROM "user" WHERE email = \$1$`).
			WithArgs("notfound@example.com").
			WillReturnError(sql.ErrNoRows)

		result, err := repo.GetUserByEmail(context.Background(), "notfound@example.com")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "user not found with email")
		assert.Equal(t, models.User{}, result)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Database error", func(t *testing.T) {
		mock.ExpectQuery(`^SELECT id, login, email, password_hash FROM "user" WHERE email = \$1$`).
			WithArgs("test@example.com").
			WillReturnError(fmt.Errorf("database error"))

		result, err := repo.GetUserByEmail(context.Background(), "test@example.com")

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "database error")
		assert.Equal(t, models.User{}, result)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		name          string
		user          models.User
		mockBehavior  func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name: "Success",
			user: models.User{
				ID:       1,
				Login:    "updatedLogin",
				Email:    "updated@example.com",
				Password: "newHashedPassword",
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE "user"`).
					WithArgs("updatedLogin", "updated@example.com", "newHashedPassword", 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Failed to Update User - Query Error",
			user: models.User{
				ID:       1,
				Login:    "updatedLogin",
				Email:    "updated@example.com",
				Password: "newHashedPassword",
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE "user"`).
					WithArgs("updatedLogin", "updated@example.com", "newHashedPassword", 1).
					WillReturnError(errors.New("update failed"))
			},
			expectedError: errors.New("update failed"),
		},
		{
			name: "No Rows Affected",
			user: models.User{
				ID:       1,
				Login:    "updatedLogin",
				Email:    "updated@example.com",
				Password: "newHashedPassword",
			},
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`UPDATE "user"`).
					WithArgs("updatedLogin", "updated@example.com", "newHashedPassword", 1).
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: nil,
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

			repo := &UserRepositoryImpl{loggerDB}

			tt.mockBehavior(mock)

			err := repo.UpdateUser(context.Background(), tt.user)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name          string
		userID        string
		mockBehavior  func(mock sqlmock.Sqlmock)
		expectedError error
	}{
		{
			name:   "Success",
			userID: "1",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM "user"`).
					WithArgs("1").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name:   "Delete Failed",
			userID: "1",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM "user"`).
					WithArgs("1").
					WillReturnError(errors.New("delete failed"))
			},
			expectedError: errors.New("delete failed"),
		},
		{
			name:   "No Rows Affected",
			userID: "1",
			mockBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM "user"`).
					WithArgs("1").
					WillReturnResult(sqlmock.NewResult(1, 0))
			},
			expectedError: nil,
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

			repo := &UserRepositoryImpl{loggerDB}

			tt.mockBehavior(mock)

			err := repo.DeleteUser(context.Background(), tt.userID)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetAvatarPathByUserId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening stub database connection: %v", err)
	}
	defer db.Close()

	var logBuffer bytes.Buffer
	handler := slog.NewTextHandler(&logBuffer, nil)
	logger := slog.New(handler)
	loggerDB := dblogger.NewDB(db, logger)

	repo := &UserRepositoryImpl{loggerDB}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectQuery(`SELECT avatar_path FROM "user"`).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"avatar_path"}).AddRow("user1.png"))

		expectedPath := "user1.png"

		avatarPath, err := repo.GetAvatarPathByUserId(context.Background(), 1)

		assert.NoError(t, err)
		assert.Equal(t, expectedPath, avatarPath)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("User Not Found", func(t *testing.T) {
		mock.ExpectQuery(`SELECT avatar_path FROM "user"`).
			WithArgs(1).
			WillReturnError(sql.ErrNoRows)

		avatarPath, err := repo.GetAvatarPathByUserId(context.Background(), 1)

		assert.ErrorIs(t, err, models.ErrNotFound)
		assert.Empty(t, avatarPath)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Error - Query Execution", func(t *testing.T) {
		mock.ExpectQuery(`SELECT avatar_path FROM "user"`).
			WithArgs(1).
			WillReturnError(errors.New("query execution error"))

		avatarPath, err := repo.GetAvatarPathByUserId(context.Background(), 1)

		assert.ErrorIs(t, err, models.ErrInternal)
		assert.Empty(t, avatarPath)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestUpdateAvatarPathByUserId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening stub database connection: %v", err)
	}
	defer db.Close()

	var logBuffer bytes.Buffer
	handler := slog.NewTextHandler(&logBuffer, nil)
	logger := slog.New(handler)
	loggerDB := dblogger.NewDB(db, logger)

	repo := &UserRepositoryImpl{loggerDB}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`UPDATE "user"`).
			WithArgs("user1.png", 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.UpdateAvatarPathByUserId(context.Background(), 1, "user1.png")

		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("No Rows Affected", func(t *testing.T) {
		mock.ExpectExec(`UPDATE "user"`).
			WithArgs("user1.png", 1).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.UpdateAvatarPathByUserId(context.Background(), 1, "user1.png")

		assert.ErrorIs(t, err, models.ErrNotFound)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Error - Query Execution", func(t *testing.T) {
		mock.ExpectExec(`UPDATE "user"`).
			WithArgs("user1.png", 1).
			WillReturnError(errors.New("query execution error"))

		err := repo.UpdateAvatarPathByUserId(context.Background(), 1, "user1.png")

		assert.ErrorIs(t, err, models.ErrInternal)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Error - RowsAffected", func(t *testing.T) {
		mock.ExpectExec(`UPDATE "user"`).
			WithArgs("user1.png", 1).
			WillReturnResult(sqlmock.NewResult(1, 0))

		err := repo.UpdateAvatarPathByUserId(context.Background(), 1, "user1.png")

		assert.ErrorIs(t, err, models.ErrNotFound)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestGetUserByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening stub database connection: %v", err)
	}
	defer db.Close()

	var logBuffer bytes.Buffer
	handler := slog.NewTextHandler(&logBuffer, nil)
	logger := slog.New(handler)
	loggerDB := dblogger.NewDB(db, logger)

	repo := &UserRepositoryImpl{loggerDB}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectQuery(`SELECT login, email, avatar_path`).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"login", "email", "avatar_path"}).
				AddRow("user1", "user1@example.com", "user1.png"))

		expectedProfile := models.UserProfile{
			Login:      "user1",
			Email:      "user1@example.com",
			AvatarPath: "user1.png",
		}

		userProfile, err := repo.GetUserByID(context.Background(), 1)

		assert.NoError(t, err)
		assert.Equal(t, expectedProfile, userProfile)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("User Not Found", func(t *testing.T) {
		mock.ExpectQuery(`SELECT login, email, avatar_path`).
			WithArgs(1).
			WillReturnError(sql.ErrNoRows)

		userProfile, err := repo.GetUserByID(context.Background(), 1)

		assert.ErrorIs(t, err, models.ErrNotFound)
		assert.Equal(t, models.UserProfile{}, userProfile)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Error - Query Execution", func(t *testing.T) {
		mock.ExpectQuery(`SELECT login, email, avatar_path`).
			WithArgs(1).
			WillReturnError(errors.New("query execution error"))

		userProfile, err := repo.GetUserByID(context.Background(), 1)

		assert.ErrorIs(t, err, models.ErrInternal)
		assert.Equal(t, models.UserProfile{}, userProfile)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Error - Scan", func(t *testing.T) {
		mock.ExpectQuery(`SELECT login, email, avatar_path`).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"login", "email", "avatar_path"}).
				AddRow("user1", "user1@example.com", nil))

		userProfile, err := repo.GetUserByID(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, models.UserProfile{}, userProfile)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestUpdatePassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening stub database connection: %v", err)
	}
	defer db.Close()

	var logBuffer bytes.Buffer
	handler := slog.NewTextHandler(&logBuffer, nil)
	logger := slog.New(handler)
	loggerDB := dblogger.NewDB(db, logger)

	repo := &UserRepositoryImpl{loggerDB}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`UPDATE user SET password = \$1 WHERE id = \$2`).
			WithArgs("newPassword123", 1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.UpdatePassword(context.Background(), 1, "newPassword123")

		assert.NoError(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Error - Query Execution", func(t *testing.T) {
		mock.ExpectExec(`UPDATE user SET password = \$1 WHERE id = \$2`).
			WithArgs("newPassword123", 1).
			WillReturnError(errors.New("query execution error"))

		err := repo.UpdatePassword(context.Background(), 1, "newPassword123")

		assert.ErrorIs(t, err, models.ErrInternal)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Error - No Rows Affected", func(t *testing.T) {
		mock.ExpectExec(`UPDATE user SET password = \$1 WHERE id = \$2`).
			WithArgs("newPassword123", 1).
			WillReturnResult(sqlmock.NewResult(0, 0))

		err := repo.UpdatePassword(context.Background(), 1, "newPassword123")

		assert.ErrorIs(t, err, models.ErrNotFound)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func TestUpdateProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening stub database connection: %v", err)
	}
	defer db.Close()

	var logBuffer bytes.Buffer
	handler := slog.NewTextHandler(&logBuffer, nil)
	logger := slog.New(handler)
	loggerDB := dblogger.NewDB(db, logger)

	repo := &UserRepositoryImpl{loggerDB}

	t.Run("Success", func(t *testing.T) {
		// Мокаем успешное выполнение запроса на обновление профиля
		mock.ExpectExec(`UPDATE user SET email = \$1, login = \$2 WHERE id = \$3`).
			WithArgs("newemail@example.com", "newuser", 1).
			WillReturnResult(sqlmock.NewResult(1, 1)) // 1 строка обновлена

		err := repo.UpdateProfile(context.Background(), 1, "newuser", "newemail@example.com")

		assert.NoError(t, err)

		// Проверка выполнения всех ожиданий mock
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Error - Query Execution", func(t *testing.T) {
		// Мокаем ошибку при выполнении запроса
		mock.ExpectExec(`UPDATE user SET email = \$1, login = \$2 WHERE id = \$3`).
			WithArgs("newemail@example.com", "newuser", 1).
			WillReturnError(errors.New("query execution error"))

		err := repo.UpdateProfile(context.Background(), 1, "newuser", "newemail@example.com")

		assert.ErrorIs(t, err, models.ErrInternal)

		// Проверка выполнения всех ожиданий mock
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("Error - No Rows Affected", func(t *testing.T) {
		// Мокаем ситуацию, когда обновление не затронуло ни одной строки
		mock.ExpectExec(`UPDATE user SET email = \$1, login = \$2 WHERE id = \$3`).
			WithArgs("newemail@example.com", "newuser", 1).
			WillReturnResult(sqlmock.NewResult(0, 0)) // 0 строк обновлено

		err := repo.UpdateProfile(context.Background(), 1, "newuser", "newemail@example.com")

		// Мы проверяем, что возвращается ошибка models.ErrNotFound
		assert.ErrorIs(t, err, models.ErrNotFound)

		// Проверка выполнения всех ожиданий mock
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
