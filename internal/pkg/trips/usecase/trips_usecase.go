package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/trips"
	"context"
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
	return u.tripRepo.UpdateTrip(ctx, trip)
}

func (u *TripsUsecaseImpl) DeleteTrip(ctx context.Context, id uint) error {
	return u.tripRepo.DeleteTrip(ctx, id)
}

func (u *TripsUsecaseImpl) ReadTripsByUserID(ctx context.Context, userID uint) ([]models.Trip, error) {
	trips, err := u.tripRepo.ReadTripsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return trips, nil
}

func (u *TripsUsecaseImpl) ReadTrip(ctx context.Context, tripID uint) (models.Trip, error) {
	trip, err := u.tripRepo.ReadTrip(ctx, tripID)
	if err != nil {
		return models.Trip{}, err
	}
	return trip, nil
}
