package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"log"
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
	log.Println("debug trip update:", trip.ID)

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
GROUP BY t.id, t.user_id, t.name, t.description, t.city_id, 
         t.start_date, t.end_date, t.private, t.created_at

UNION

SELECT 
    t.id, t.user_id, t.name, t.description, t.city_id, 
    t.start_date, t.end_date, t.private, t.created_at, 
    COALESCE(ARRAY_AGG(tp.photo_path) FILTER (WHERE tp.photo_path IS NOT NULL), '{}') AS photos
FROM trip t
LEFT JOIN trip_photo tp ON t.id = tp.trip_id
WHERE t.id IN (
    SELECT trip_id FROM user_shared_trip WHERE user_id = $1
)
GROUP BY t.id, t.user_id, t.name, t.description, t.city_id, 
         t.start_date, t.end_date, t.private, t.created_at

ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

`

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
	// sharedTripsQuery := `SELECT trip_id FROM user_shared_trip WHERE user_id = $1`
	// tripIds, err := r.db.QueryContext(ctx, sharedTripsQuery, userID)
	// for tripIds.Next() {
	// 	findTripQuery := `SELECT `
	// }
	return trips, nil
}

func (r *TripRepository) GetTrip(ctx context.Context, tripID uint) (models.Trip, []models.UserProfile, error) {
	var trip models.Trip

	// Получение информации о поездке
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
			return trip, nil, models.ErrNotFound
		}
		return trip, nil, fmt.Errorf("failed to get trip: %w", err)
	}

	// Получение фотографий поездки
	photoQuery := `
        SELECT photo_path 
        FROM trip_photo
        WHERE trip_id = $1
    `
	rows, err := r.db.QueryContext(ctx, photoQuery, tripID)
	if err != nil {
		return trip, nil, fmt.Errorf("failed to get trip photos: %w", err)
	}
	defer rows.Close()

	var photos []string
	for rows.Next() {
		var photoPath string
		if err := rows.Scan(&photoPath); err != nil {
			return trip, nil, fmt.Errorf("failed to scan photo: %w", err)
		}
		photos = append(photos, photoPath)
	}

	trip.Photos = photos

	// Получение списка user_id, которые участвуют в поездке
	userIDQuery := `SELECT user_id FROM user_shared_trip WHERE trip_id = $1`
	rows, err = r.db.QueryContext(ctx, userIDQuery, tripID)
	if err != nil {
		return trip, nil, fmt.Errorf("failed to get user ids for trip: %w", err)
	}
	defer rows.Close()

	var userIDs []uint
	for rows.Next() {
		var userID uint
		if err := rows.Scan(&userID); err != nil {
			return trip, nil, fmt.Errorf("failed to scan user_id: %w", err)
		}
		userIDs = append(userIDs, userID)
	}

	// Для каждого user_id получить информацию о пользователе
	var userProfiles []models.UserProfile
	for _, userID := range userIDs {
		userQuery := `SELECT login, avatar_path, email FROM "user" WHERE id = $1`
		var userProfile models.UserProfile
		err := r.db.QueryRowContext(ctx, userQuery, userID).Scan(
			&userProfile.Login,
			&userProfile.AvatarPath,
			&userProfile.Email,
		)
		if err != nil {
			return trip, nil, fmt.Errorf("failed to get user profile for user_id %d: %w", userID, err)
		}
		userProfiles = append(userProfiles, userProfile)
	}

	return trip, userProfiles, nil
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

func (r *TripRepository) CreateSharingLink(ctx context.Context, tripID uint, token string, sharingOption string) error {
	query := `
        INSERT INTO sharing_token (trip_id, token, sharing_option, expires_at)
        VALUES ($1, $2, $3, NOW() + INTERVAL '7 days')
    `
	_, err := r.db.ExecContext(ctx, query, tripID, token, sharingOption)
	if err != nil {
		return fmt.Errorf("failed to insert sharing link into database: %w", err)
	}
	return nil
}

func (r *TripRepository) GetSharingToken(ctx context.Context, tripID uint) (models.SharingToken, error) {
	query := `SELECT token, sharing_option, expires_at FROM sharing_token WHERE trip_id = $1`
	var token models.SharingToken
	err := r.db.QueryRowContext(ctx, query, tripID).Scan(
		&token.Token,
		&token.SharingOption,
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

func (r *TripRepository) GetTripBySharingToken(ctx context.Context, token string) (models.Trip, []models.UserProfile, error) {
	tripIdQuery := `SELECT trip_id FROM sharing_token WHERE token = $1`
	var tripID int
	err := r.db.QueryRowContext(ctx, tripIdQuery, token).Scan(
		&tripID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Trip{}, nil, models.ErrNotFound
		}
		return models.Trip{}, nil, fmt.Errorf("failed to retrive trip ID by sharing token: %w", err)
	}

	query := `SELECT id, user_id, name, description, city_id, start_date, end_date, private, created_at FROM trip WHERE id = $1`
	var trip models.Trip
	err = r.db.QueryRowContext(ctx, query, tripID).Scan(
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
			return trip, nil, models.ErrNotFound
		}
		return trip, nil, fmt.Errorf("failed to get trip: %w", err)
	}

	userIDQuery := `SELECT user_id FROM user_shared_trip WHERE trip_id = $1`
	usersRows, err := r.db.QueryContext(ctx, userIDQuery, tripID)
	if err != nil {
		return trip, nil, fmt.Errorf("failed to get user ids for trip: %w", err)
	}
	defer usersRows.Close()

	var userIDs []uint
	for usersRows.Next() {
		var userID uint
		if err := usersRows.Scan(&userID); err != nil {
			return trip, nil, fmt.Errorf("failed to scan user_id: %w", err)
		}
		userIDs = append(userIDs, userID)
	}

	// Для каждого user_id получить информацию о пользователе
	var userProfiles []models.UserProfile
	for _, userID := range userIDs {
		userQuery := `SELECT login, avatar_path, email FROM "user" WHERE id = $1`
		var userProfile models.UserProfile
		err := r.db.QueryRowContext(ctx, userQuery, userID).Scan(
			&userProfile.Login,
			&userProfile.AvatarPath,
			&userProfile.Email,
		)
		if err != nil {
			return trip, nil, fmt.Errorf("failed to get user profile for user_id %d: %w", userID, err)
		}
		userProfiles = append(userProfiles, userProfile)
	}

	return trip, userProfiles, nil
}

func (r *TripRepository) AddUserToTrip(ctx context.Context, tripId, userId uint) (bool, error) {
	findOptionQuery := `SELECT sharing_option FROM sharing_token WHERE trip_id = $1`
	var sharingOption string
	err := r.db.QueryRowContext(ctx, findOptionQuery, tripId).Scan(&sharingOption)
	if err != nil {
		return false, fmt.Errorf("failed to find sharing option: %w", err)
	}
	findUserQuery := `SELECT COUNT(*) FROM user_shared_trip WHERE user_id = $1 AND trip_id = $2 and sharing_option = $3`
	var count int
	err = r.db.QueryRowContext(ctx, findUserQuery, userId, tripId, sharingOption).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to find shared trip for this user: %w", err)
	}
	if count > 0 {
		return false, nil
	}

	query := `INSERT INTO user_shared_trip (trip_id, user_id, sharing_option, created_at)
        VALUES ($1, $2, $3, NOW())`
	_, err = r.db.ExecContext(ctx, query, tripId, userId, sharingOption)
	if err != nil {
		return false, fmt.Errorf("failed to add user to a trip: %w", err)
	}
	return true, nil
}

func (r *TripRepository) GetSharingOption(ctx context.Context, userId, tripId uint) (string, error) {
	tripQuery := `SELECT user_id FROM trip WHERE id = $1`
	var owner int
	var sharingOption string
	err := r.db.QueryRowContext(ctx, tripQuery, tripId).Scan(&owner)
	if err != nil {
		return "", fmt.Errorf("failed to find sharing option: %w", err)
	}
	if owner == int(userId) {
		sharingOption = "editing"
		return sharingOption, nil
	}
	query := `SELECT sharing_option FROM user_shared_trip WHERE user_id = $1 AND trip_id = $2`
	err = r.db.QueryRowContext(ctx, query, userId, tripId).Scan(&sharingOption)
	if err != nil {
		return "", fmt.Errorf("failed to find sharing option: %w", err)
	}
	return sharingOption, nil
}
