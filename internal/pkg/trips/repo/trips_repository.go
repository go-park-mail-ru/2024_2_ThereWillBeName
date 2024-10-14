package repo

import (
	"2024_2_ThereWillBeName/internal/models"

	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
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
		log.Println(err)
		return fmt.Errorf("failed to create a trip: %w", models.ErrInternal.CustomError)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to retrieve rows affected: %w", models.ErrInternal.CustomError)
	}
	if rowsAffected == 0 {
		log.Println("no rows were created")
		return fmt.Errorf("no rows were created: %w", models.ErrNotFound.CustomError)
	}

	return nil
}

func (r *TripRepository) UpdateTrip(ctx context.Context, trip models.Trip) error {
	query := `UPDATE trips 
              SET name = $1, description = $2, city_id = $3, start_date = $4, end_date = $5, private = $6 
              WHERE id = $7`

	result, err := r.db.ExecContext(ctx, query, trip.Name, trip.Description, trip.CityID, trip.StartDate, trip.EndDate, trip.Private, trip.ID)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to execute update query: %w", models.ErrInternal.CustomError)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to retrieve rows affected: %w", models.ErrInternal.CustomError)
	}
	if rowsAffected == 0 {
		log.Println("no rows were updated")
		return fmt.Errorf("no rows were updated: %w", models.ErrNotFound.CustomError)
	}

	return nil
}

func (r *TripRepository) DeleteTrip(ctx context.Context, id uint) error {
	query := `DELETE FROM trips WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to delete trip: %w", models.ErrInternal.CustomError)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return fmt.Errorf("failed to retrieve rows affected %w", err)
	}
	if rowsAffected == 0 {
		log.Println("no rows were deleted")
		return fmt.Errorf("no rows were deleted: %w", models.ErrNotFound.CustomError)
	}
	return nil
}

func (r *TripRepository) GetTripsByUserID(ctx context.Context, userID uint) ([]models.Trip, error) {
	var exists bool
	checkUserQuery := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`
	err := r.db.QueryRowContext(ctx, checkUserQuery, userID).Scan(&exists)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to find user: %w", models.ErrInternal.CustomError)
	}

	if !exists {
		log.Println("user not found")
		return nil, fmt.Errorf("user not found: %w", models.ErrUserNotFound.CustomError)
	}

	query := `SELECT id, user_id, name, description, city_id, start_date, end_date, private, created_at 
              FROM trips 
              WHERE user_id = $1
              ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, userID)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to retrieve trips: %w", models.ErrInternal.CustomError)
	}
	defer rows.Close()

	var tripRows []models.Trip
	for rows.Next() {
		var trip models.Trip
		if err := rows.Scan(&trip.ID, &trip.UserID, &trip.Name, &trip.Description, &trip.CityID, &trip.StartDate, &trip.EndDate, &trip.Private, &trip.CreatedAt); err != nil {
			log.Println(err)
			return nil, fmt.Errorf("failed to scan trip row: %w", models.ErrInternal.CustomError)
		}
		tripRows = append(tripRows, trip)
	}

	if len(tripRows) == 0 {
		log.Println("no trips were found")
		return nil, fmt.Errorf("no trips found: %w", models.ErrNotFound.CustomError)
	}

	return tripRows, nil
}

func (r *TripRepository) GetTrip(ctx context.Context, tripID uint) (models.Trip, error) {
	query := `SELECT id, user_id, name, description, city_id, start_date, end_date, private, created_at 
              FROM trips 
              WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, tripID)

	var trip models.Trip
	err := row.Scan(&trip.ID, &trip.UserID, &trip.Name, &trip.Description, &trip.CityID, &trip.StartDate, &trip.EndDate, &trip.Private, &trip.CreatedAt)
	if err != nil {
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return models.Trip{}, fmt.Errorf("trip not found: %w", models.ErrNotFound.CustomError)
		}
		return models.Trip{}, fmt.Errorf("failed to scan trip row: %w", models.ErrInternal.CustomError)
	}

	return trip, nil
}
