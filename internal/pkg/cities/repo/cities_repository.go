package repo

import (
	"2024_2_ThereWillBeName/internal/models"

	"context"
	"database/sql"
	"errors"
	"fmt"
)

type CitiesRepository struct {
	db *sql.DB
}

func NewCitiesRepository(db *sql.DB) *CitiesRepository {
	return &CitiesRepository{db: db}
}

func (r *CitiesRepository) SearchCitiesByName(ctx context.Context, query string) ([]models.City, error) {
	var cities []models.City

	searchQuery := `SELECT id, name, created_at FROM cities WHERE name ILIKE $1`

	rows, err := r.db.QueryContext(ctx, searchQuery, query+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to search cities: %w", models.ErrInternal)
	}
	defer rows.Close()

	for rows.Next() {
		var city models.City
		err := rows.Scan(&city.ID, &city.Name, &city.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan city row: %w", models.ErrInternal)
		}
		cities = append(cities, city)
	}

	if len(cities) == 0 {
		return nil, models.ErrNotFound
	}

	return cities, nil
}

func (r *CitiesRepository) SearchCityByID(ctx context.Context, id uint) (models.City, error) {
	var city models.City

	query := `SELECT id, name, created_at FROM cities WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&city.ID, &city.Name, &city.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return city, models.ErrNotFound
		}
		return city, fmt.Errorf("failed to retrieve city: %w", models.ErrInternal)
	}

	return city, nil
}
