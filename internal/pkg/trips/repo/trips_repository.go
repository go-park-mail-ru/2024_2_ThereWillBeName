package repo

import (
	"2024_2_ThereWillBeName/internal/models"

	"context"
	"database/sql"
	"errors"
	"fmt"
)

type TripRepository struct {
	db *sql.DB
}

func NewTripRepository(db *sql.DB) *TripRepository {
	return &TripRepository{db: db}
}

func (r *TripRepository) CreateTrip(ctx context.Context, trip models.Trip) error {
	query := `INSERT INTO trip (user_id, name, description, city_id, start_date, end_date, private, created_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())`

	result, err := r.db.ExecContext(ctx, query, trip.UserID, trip.Name, trip.Description, trip.CityID, trip.StartDate, trip.EndDate, trip.Private)
	if err != nil {
		return fmt.Errorf("failed to create a trip: %w", models.ErrInternal)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %w", models.ErrInternal)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created: %w", models.ErrNotFound)
	}

	return nil
}

func (r *TripRepository) UpdateTrip(ctx context.Context, trip models.Trip) error {
	query := `UPDATE trip 
              SET name = $1, description = $2, city_id = $3, start_date = $4, end_date = $5, private = $6, updated_at = NOW() 
              WHERE id = $7`

	result, err := r.db.ExecContext(ctx, query, trip.Name, trip.Description, trip.CityID, trip.StartDate, trip.EndDate, trip.Private, trip.ID)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %w", models.ErrInternal)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %w", models.ErrInternal)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated: %w", models.ErrNotFound)
	}

	return nil
}

func (r *TripRepository) DeleteTrip(ctx context.Context, id uint) error {
	query := `DELETE FROM trip WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete trip: %w", models.ErrInternal)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected %w", models.ErrInternal)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted: %w", models.ErrNotFound)
	}
	return nil
}

func (r *TripRepository) GetTripsByUserID(ctx context.Context, userID uint, limit, offset int) ([]models.Trip, error) {
	query := `SELECT id, user_id, name, description, city_id, start_date, end_date, private, created_at 
              FROM trip 
              WHERE user_id = $1
              ORDER BY created_at DESC
			  LIMIT $2 OFFSET $3`

	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve trips: %w", models.ErrInternal)
	}
	defer rows.Close()

	var tripRows []models.Trip
	for rows.Next() {
		var trip models.Trip
		if err := rows.Scan(&trip.ID, &trip.UserID, &trip.Name, &trip.Description, &trip.CityID, &trip.StartDate, &trip.EndDate, &trip.Private, &trip.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan trip row: %w", models.ErrInternal)
		}
		tripRows = append(tripRows, trip)
	}

	if len(tripRows) == 0 {
		return nil, fmt.Errorf("no trips found: %w", models.ErrNotFound)
	}

	return tripRows, nil
}

func (r *TripRepository) GetTrip(ctx context.Context, tripID uint) (models.Trip, error) {
	query := `SELECT id, user_id, name, description, city_id, start_date, end_date, private, created_at 
              FROM trip 
              WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, tripID)

	var trip models.Trip
	err := row.Scan(&trip.ID, &trip.UserID, &trip.Name, &trip.Description, &trip.CityID, &trip.StartDate, &trip.EndDate, &trip.Private, &trip.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Trip{}, fmt.Errorf("trip not found: %w", models.ErrNotFound)
		}
		return models.Trip{}, fmt.Errorf("failed to scan trip row: %w", models.ErrInternal)
	}

	return trip, nil
}

func (r *TripRepository) AddPlaceToTrip(ctx context.Context, tripID uint, placeID uint) error {
	query := `INSERT INTO trip_place (trip_id, place_id, created_at) 
              VALUES ($1, $2, NOW())`

	result, err := r.db.ExecContext(ctx, query, tripID, placeID)

	if err != nil {
		return fmt.Errorf("failed to add place to a trip: %w", models.ErrInternal)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %w", models.ErrInternal)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created: %w", models.ErrNotFound)
	}

	return nil
}
