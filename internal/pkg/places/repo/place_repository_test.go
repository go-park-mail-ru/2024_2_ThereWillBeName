package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaceRepository_GetPlaces(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open mock sql database: %v", err)
	}
	defer db.Close()

	categories := []string{"Park", "Recreation", "Nature"}
	//categoriesStr := strings.Join(categories, ",")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT p.id, p.name, p.imagePath, p.description, p.rating, p.numberOfReviews, p.address, p.phoneNumber, c.name AS city_name, ARRAY_AGG(ca.name) AS categories FROM places p JOIN cities c ON p.cityId = c.id JOIN places_categories pc ON p.id = pc.place_id JOIN categories ca ON pc.category_id = ca.id GROUP BY p.id, c.name ORDER BY p.id LIMIT $1 OFFSET $2")).
		WithArgs(10, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "imagePath", "description", "rating", "numberOfReviews", "address", "city", "phoneNumber", "categories"}).
			AddRow(1, "Central Park", "/images/central_park.jpg", "A large public park in New York City, offering a variety of recreational activities.", 5, 2500, "59th St to 110th St, New York, NY 10022", "+1 212-310-6600", "New York", pq.Array(categories)))

	expectedCode := []models.GetPlace{{
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
	}}

	r := NewPLaceRepository(db)
	places, err := r.GetPlaces(context.Background(), 10, 0)

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
	places, err := r.GetPlaces(context.Background(), 10, 0)

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
	places, err := r.GetPlaces(context.Background(), 10, 0)
	fmt.Println(places, err)
	assert.Error(t, err)
	assert.Nil(t, places)
}
