package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"time"

	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type TripRepository struct {
	db *dblogger.DB
}

func NewTripRepository(db *dblogger.DB) *TripRepository {
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
	query := `
		SELECT 
			t.id, t.user_id, t.name, t.description, t.city_id, 
			t.start_date, t.end_date, t.private, t.created_at, 
			COALESCE(ARRAY_AGG(tp.photo_path) FILTER (WHERE tp.photo_path IS NOT NULL), '{}') AS photos
		FROM trip t
		LEFT JOIN trip_photo tp ON t.id = tp.trip_id
		WHERE t.user_id = $1
		GROUP BY t.id
		ORDER BY t.created_at DESC
		LIMIT $2 OFFSET $3`

	rows, err := r.db.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve trips: %w", models.ErrInternal)
	}
	defer rows.Close()

	var trips []models.Trip
	for rows.Next() {
		var trip models.Trip
		if err := rows.Scan(
			&trip.ID, &trip.UserID, &trip.Name, &trip.Description,
			&trip.CityID, &trip.StartDate, &trip.EndDate, &trip.Private,
			&trip.CreatedAt, pq.Array(&trip.Photos),
		); err != nil {
			return nil, fmt.Errorf("failed to scan trip row: %w", models.ErrInternal)
		}

		trips = append(trips, trip)
	}

	if len(trips) == 0 {
		return nil, fmt.Errorf("no trips found: %w", models.ErrNotFound)
	}

	return trips, nil
}

func (r *TripRepository) GetTrip(ctx context.Context, tripID uint) (models.Trip, error) {
	var trip models.Trip

	query := `
        SELECT 
            id, user_id, name, description, city_id, start_date, end_date, private, created_at 
        FROM 
            trip
        WHERE 
            id = $1
    `
	err := r.db.QueryRowContext(ctx, query, tripID).Scan(
		&trip.ID,
		&trip.UserID,
		&trip.Name,
		&trip.Description,
		&trip.CityID,
		&trip.StartDate,
		&trip.EndDate,
		&trip.Private,
		&trip.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return trip, models.ErrNotFound
		}
		return trip, fmt.Errorf("failed to get trip: %w", err)
	}

	photoQuery := `
        SELECT photo_path 
        FROM trip_photo
        WHERE trip_id = $1
    `
	rows, err := r.db.QueryContext(ctx, photoQuery, tripID)
	if err != nil {
		return trip, fmt.Errorf("failed to get trip photos: %w", err)
	}
	defer rows.Close()

	var photos []string
	for rows.Next() {
		var photoPath string
		if err := rows.Scan(&photoPath); err != nil {
			return trip, fmt.Errorf("failed to scan photo: %w", err)
		}
		photos = append(photos, photoPath)
	}

	trip.Photos = photos

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

func (r *TripRepository) AddPhotoToTrip(ctx context.Context, tripID uint, photoPath string) error {
	query := `
        INSERT INTO trip_photo (trip_id, photo_path)
        VALUES ($1, $2)
    `
	_, err := r.db.ExecContext(ctx, query, tripID, photoPath)
	if err != nil {
		return fmt.Errorf("failed to insert photo into database: %w", err)
	}
	return nil
}

func (r *TripRepository) DeletePhotoFromTrip(ctx context.Context, tripID uint, photoPath string) error {
	query := `DELETE FROM trip_photo WHERE trip_id = $1 AND photo_path = $2`
	result, err := r.db.ExecContext(ctx, query, tripID, photoPath)
	if err != nil {
		return fmt.Errorf("failed to delete photo from database: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("photo not found in trip: %w", models.ErrNotFound)
	}

	return nil
}

func (r *TripRepository) CreateSharingLink(ctx context.Context, tripID uint, token string) error {
	query := `
        INSERT INTO shared_link (trip_id, token, expires_at)
        VALUES ($1, $2, NOW() + INTERVAL '7 days')
    `
	_, err := r.db.ExecContext(ctx, query, tripID, token)
	if err != nil {
		return fmt.Errorf("failed to insert sharing link into database: %w", err)
	}
	return nil
}

func (r *TripRepository) GetSharingToken(ctx context.Context, tripID uint) (models.SharingToken, error) {
	query := `SELECT token, expires_at FROM sharing_link WHERE trip_id = $1`
	var token models.SharingToken
	err := r.db.QueryRowContext(ctx, query, tripID).Scan(
		&token.Token,
		&token.ExpiresAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.SharingToken{}, nil
		}
		return models.SharingToken{}, fmt.Errorf("failed to retrive sharing token: %w", err)
	}
	if token.ExpiresAt.Before(time.Now()) {
		deleteQuery := `DELETE from shared_link WHERE trip_id = $1`
		result, err := r.db.ExecContext(ctx, deleteQuery, tripID)
		if err != nil {
			return models.SharingToken{}, fmt.Errorf("failed to delete sharing token: %w", models.ErrInternal)
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return models.SharingToken{}, fmt.Errorf("failed to retrieve rows affected %w", models.ErrInternal)
		}
		if rowsAffected == 0 {
			return models.SharingToken{}, fmt.Errorf("no rows were deleted: %w", models.ErrNotFound)
		}
	}
	return token, nil
}
