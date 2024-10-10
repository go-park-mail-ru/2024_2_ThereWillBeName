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

func (r *PlaceRepository) GetPlaces(ctx context.Context) ([]models.Place, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, imagePath, description FROM places")
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
	result, err := r.db.ExecContext(ctx, "INSERT INTO places (name, imagePath, description) VALUES ($1, $2, $3)", place.Name, place.ImagePath, place.Description)
	if err != nil {
		return fmt.Errorf("coldn't create place: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("couldn't get number of rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created")
	}
	return nil
}

func (r *PlaceRepository) GetPlace(ctx context.Context, name string) (models.Place, error) {
	var dataPlace models.Place
	row := r.db.QueryRowContext(ctx, "SELECT id, name, imagePath, description FROM places where name = $1", name)
	err := row.Scan(&dataPlace.ID, &dataPlace.Name, &dataPlace.ImagePath, &dataPlace.Description)
	if err != nil {
		return models.Place{}, fmt.Errorf("couldn't get place by name: %w", err)
	}
	return dataPlace, nil
}

func (r *PlaceRepository) UpdatePlace(ctx context.Context, place models.Place) error {
	result, err := r.db.ExecContext(ctx, "UPDATE places SET name = $1, imagePath = $2, description = $3", place.Name, place.ImagePath, place.Description)
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

func (r *PlaceRepository) DeletePlace(ctx context.Context, name string) error {
	result, err := r.db.ExecContext(ctx, "DELETE FROM places WHERE name=$1", name)
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

func (r *PlaceRepository) SearchPlaces(ctx context.Context, name string) ([]models.Place, error) {
	var places []models.Place
	stmt, err := r.db.PrepareContext(ctx, "SELECT id, name, imagePath, description FROM places where name LIKE $1")
	if err != nil {
		return nil, fmt.Errorf("couldn't prepare query: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, "%"+name+"%")
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
