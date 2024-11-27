package search

import (
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSearchCitiesAndPlacesBySubString(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening stub database connection: %v", err)
	}
	defer db.Close()

	// Создаем репозиторий
	var logBuffer bytes.Buffer

	handler := slog.NewTextHandler(&logBuffer, nil)
	logger := slog.New(handler)
	loggerDB := dblogger.NewDB(db, logger)

	repo := NewSearchRepository(loggerDB)
	// Тест 1: Успешный поиск
	t.Run("Success", func(t *testing.T) {
		// Мокируем возвращаемые строки (результаты поиска)
		rows := sqlmock.NewRows([]string{"id", "name", "type"}).
			AddRow(1, "New York", "city").
			AddRow(2, "Central Park", "place")

		// Ожидаем выполнение запроса
		mock.ExpectQuery(`SELECT id, name, 'city' AS type`).
			WithArgs("New").
			WillReturnRows(rows)

		// Выполняем тестируемую функцию
		query := "New"
		results, err := repo.SearchCitiesAndPlacesBySubString(context.Background(), query)

		// Проверяем, что ошибки нет и данные соответствуют ожиданиям
		assert.NoError(t, err)
		assert.Len(t, results, 2)
		assert.Equal(t, "New York", results[0].Name)
		assert.Equal(t, "city", results[0].Type)
		assert.Equal(t, "Central Park", results[1].Name)
		assert.Equal(t, "place", results[1].Type)

		// Проверяем, что все ожидания выполнены
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	// Тест 2: Отсутствие результатов
	t.Run("No Results Found", func(t *testing.T) {
		// Мокируем пустой набор строк (нет результатов)
		rows := sqlmock.NewRows([]string{"id", "name", "type"})

		// Ожидаем выполнение запроса
		mock.ExpectQuery(`SELECT id, name, 'city' AS type`).
			WithArgs("Unknown").
			WillReturnRows(rows)

		// Выполняем тестируемую функцию
		query := "Unknown"
		results, err := repo.SearchCitiesAndPlacesBySubString(context.Background(), query)

		// Проверяем, что возникла ошибка о ненахождении результатов
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no results found")
		assert.Nil(t, results)

		// Проверяем, что все ожидания выполнены
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	// Тест 3: Ошибка выполнения запроса
	t.Run("Query Execution Error", func(t *testing.T) {
		// Мокируем ошибку при выполнении запроса
		mock.ExpectQuery(`SELECT id, name, 'city' AS type`).
			WithArgs("New").
			WillReturnError(fmt.Errorf("query execution error"))

		// Выполняем тестируемую функцию
		query := "New"
		results, err := repo.SearchCitiesAndPlacesBySubString(context.Background(), query)

		// Проверяем, что ошибка выполнения запроса возвращена
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to execute search query")
		assert.Nil(t, results)

		// Проверяем, что все ожидания выполнены
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
