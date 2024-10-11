package reviews

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/reviews"
	"context"
)

type reviewsUsecaseImpl struct {
	repo reviews.ReviewsRepo
}

func NewReviewsUsecase(repo reviews.ReviewsRepo) *reviewsUsecaseImpl {
	return &reviewsUsecaseImpl{repo: repo}
}

func (u *reviewsUsecaseImpl) CreateReview(ctx context.Context, review models.Review) error {
	return u.repo.CreateReview(ctx, review)
}

func (u *reviewsUsecaseImpl) UpdateReview(ctx context.Context, review models.Review) error {
	return u.repo.UpdateReview(ctx, review)
}

func (u *reviewsUsecaseImpl) DeleteReview(ctx context.Context, reviewID uint) error {
	return u.repo.DeleteReview(ctx, reviewID)
}

func (u *reviewsUsecaseImpl) GetReviewsByPlaceID(ctx context.Context, placeID uint) ([]models.Review, error) {
	return u.repo.GetReviewsByPlaceID(ctx, placeID)
}

func (u *reviewsUsecaseImpl) GetReview(ctx context.Context, reviewID uint) (models.Review, error) {
	return u.repo.GetReview(ctx, reviewID)
}
