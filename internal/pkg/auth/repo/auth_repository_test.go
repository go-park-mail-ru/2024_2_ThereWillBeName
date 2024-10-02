package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"fmt"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()

	repository := NewAuthRepository(db)

	mock.ExpectExec("INSERT INTO users").
		WithArgs("testuser", "testpass").
		WillReturnResult(sqlmock.NewResult(1, 1))
	user := models.User{
		Login:    "testuser",
		Password: "testpass",
	}

	err = repository.CreateUser(context.Background(), user)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByLogin_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()

	repository := NewAuthRepository(db)

	createdAt := time.Now()
	mock.ExpectQuery("SELECT id, login, password, created_at").
		WithArgs("testuser").
		WillReturnRows(sqlmock.NewRows([]string{"id", "login", "password", "created_at"}).
			AddRow(1, "testuser", "testpass", createdAt))

	user, err := repository.GetUserByLogin(context.Background(), "testuser")

	assert.NoError(t, err)
	assert.Equal(t, user.Login, "testuser")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateUser_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()

	repository := NewAuthRepository(db)

	mock.ExpectExec("INSERT INTO users").
		WithArgs("testuser", "testpass").
		WillReturnError(fmt.Errorf("database error"))

	user := models.User{
		Login:    "testuser",
		Password: "testpass",
	}

	err = repository.CreateUser(context.Background(), user)

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()

	repository := NewAuthRepository(db)

	mock.ExpectExec("UPDATE users").
		WithArgs("updateduser", "newpassword", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	user := models.User{
		ID:       1,
		Login:    "updateduser",
		Password: "newpassword",
	}

	err = repository.UpdateUser(context.Background(), user)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteUser_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()

	repository := NewAuthRepository(db)

	mock.ExpectExec("DELETE FROM users").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repository.DeleteUser(context.Background(), "1")

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUsers_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()

	repository := NewAuthRepository(db)

	mock.ExpectQuery("SELECT id, login, created_at").
		WithArgs(int64(10), int64(0)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "login", "created_at"}).
			AddRow(1, "testuser", time.Now()))

	users, err := repository.GetUsers(context.Background(), 10, 0)

	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, users[0].Login, "testuser")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByLogin_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()

	repository := NewAuthRepository(db)

	mock.ExpectQuery("SELECT id, login, password, created_at").
		WithArgs("nonexistentuser").
		WillReturnRows(sqlmock.NewRows([]string{"id", "login", "password", "created_at"}))

	user, err := repository.GetUserByLogin(context.Background(), "nonexistentuser")

	assert.Error(t, err)
	assert.Equal(t, models.User{}, user)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateUser_AlreadyExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()

	repository := NewAuthRepository(db)

	mock.ExpectExec("INSERT INTO users").
		WithArgs("existinguser", "testpass").
		WillReturnError(fmt.Errorf("duplicate entry error"))

	user := models.User{
		Login:    "existinguser",
		Password: "testpass",
	}

	err = repository.CreateUser(context.Background(), user)

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUsers_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()

	repository := NewAuthRepository(db)

	mock.ExpectQuery("SELECT id, login, created_at").
		WithArgs(int64(10), int64(0)).
		WillReturnError(fmt.Errorf("database error"))

	users, err := repository.GetUsers(context.Background(), 10, 0)

	assert.Error(t, err)
	assert.Empty(t, users)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByLogin_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()

	repository := NewAuthRepository(db)

	mock.ExpectQuery("SELECT id, login, password, created_at").
		WithArgs("testuser").
		WillReturnError(fmt.Errorf("database error"))

	user, err := repository.GetUserByLogin(context.Background(), "testuser")

	assert.Error(t, err)
	assert.Equal(t, models.User{}, user)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
