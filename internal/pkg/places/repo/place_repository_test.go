package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"

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

	mock.ExpectQuery(regexp.QuoteMeta("SELECT p.id, p.name, p.image_path, p.description, p.rating, p.address, p.phone_number, c.name AS city_name, ARRAY_AGG(ca.name) AS categories FROM place p JOIN city c ON p.city_id = c.id JOIN place_category pc ON p.id = pc.place_id JOIN category ca ON pc.category_id = ca.id GROUP BY p.id, c.name ORDER BY p.id LIMIT $1 OFFSET $2")).
		WithArgs(10, 0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "image_path", "description", "rating", "address", "city", "phone_number", "categories"}).
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

	mock.ExpectQuery("SELECT id, name, image, description FROM place").
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
	mock.ExpectQuery("SELECT id, name, image, description FROM place").
		WillReturnRows(rows)
	r := NewPLaceRepository(db)
	places, err := r.GetPlaces(context.Background(), 10, 0)
	fmt.Println(places, err)
	assert.Error(t, err)
	assert.Nil(t, places)
}

//func TestPlaceRepository_CreatePlace(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("failed to open mock sql database: %v", err)
//	}
//	defer db.Close()
//	r := NewPLaceRepository(db)
//
//	tests := []struct {
//		name        string
//		place       models.CreatePlace
//		mockSetup   func()
//		expectedErr error
//	}{
//		{
//			name: "succesfull",
//			place: models.CreatePlace{
//				Name:            "Test Place",
//				ImagePath:       "/path/to/image",
//				Description:     "Test Description",
//				Rating:          4,
//				NumberOfReviews: 10,
//				Address:         "Test Address",
//				CityId:          1,
//				PhoneNumber:     "1234567890",
//				CategoriesId:    []int{1, 2, 3},
//			},
//			mockSetup: func() {
//				mock.ExpectQueryRow("INSERT INTO place (name, imagePath, description, rating, numberOfReviews, address, cityId, phoneNumber) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id").WithArgs("Test Place", "/path/to/image", "Test Description", 4, 10, "Test Address", 1, "1234567890").
//					WillReturnResult(sqlmock.NewResult(1, 1))
//			},
//			expectedErr: nil,
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.mockSetup()
//			err := r.CreatePlace(context.Background(), tt.place)
//			if tt.expectedErr != nil {
//				assert.EqualError(t, err, tt.expectedErr.Error())
//			} else {
//				assert.NoError(t, err)
//			}
//			assert.NoError(t, mock.ExpectationsWereMet())
//		})
//	}
//
//}

//func TestPlaceRepository_CreatePlace(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockPlaceRepo := mock_places.NewMockPlaceRepo(ctrl)
//
//	ctx := context.Background()
//
//	place := models.CreatePlace{
//		Name:            "Test Place",
//		ImagePath:       "/path/to/image",
//		Description:     "Test Description",
//		Rating:          4,
//		NumberOfReviews: 10,
//		Address:         "Test Address",
//		CityId:          1,
//		PhoneNumber:     "1234567890",
//		CategoriesId:    []int{1, 2, 3},
//	}
//
//	t.Run("success", func(t *testing.T) {
//		// Ожидаем вызов QueryRowContext
//		mockPlaceRepo.EXPECT().QueryRowContext(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
//			DoAndReturn(func(ctx context.Context, query string, args ...interface{}) *sql.Row {
//				// Возвращаем mock для Row, который вернет id
//				return &sql.Row{
//					Scan: func(dest ...interface{}) error {
//						// Устанавливаем id в 1
//						*(dest[0].(*int)) = 1
//						return nil
//					},
//				}
//			})
//
//		// Ожидаем вызов ExecContext для каждой категории
//		for _, categoryID := range place.CategoriesId {
//			mockPlaceRepo.EXPECT().ExecContext(ctx, gomock.Any(), gomock.Any(), categoryID).
//				Return(sql.Result(&sql.RowsAffected{}), nil)
//		}
//
//		err := mockPlaceRepo.CreatePlace(ctx, place)
//		assert.NoError(t, err)
//	})
//
//	t.Run("error on create place", func(t *testing.T) {
//		// Ожидаем вызов QueryRowContext с ошибкой
//		mockPlaceRepo.EXPECT().QueryRowContext(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
//			Return(&sql.Row{
//				Scan: func(dest ...interface{}) error {
//					return errors.New("database error")
//				},
//			})
//
//		err := mockPlaceRepo.CreatePlace(ctx, place)
//		assert.Error(t, err)
//		assert.EqualError(t, err, "coldn't create place: database error")
//	})
//
//	t.Run("error on create place_category", func(t *testing.T) {
//		// Ожидаем вызов QueryRowContext
//		mockPlaceRepo.EXPECT().QueryRowContext(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
//			DoAndReturn(func(ctx context.Context, query string, args ...interface{}) *sql.Row {
//				// Возвращаем mock для Row, который вернет id
//				return &sql.Row{
//					Scan: func(dest ...interface{}) error {
//						// Устанавливаем id в 1
//						*(dest[0].(*int)) = 1
//						return nil
//					},
//				}
//			})
//
//		// Ожидаем вызов ExecContext с ошибкой
//		mockPlaceRepo.EXPECT().ExecContext(ctx, gomock.Any(), gomock.Any(), gomock.Any()).
//			Return(nil, errors.New("database error"))
//
//		err := mockPlaceRepo.CreatePlace(ctx, place)
//		assert.Error(t, err)
//		assert.EqualError(t, err, "coldn't create place_category: database error")
//	})
//}
