package trips

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

type TripsUsecase interface {
	CreateTrip(ctx context.Context, trip models.Trip) error
	UpdateTrip(ctx context.Context, user models.Trip) error
	DeleteTrip(ctx context.Context, id uint) error
	GetTripsByUserID(ctx context.Context, userID uint) ([]models.Trip, error)
	GetTrip(ctx context.Context, tripID uint) (models.Trip, error)
}

type TripsRepo interface {
	CreateTrip(ctx context.Context, user models.Trip) error
	UpdateTrip(ctx context.Context, user models.Trip) error
	DeleteTrip(ctx context.Context, id uint) error
	GetTripsByUserID(ctx context.Context, userID uint) ([]models.Trip, error)
	GetTrip(ctx context.Context, tripID uint) (models.Trip, error)
}
