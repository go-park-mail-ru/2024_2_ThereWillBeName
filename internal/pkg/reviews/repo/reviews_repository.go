package reviews

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"

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

	_, err := r.db.ExecContext(ctx, query, review.UserID, review.PlaceID, review.Rating, review.ReviewText)
	return err
}

func (r *ReviewRepository) UpdateReview(ctx context.Context, review models.Review) error {
	query := `UPDATE reviews 
              SET rating = $1, review_text = $2 
              WHERE id = $3`

	_, err := r.db.ExecContext(ctx, query, review.Rating, review.ReviewText, review.ID)
	return err
}

func (r *ReviewRepository) DeleteReview(ctx context.Context, reviewID uint) error {
	query := `DELETE FROM reviews WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, reviewID)
	return err
}

func (r *ReviewRepository) GetReviewsByPlaceID(ctx context.Context, placeID uint) ([]models.Review, error) {
	query := `SELECT id, user_id, place_id, rating, review_text, created_at 
              FROM reviews 
              WHERE place_id = $1
              ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, placeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var review models.Review
		if err := rows.Scan(&review.ID, &review.UserID, &review.PlaceID, &review.Rating, &review.ReviewText, &review.CreatedAt); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
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
		return models.Review{}, err
	}

	return review, nil
}
