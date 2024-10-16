package reviews

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type ReviewRepository struct {
	db *sql.DB
}

func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) CreateReview(ctx context.Context, review models.Review) error {
	query := `INSERT INTO reviews (user_id, place_id, rating, review_text, created_at) 
              VALUES ($1, $2, $3, $4, NOW())`

	result, err := r.db.ExecContext(ctx, query, review.UserID, review.PlaceID, review.Rating, review.ReviewText)
	if err != nil {
		log.Println("failed to execute insert review query:", err)
		return fmt.Errorf("failed to create review: %w", models.ErrInternal.CustomError)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("failed to retrieve rows affected:", err)
		return fmt.Errorf("failed to retrieve rows affected: %w", models.ErrInternal.CustomError)
	}

	if rowsAffected == 0 {
		log.Println("no rows were created for the review")
		return fmt.Errorf("no rows were created for the review: %w", models.ErrNotFound.CustomError)
	}
	return nil
}

func (r *ReviewRepository) UpdateReview(ctx context.Context, review models.Review) error {
	query := `UPDATE reviews 
              SET rating = $1, review_text = $2 
              WHERE id = $3`

	result, err := r.db.ExecContext(ctx, query, review.Rating, review.ReviewText, review.ID)
	if err != nil {
		log.Println("failed to execute update query:", err)
		return fmt.Errorf("failed to update review: %w", models.ErrInternal.CustomError)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("failed to retrieve rows affected:", err)
		return fmt.Errorf("failed to retrieve rows affected: %w", models.ErrInternal.CustomError)
	}

	if rowsAffected == 0 {
		log.Println("no rows were updated for review ID:", review.ID)
		return fmt.Errorf("no rows were updated: %w", models.ErrNotFound.CustomError)
	}

	return nil
}

func (r *ReviewRepository) DeleteReview(ctx context.Context, reviewID uint) error {
	query := `DELETE FROM reviews WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, reviewID)
	if err != nil {
		log.Println("failed to execute delete query:", err)
		return fmt.Errorf("failed to delete review: %w", models.ErrInternal.CustomError)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("failed to retrieve rows affected:", err)
		return fmt.Errorf("failed to retrieve rows affected: %w", models.ErrInternal.CustomError)
	}

	if rowsAffected == 0 {
		log.Println("No rows were deleted for review ID:", reviewID)
		return fmt.Errorf("review not found: %w", models.ErrNotFound.CustomError)
	}

	return nil
}

func (r *ReviewRepository) GetReviewsByPlaceID(ctx context.Context, placeID uint, limit, offset int) ([]models.Review, error) {
	query := `SELECT id, user_id, place_id, rating, review_text, created_at 
              FROM reviews 
              WHERE place_id = $1
              ORDER BY created_at DESC
			  LIMIT $2 OFFSET $3`

	rows, err := r.db.QueryContext(ctx, query, placeID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve reviews: %w", models.ErrInternal.CustomError)
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ID, &review.UserID, &review.PlaceID, &review.Rating, &review.ReviewText, &review.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan review row: %w", models.ErrInternal.CustomError)
		}
		reviews = append(reviews, review)
	}

	if len(reviews) == 0 {
		log.Println("no reviews found for the given place")
		return nil, fmt.Errorf("no reviews found for place with ID %d: %w", placeID, models.ErrNotFound.CustomError)
	}

	return reviews, nil
}

func (r *ReviewRepository) GetReview(ctx context.Context, reviewID uint) (models.Review, error) {
	query := `SELECT id, user_id, place_id, rating, review_text, created_at 
              FROM reviews 
              WHERE id = $1`

	row := r.db.QueryRowContext(ctx, query, reviewID)

	var review models.Review
	err := row.Scan(&review.ID, &review.UserID, &review.PlaceID, &review.Rating, &review.ReviewText, &review.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Review{}, fmt.Errorf("review with ID %d did not found: %w", reviewID, models.ErrNotFound.CustomError)
		}
		return models.Review{}, fmt.Errorf("failed to scan review: %w", models.ErrInternal.CustomError)
	}

	return review, nil
}
