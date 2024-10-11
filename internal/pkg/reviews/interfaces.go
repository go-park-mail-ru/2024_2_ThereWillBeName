package reviews

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go
type ReviewsUsecase interface {
	CreateReview(ctx context.Context, review models.Review) error
	UpdateReview(ctx context.Context, review models.Review) error
	DeleteReview(ctx context.Context, reviewID uint) error
	GetReviewsByPlaceID(ctx context.Context, placeID uint) ([]models.Review, error)
	GetReview(ctx context.Context, reviewID uint) (models.Review, error)
}

type ReviewsRepo interface {
	CreateReview(ctx context.Context, review models.Review) error
	UpdateReview(ctx context.Context, review models.Review) error
	DeleteReview(ctx context.Context, reviewID uint) error
	GetReviewsByPlaceID(ctx context.Context, placeID uint) ([]models.Review, error)
	GetReview(ctx context.Context, reviewID uint) (models.Review, error)
}
