package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/trips"
	"context"
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("trip not found")
	ErrConflict = errors.New("foreign key constraint violation")
	ErrInternal = errors.New("internal repository error")
	//ErrForeignKeyViolation = errors.New("foreign key constraint violation")
	ErrUserNotFound = errors.New("user not found")
)

type TripsUsecaseImpl struct {
	tripRepo trips.TripsRepo
}

func NewTripsUsecase(repo trips.TripsRepo) *TripsUsecaseImpl {
	return &TripsUsecaseImpl{
		tripRepo: repo,
	}
}

func (u *TripsUsecaseImpl) CreateTrip(ctx context.Context, trip models.Trip) error {
	err := u.tripRepo.CreateTrip(ctx, trip)
	if err != nil {
		if errors.Is(err, ErrConflict) {
			return fmt.Errorf("invalid request: %w", ErrConflict)
		} else if errors.Is(err, ErrNotFound) {
			return fmt.Errorf("invalid request: %w", ErrNotFound)
		} else {
			return fmt.Errorf("internal error: %w", err)
		}
	}

	return nil
}

func (u *TripsUsecaseImpl) UpdateTrip(ctx context.Context, trip models.Trip) error {
	err := u.tripRepo.UpdateTrip(ctx, trip)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return fmt.Errorf("invalid request: %w", ErrNotFound)
		} else {
			return fmt.Errorf("internal error: %w", err)
		}
	}

	return nil
}

func (u *TripsUsecaseImpl) DeleteTrip(ctx context.Context, id uint) error {
	err := u.tripRepo.DeleteTrip(ctx, id)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return fmt.Errorf("invalid request: %w", ErrNotFound)
		} else if errors.Is(err, ErrConflict) {
			return fmt.Errorf("invalid request: %w", ErrConflict)
		}
		return fmt.Errorf("internal error: %w", err)
	}

	return nil
}

func (u *TripsUsecaseImpl) GetTripsByUserID(ctx context.Context, userID uint) ([]models.Trip, error) {
	trips, err := u.tripRepo.GetTripsByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return nil, fmt.Errorf("invalid request: %w", ErrUserNotFound)
		} else if errors.Is(err, ErrNotFound) {
			return nil, fmt.Errorf("invalid request: %w", ErrNotFound)
		}
		return nil, fmt.Errorf("internal error: %w", err)
	}
	return trips, nil
}

func (u *TripsUsecaseImpl) GetTrip(ctx context.Context, tripID uint) (models.Trip, error) {
	trip, err := u.tripRepo.GetTrip(ctx, tripID)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return models.Trip{}, fmt.Errorf("invalid request: %w", ErrNotFound)
		}
		return models.Trip{}, fmt.Errorf("internal error^ %w", err)
	}
	return trip, nil
}
