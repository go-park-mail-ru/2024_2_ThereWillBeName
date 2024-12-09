package trips

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

type TripsUsecase interface {
	CreateTrip(ctx context.Context, trip models.Trip) error
	UpdateTrip(ctx context.Context, user models.Trip) error
	DeleteTrip(ctx context.Context, id uint) error
	GetTripsByUserID(ctx context.Context, userID uint, limit, offset int) ([]models.Trip, error)
	GetTrip(ctx context.Context, tripID uint) (models.Trip, error)
	AddPlaceToTrip(ctx context.Context, tripID uint, placeID uint) error
	AddPhotosToTrip(ctx context.Context, tripID uint, photos []string) error
	DeletePhotoFromTrip(ctx context.Context, tripID uint, photoPath string) error
	CreateSharingLink(ctx context.Context, tripID uint, token string, sharingOption string) error
	GetSharingToken(ctx context.Context, tripID uint) (models.SharingToken, error)
}

type TripsRepo interface {
	CreateTrip(ctx context.Context, user models.Trip) error
	UpdateTrip(ctx context.Context, user models.Trip) error
	DeleteTrip(ctx context.Context, id uint) error
	GetTripsByUserID(ctx context.Context, userID uint, limit, offset int) ([]models.Trip, error)
	GetTrip(ctx context.Context, tripID uint) (models.Trip, error)
	AddPlaceToTrip(ctx context.Context, tripID uint, placeID uint) error
	AddPhotoToTrip(ctx context.Context, tripID uint, photoPath string) error
	DeletePhotoFromTrip(ctx context.Context, tripID uint, photoPath string) error
	CreateSharingLink(ctx context.Context, tripID uint, token string, sharingOption string) error
	GetSharingToken(ctx context.Context, tripID uint) (models.SharingToken, error)
}
