package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"github.com/lib/pq"
)

type PlaceRepository struct {
	db *sql.DB
}

func NewPLaceRepository(db *sql.DB) *PlaceRepository {
	return &PlaceRepository{db: db}
}

func (r *PlaceRepository) GetPlaces(ctx context.Context, limit, offset int) ([]models.GetPlace, error) {
	query := "SELECT p.id, p.name, p.imagePath, p.description, p.rating, p.numberOfReviews, p.address, p.phoneNumber, c.name AS city_name, ARRAY_AGG(ca.name) AS categories FROM places p JOIN cities c ON p.cityId = c.id JOIN places_categories pc ON p.id = pc.place_id JOIN categories ca ON pc.category_id = ca.id GROUP BY p.id, c.name ORDER BY p.id LIMIT $1 OFFSET $2"
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("couldn't get places: %w", err)
	}
	defer rows.Close()
	var places []models.GetPlace
	for rows.Next() {
		var place models.GetPlace
		err := rows.Scan(&place.ID, &place.Name, &place.ImagePath, &place.Description, &place.Rating, &place.NumberOfReviews, &place.Address, &place.PhoneNumber, &place.City, pq.Array(&place.Categories))
		if err != nil {
			return nil, fmt.Errorf("couldn't unmarshal list of places: %w", err)
		}
		places = append(places, place)
	}
	return places, nil
}

func (r *PlaceRepository) CreatePlace(ctx context.Context, place models.CreatePlace) error {
	query := "INSERT INTO places (name, imagePath, description, rating, numberOfReviews, address, cityId, phoneNumber) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	var id int
	err := r.db.QueryRowContext(ctx, query, place.Name, place.ImagePath, place.Description, place.Rating, place.NumberOfReviews, place.Address, place.CityId, place.PhoneNumber).Scan(&id)
	if err != nil {
		return fmt.Errorf("coldn't create place: %w", err)
	}
	for _, categoryId := range place.CategoriesId {
		query = "INSERT INTO places_categories (place_id, category_id) VALUES ($1, $2)"
		result, err := r.db.ExecContext(ctx, query, id, categoryId)
		if err != nil {
			return fmt.Errorf("coldn't create place_categories: %w", err)
		}
		if _, err = result.RowsAffected(); err != nil {
			return fmt.Errorf("couldn't get number of rows affected: %w", err)
		}
	}
	return nil
}

func (r *PlaceRepository) GetPlace(ctx context.Context, id uint) (models.GetPlace, error) {
	var place models.GetPlace
	query := "SELECT p.id, p.name, p.imagePath, p.description, p.rating, p.numberOfReviews, p.address, p.phoneNumber, c.name AS city_name, ARRAY_AGG(ca.name) AS categories FROM places p JOIN cities c ON p.cityId = c.id JOIN places_categories pc ON p.id = pc.place_id JOIN categories ca ON pc.category_id = ca.id WHERE p.id = $1 GROUP BY p.id, c.name ORDER BY p.id"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&place.ID, &place.Name, &place.ImagePath, &place.Description, &place.Rating, &place.NumberOfReviews, &place.Address, &place.PhoneNumber, &place.City, pq.Array(&place.Categories))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.GetPlace{}, fmt.Errorf("place not found")
		}
		return models.GetPlace{}, fmt.Errorf("couldn't get place by name: %w", err)
	}
	return place, nil
}

func (r *PlaceRepository) UpdatePlace(ctx context.Context, place models.UpdatePlace) error {
	query := "UPDATE places SET name = $1, imagePath = $2, description = $3, rating = $4, numberOfReviews = $5, address = $6, phoneNumber = $7 WHERE id=$8"
	result, err := r.db.ExecContext(ctx, query, place.Name, place.ImagePath, place.Description, place.Rating, place.NumberOfReviews, place.Address, place.PhoneNumber, place.ID)
	if err != nil {
		return fmt.Errorf("couldn't update place: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("couldn't get number of rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	query = "DELETE FROM places_categories WHERE place_id=$1"
	result, err = r.db.ExecContext(ctx, query, place.ID)
	if err != nil {
		return fmt.Errorf("couldn't delete places_categories: %w", err)
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("couldn't get number of rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted")
	}

	for _, categoryId := range place.CategoriesId {
		query = "INSERT INTO places_categories (place_id, category_id) VALUES ($1, $2)"
		result, err := r.db.ExecContext(ctx, query, place.ID, categoryId)
		if err != nil {
			return fmt.Errorf("coldn't create place_categories: %w", err)
		}
		if _, err = result.RowsAffected(); err != nil {
			return fmt.Errorf("couldn't get number of rows affected: %w", err)
		}
	}
	return nil
}

func (r *PlaceRepository) DeletePlace(ctx context.Context, id uint) error {
	query := "DELETE FROM places WHERE id=$1"
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("couldn't delete place: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("couldn't get number of rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted: %w", models.ErrNotFound)
	}
	return nil
}

func (r *PlaceRepository) SearchPlaces(ctx context.Context, name string, limit, offset int) ([]models.GetPlace, error) {
	var places []models.GetPlace
	query := "SELECT p.id, p.name, p.imagePath, p.description, p.rating, p.numberOfReviews, p.address, p.phoneNumber, c.name AS city_name, ARRAY_AGG(ca.name) AS categories FROM places p JOIN cities c ON p.cityId = c.id JOIN places_categories pc ON p.id = pc.place_id JOIN categories ca ON pc.category_id = ca.id WHERE p.name LIKE '%' || $1 || '%' GROUP BY p.id, c.name ORDER BY p.id LIMIT $2 OFFSET $3"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("couldn't prepare query: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, name, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("couldn't get places: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var place models.GetPlace
		err := rows.Scan(&place.ID, &place.Name, &place.ImagePath, &place.Description, &place.Rating, &place.NumberOfReviews, &place.Address, &place.PhoneNumber, &place.City, pq.Array(&place.Categories))
		if err != nil {
			return nil, fmt.Errorf("couldn't unmarshal list of places: %w", err)
		}
		places = append(places, place)
	}
	return places, nil
}
