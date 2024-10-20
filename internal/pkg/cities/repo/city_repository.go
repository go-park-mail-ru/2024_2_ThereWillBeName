package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type CitiesRepository struct {
	db *sql.DB
}

func NewCitiesRepository(db *sql.DB) *CitiesRepository { return &CitiesRepository{db} }

func (c *CitiesRepository) CreateCity(ctx context.Context, city models.City) error {
	query := "INSERT INTO cities (name, created_at) VALUES ($1, NOW())"
	log.Println(query, city.Name)
	_, err := c.db.ExecContext(ctx, query, city.Name)
	return err
}

func (c *CitiesRepository) GetCity(ctx context.Context, id int) (models.City, error) {
	var city models.City
	query := "SELECT id, name FROM cities WHERE id = $1"
	row := c.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&city.ID, &city.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.City{}, fmt.Errorf("city not found: %w", models.ErrNotFound)
		}
		return models.City{}, fmt.Errorf("couldn't get place by name: %w", err)
	}
	return city, nil
}

func (c *CitiesRepository) UpdateCity(ctx context.Context, city models.City) error {
	query := "UPDATE cities SET name = $1, created_at = NOW() WHERE id = $2"
	result, err := c.db.ExecContext(ctx, query, city.Name, city.ID)
	if err != nil {
		return fmt.Errorf("couldn't update city: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("couldn't get number of rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated: %w", models.ErrNotFound)
	}
	return nil
}

func (c *CitiesRepository) DeleteCity(ctx context.Context, id int) error {
	query := "DELETE FROM cities WHERE id = $1"
	result, err := c.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("couldn't delete city: %w", err)
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

func (c *CitiesRepository) GetCities(ctx context.Context, limit, offset int) ([]models.City, error) {
	var cities []models.City
	query := "SELECT id, name, created_at FROM cities LIMIT $1 OFFSET $2"
	rows, err := c.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("couldn't get cities: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var city models.City
		log.Println(rows)
		err := rows.Scan(&city.ID, &city.Name, &city.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("couldn't unmarshal place by name: %w", err)
		}
		cities = append(cities, city)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("couldn't  unmarshal cities: %w", err)
	}
	log.Println(cities)
	return cities, nil
}
