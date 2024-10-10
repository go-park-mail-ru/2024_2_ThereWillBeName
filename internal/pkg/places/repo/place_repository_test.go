package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaceRepository_GetPlaces(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, name, image, description FROM places").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "image", "description"}).
			AddRow(0, "testName", "testImage.png", "testDescription"))

	expectedCode := []models.Place{{ID: 0, Name: "testName", Image: "testImage.png", Description: "testDescription"}}

	r := NewPLaceRepository(db)
	places, err := r.GetPlaces(context.Background())

	assert.NoError(t, err)
	assert.Len(t, places, len(expectedCode))
	assert.Equal(t, expectedCode, places)
}

func TestPlaceRepository_GetPlaces_DbError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, name, image, description FROM places").
		WillReturnError(fmt.Errorf("couldn't get places: %w", err))

	r := NewPLaceRepository(db)
	places, err := r.GetPlaces(context.Background())

	assert.Error(t, err)
	assert.Nil(t, places)
}

func TestPlaceRepository_GetPlaces_ParseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "name", "image", "description", "fail"}).
		AddRow(0, "name", "image", "description", "fail")
	mock.ExpectQuery("SELECT id, name, image, description FROM places").
		WillReturnRows(rows)
	r := NewPLaceRepository(db)
	places, err := r.GetPlaces(context.Background())
	fmt.Println(places, err)
	assert.Error(t, err)
	assert.Nil(t, places)
}
