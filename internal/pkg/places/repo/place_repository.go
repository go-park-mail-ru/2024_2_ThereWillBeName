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
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, image, description FROM places")
	if err != nil {
		return nil, fmt.Errorf("couldn't get places: %w", err)
	}
	defer rows.Close()
	var places []models.Place
	for rows.Next() {
		var place models.Place
		err := rows.Scan(&place.ID, &place.Name, &place.Image, &place.Description)
		if err != nil {
			return nil, fmt.Errorf("не получилось распарсить в place: %w", err)
		}
		places = append(places, place)
	}
	return places, nil
}

func (r *PlaceRepository) CreatePlace(ctx context.Context, place models.Place) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO places (name, image, description) VALUES ($1, $2, $3)", place.Name, place.Image, place.Description)
	if err != nil {
		return fmt.Errorf("coldn't create place: %w", err)
	}
	return nil
}

func (r *PlaceRepository) GetPlace(ctx context.Context, name string) (models.Place, error) {
	var dataPlace models.Place
	row := r.db.QueryRowContext(ctx, "SELECT id, name, image, description FROM places where name = $1", name)
	err := row.Scan(&dataPlace.ID, &dataPlace.Name, &dataPlace.Image, &dataPlace.Description)
	if err != nil {
		return models.Place{}, fmt.Errorf("couldn't get place by name: %w", err)
	}
	return dataPlace, nil
}

func (r *PlaceRepository) UpdatePlace(ctx context.Context, place models.Place) error {
	_, err := r.db.ExecContext(ctx, "UPDATE places SET name = $1, image = $2, description = $3", place.Name, place.Image, place.Description)
	if err != nil {
		return fmt.Errorf("couldn't update place: %w", err)
	}
	return nil
}

func (r *PlaceRepository) DeletePlace(ctx context.Context, name string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM places WHERE name=$1", name)
	if err != nil {
		return fmt.Errorf("couldn't delete place: %w", err)
	}
	return nil
}

func (r *PlaceRepository) GetPlacesBySearch(ctx context.Context, name string) ([]models.Place, error) {
	var places []models.Place
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, image, description FROM places where name LIKE $1", "%"+name+"%")
	if err != nil {
		return nil, fmt.Errorf("couldn't get places: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var place models.Place
		err := rows.Scan(&place.ID, &place.Name, &place.Image, &place.Description)
		if err != nil {
			return nil, fmt.Errorf("не получилось распарсить в place: %w", err)
		}
		places = append(places, place)
	}
	return places, nil
}
