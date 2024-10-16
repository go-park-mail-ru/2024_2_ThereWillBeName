package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/trips"
	"context"
	"errors"
	"fmt"
	"log"
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
		log.Println(err)
		if errors.Is(err, models.ErrNotFound) {
			return fmt.Errorf("invalid request: %w", models.ErrNotFound)
		} else {
			return fmt.Errorf("internal error: %w", models.ErrInternal)
		}
	}

	return nil
}

func (u *TripsUsecaseImpl) UpdateTrip(ctx context.Context, trip models.Trip) error {
	err := u.tripRepo.UpdateTrip(ctx, trip)
	if err != nil {
		log.Println(err)
		if errors.Is(err, models.ErrNotFound) {
			return fmt.Errorf("invalid request: %w", models.ErrNotFound)
		} else {
			return fmt.Errorf("internal error: %w", models.ErrInternal)
		}
	}

	return nil
}

func (u *TripsUsecaseImpl) DeleteTrip(ctx context.Context, id uint) error {
	err := u.tripRepo.DeleteTrip(ctx, id)
	if err != nil {
		log.Println(err)
		if errors.Is(err, models.ErrNotFound) {
			return fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		return fmt.Errorf("internal error: %w", models.ErrInternal)
	}

	return nil
}

func (u *TripsUsecaseImpl) GetTripsByUserID(ctx context.Context, userID uint, limit, offset int) ([]models.Trip, error) {
	tripsFound, err := u.tripRepo.GetTripsByUserID(ctx, userID, limit, offset)
	if err != nil {
		log.Println(err)
		if errors.Is(err, models.ErrNotFound) {
			return nil, fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		return nil, fmt.Errorf("internal error: %w", models.ErrInternal)
	}
	return tripsFound, nil
}

func (u *TripsUsecaseImpl) GetTrip(ctx context.Context, tripID uint) (models.Trip, error) {
	trip, err := u.tripRepo.GetTrip(ctx, tripID)
	if err != nil {
		log.Println(err)
		if errors.Is(err, models.ErrNotFound) {
			return models.Trip{}, fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		return models.Trip{}, fmt.Errorf("internal error^ %w", models.ErrInternal)
	}
	return trip, nil
}