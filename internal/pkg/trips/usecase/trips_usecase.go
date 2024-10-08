package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/trips"
	"context"
	"errors"
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
	return u.tripRepo.CreateTrip(ctx, trip)
}

func (u *TripsUsecaseImpl) UpdateTrip(ctx context.Context, trip models.Trip) error {
	err := u.tripRepo.UpdateTrip(ctx, trip)
	if err != nil {
		if err.Error() == "no rows were updated" {
			return errors.New("trip not found")
		}
		return err
	}

	return nil
}

func (u *TripsUsecaseImpl) DeleteTrip(ctx context.Context, id uint) error {
	err := u.tripRepo.DeleteTrip(ctx, id)
	if err != nil {
		if err.Error() == "no rows were deleted" {
			return errors.New("trip not found")
		}
		return err
	}

	return nil
}

func (u *TripsUsecaseImpl) GetTripsByUserID(ctx context.Context, userID uint) ([]models.Trip, error) {
	trips, err := u.tripRepo.GetTripsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return trips, nil
}

func (u *TripsUsecaseImpl) GetTrip(ctx context.Context, tripID uint) (models.Trip, error) {
	trip, err := u.tripRepo.GetTrip(ctx, tripID)
	if err != nil {
		return models.Trip{}, err
	}
	return trip, nil
}
