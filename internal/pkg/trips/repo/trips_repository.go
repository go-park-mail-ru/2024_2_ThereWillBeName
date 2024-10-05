package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type TripRepository struct {
	db *sql.DB
}

func NewTripRepository(db *sql.DB) *TripRepository {
	return &TripRepository{db: db}
}

func (r *TripRepository) CreateTrip(ctx context.Context, trip models.Trip) error {
	query := `INSERT INTO trips (user_id, name, description, destination, start_date, end_date, private, created_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())`

	_, err := r.db.ExecContext(ctx, query, trip.UserID, trip.Name, trip.Description, trip.Destination, trip.StartDate, trip.EndDate, trip.Private)
	return err
}

func (r *TripRepository) UpdateTrip(ctx context.Context, trip models.Trip) error {
	query := `UPDATE trips 
              SET name = $1, description = $2, destination = $3, start_date = $4, end_date = $5, private = $6 
              WHERE id = $7`

	_, err := r.db.ExecContext(ctx, query, trip.Name, trip.Description, trip.Destination, trip.StartDate, trip.EndDate, trip.Private, trip.ID)
	return err
}

func (r *TripRepository) DeleteTrip(ctx context.Context, id uint) error {
	query := `DELETE FROM trips WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *TripRepository) ReadTripsByUserID(ctx context.Context, userID uint) ([]models.Trip, error) {
	query := `SELECT id, user_id, name, description, destination, start_date, end_date, private, created_at 
              FROM trips 
              WHERE user_id = $1
              ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trips []models.Trip
	for rows.Next() {
		var trip models.Trip
		if err := rows.Scan(&trip.ID, &trip.UserID, &trip.Name, &trip.Description, &trip.Destination, &trip.StartDate, &trip.EndDate, &trip.Private, &trip.CreatedAt); err != nil {
			return nil, err
		}
		trips = append(trips, trip)
	}

	return trips, nil
}

func (r *TripRepository) ReadTrip(ctx context.Context, tripID uint) (models.Trip, error) {
	query := `SELECT id, user_id, name, description, destination, start_date, end_date, private, created_at 
              FROM trips 
              WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, tripID)

	var trip models.Trip
	err := row.Scan(&trip.ID, &trip.UserID, &trip.Name, &trip.Description, &trip.Destination, &trip.StartDate, &trip.EndDate, &trip.Private, &trip.CreatedAt)
	if err != nil {
		return models.Trip{}, err
	}

	return trip, nil
}
