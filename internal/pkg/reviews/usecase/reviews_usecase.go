package reviews

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/reviews"
	"context"
	"errors"
	"fmt"
)

type reviewsUsecaseImpl struct {
	repo reviews.ReviewsRepo
}

func NewReviewsUsecase(repo reviews.ReviewsRepo) *reviewsUsecaseImpl {
	return &reviewsUsecaseImpl{repo: repo}
}

func (u *reviewsUsecaseImpl) CreateReview(ctx context.Context, review models.Review) (models.GetReview, error) {
	createdReview, err := u.repo.CreateReview(ctx, review)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return models.GetReview{}, fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		return models.GetReview{}, fmt.Errorf("internal error: %w", models.ErrInternal)
	}

	return createdReview, nil
}

func (u *reviewsUsecaseImpl) UpdateReview(ctx context.Context, review models.Review) error {
	err := u.repo.UpdateReview(ctx, review)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		return fmt.Errorf("internal error: %w", models.ErrInternal)
	}

	return nil
}

func (u *reviewsUsecaseImpl) DeleteReview(ctx context.Context, reviewID uint) error {
	err := u.repo.DeleteReview(ctx, reviewID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		return fmt.Errorf("internal error: %w", models.ErrInternal)
	}

	return nil
}

func (u *reviewsUsecaseImpl) GetReviewsByPlaceID(ctx context.Context, placeID uint, limit, offset int) ([]models.GetReview, error) {
	reviewsFound, err := u.repo.GetReviewsByPlaceID(ctx, placeID, limit, offset)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		return nil, fmt.Errorf("internal error: %w", models.ErrInternal)
	}
	return reviewsFound, nil
}

func (u *reviewsUsecaseImpl) GetReviewsByUserID(ctx context.Context, userID uint, limit, offset int) ([]models.GetReviewByUserID, error) {
	reviewsFound, err := u.repo.GetReviewsByUserID(ctx, userID, limit, offset)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		return nil, fmt.Errorf("internal error: %w", models.ErrInternal)
	}

	return reviewsFound, nil
}

func (u *reviewsUsecaseImpl) GetReview(ctx context.Context, reviewID uint) (models.GetReview, error) {
	reviewFound, err := u.repo.GetReview(ctx, reviewID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return models.GetReview{}, fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		return models.GetReview{}, fmt.Errorf("internal error: %w", models.ErrInternal)
	}

	return reviewFound, nil
}
