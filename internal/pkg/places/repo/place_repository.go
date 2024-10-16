package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	_ "embed"
	"fmt"
)

type PlaceRepository struct {
	db *sql.DB
}

func NewPLaceRepository(db *sql.DB) *PlaceRepository {
	return &PlaceRepository{db: db}
}

func (r *PlaceRepository) GetPlaces(ctx context.Context, limit, offset int) ([]models.Place, error) {
	query := "SELECT id, name, imagePath, description FROM places ORDER BY id LIMIT $1 OFFSET $2"
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("couldn't get places: %w", err)
	}
	defer rows.Close()
	var places []models.Place
	for rows.Next() {
		var place models.Place
		err := rows.Scan(&place.ID, &place.Name, &place.ImagePath, &place.Description)
		if err != nil {
			return nil, fmt.Errorf("couldn't unmarshal list of places: %w", err)
		}
		places = append(places, place)
	}
	return places, nil
}

func (r *PlaceRepository) CreatePlace(ctx context.Context, place models.Place) error {
	query := "INSERT INTO places (name, imagePath, description, rating, numberOfReviews, address, city, phoneNumber, category) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	result, err := r.db.ExecContext(ctx, query, place.Name, place.ImagePath, place.Description, place.Rating, place.NumberOfReviews, place.Address, place.City, place.PhoneNumber, place.Category)
	if err != nil {
		return fmt.Errorf("coldn't create place: %w", err)
	}
	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("couldn't get number of rows affected: %w", err)
	}
	return nil
}

func (r *PlaceRepository) GetPlace(ctx context.Context, id int) (models.Place, error) {
	var dataPlace models.Place
	query := "SELECT id, name, imagePath, description FROM places where id = $1"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&dataPlace.ID, &dataPlace.Name, &dataPlace.ImagePath, &dataPlace.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Place{}, fmt.Errorf("place not found")
		}
		return models.Place{}, fmt.Errorf("couldn't get place by name: %w", err)
	}
	return dataPlace, nil
}

func (r *PlaceRepository) UpdatePlace(ctx context.Context, place models.Place) error {
	query := "UPDATE places SET name = $1, imagePath = $2, description = $3 WHERE id=$4"
	result, err := r.db.ExecContext(ctx, query, place.Name, place.ImagePath, place.Description, place.ID)
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
	return nil
}

func (r *PlaceRepository) DeletePlace(ctx context.Context, id int) error {
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
		return fmt.Errorf("no rows were deleted")
	}
	return nil
}

func (r *PlaceRepository) SearchPlaces(ctx context.Context, name string, limit, offset int) ([]models.Place, error) {
	var places []models.Place
	query := "SELECT id, name, imagePath, description FROM places where name LIKE '%' || $1 || '%' ORDER BY id LIMIT $2 OFFSET $3"
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
		var place models.Place
		err := rows.Scan(&place.ID, &place.Name, &place.ImagePath, &place.Description)
		if err != nil {
			return nil, fmt.Errorf("couldn't unmarshal list of places: %w", err)
		}
		places = append(places, place)
	}
	return places, nil
}
