package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

var (
	ErrNotFound = errors.New("trip not found")
	ErrConflict = errors.New("foreign key constraint violation")
	ErrInternal = errors.New("internal repository error")
	//ErrForeignKeyViolation = errors.New("foreign key constraint violation")
	ErrUserNotFound = errors.New("user not found")
)

type TripRepository struct {
	db *sql.DB
}

func NewTripRepository(db *sql.DB) *TripRepository {
	return &TripRepository{db: db}
}

func (r *TripRepository) CreateTrip(ctx context.Context, trip models.Trip) error {
	query := `INSERT INTO trips (user_id, name, description, city_id, start_date, end_date, private, created_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())`

	result, err := r.db.ExecContext(ctx, query, trip.UserID, trip.Name, trip.Description, trip.CityID, trip.StartDate, trip.EndDate, trip.Private)
	if err != nil {
		// Проверяем на нарушение внешнего ключа
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			return fmt.Errorf("failed to create a trip: %w", ErrConflict)
		}
		return fmt.Errorf("failed to create a trip: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were created: %w", ErrNotFound)
	}

	return nil
}

func (r *TripRepository) UpdateTrip(ctx context.Context, trip models.Trip) error {
	query := `UPDATE trips 
              SET name = $1, description = $2, city_id = $3, start_date = $4, end_date = $5, private = $6 
              WHERE id = $7`

	result, err := r.db.ExecContext(ctx, query, trip.Name, trip.Description, trip.CityID, trip.StartDate, trip.EndDate, trip.Private, trip.ID)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated: %w", ErrNotFound)
	}

	return nil
}

func (r *TripRepository) DeleteTrip(ctx context.Context, id uint) error {
	query := `DELETE FROM trips WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			return fmt.Errorf("failed to delete trip: %w", ErrConflict)
		}
		return fmt.Errorf("failed to delete trip: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve rows affected %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted: %w", ErrNotFound)
	}
	return nil
}

func (r *TripRepository) GetTripsByUserID(ctx context.Context, userID uint) ([]models.Trip, error) {
	query := `SELECT id, user_id, name, description, city_id, start_date, end_date, private, created_at 
              FROM trips 
              WHERE user_id = $1
              ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		// Проверяем на нарушение внешнего ключа
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23503" {
			return nil, fmt.Errorf("failed to retrieve trips: %w", ErrUserNotFound)
		}
		return nil, fmt.Errorf("failed to retrieve trips: %w", err)
	}
	defer rows.Close()

	var trips []models.Trip
	for rows.Next() {
		var trip models.Trip
		if err := rows.Scan(&trip.ID, &trip.UserID, &trip.Name, &trip.Description, &trip.CityID, &trip.StartDate, &trip.EndDate, &trip.Private, &trip.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan trip row: %w", err)
		}
		trips = append(trips, trip)
	}

	if len(trips) == 0 {
		return nil, fmt.Errorf("no trips found: %w", ErrNotFound)
	}

	return trips, nil
}

func (r *TripRepository) GetTrip(ctx context.Context, tripID uint) (models.Trip, error) {
	query := `SELECT id, user_id, name, description, city_id, start_date, end_date, private, created_at 
              FROM trips 
              WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, tripID)

	var trip models.Trip
	err := row.Scan(&trip.ID, &trip.UserID, &trip.Name, &trip.Description, &trip.CityID, &trip.StartDate, &trip.EndDate, &trip.Private, &trip.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Trip{}, fmt.Errorf("trip not found: %w", ErrNotFound)
		}
		return models.Trip{}, fmt.Errorf("failed to scan trip row: %w", err)
	}

	return trip, nil
}
