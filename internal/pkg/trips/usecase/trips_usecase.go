package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/trips"
	"context"
	"errors"
	"fmt"
	"log"
	"path"
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
		if errors.Is(err, models.ErrNotFound) {
			return fmt.Errorf("invalid request: %w", models.ErrNotFound)
		} else {
			return fmt.Errorf("internal error: %w", models.ErrInternal)
		}
	}

	return nil
}

func (u *TripsUsecaseImpl) UpdateTrip(ctx context.Context, trip models.Trip) error {
	log.Println("debug in usecase", trip.ID)
	err := u.tripRepo.UpdateTrip(ctx, trip)
	if err != nil {
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
		if errors.Is(err, models.ErrNotFound) {
			return nil, fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		return nil, fmt.Errorf("internal error: %w", models.ErrInternal)
	}
	return tripsFound, nil
}

func (u *TripsUsecaseImpl) GetTrip(ctx context.Context, tripID uint) (models.Trip, []models.UserProfile, error) {
	trip, userProfiles, err := u.tripRepo.GetTrip(ctx, tripID)
	if err != nil {
		return models.Trip{}, nil, err
	}
	return trip, userProfiles, nil
}

func (u *TripsUsecaseImpl) AddPlaceToTrip(ctx context.Context, tripID uint, placeID uint) error {
	err := u.tripRepo.AddPlaceToTrip(ctx, tripID, placeID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return fmt.Errorf("invalid request: %w", models.ErrNotFound)
		} else {
			return fmt.Errorf("internal error: %w", models.ErrInternal)
		}
	}

	return nil
}

func (u *TripsUsecaseImpl) AddPhotosToTrip(ctx context.Context, tripID uint, photoPaths []string) error {
	for _, fullpath := range photoPaths {
		filename := path.Base(fullpath)
		err := u.tripRepo.AddPhotoToTrip(ctx, tripID, filename)
		if err != nil {
			return fmt.Errorf("failed to add photo to trip: %w", err)
		}
	}
	return nil
}

func (u *TripsUsecaseImpl) DeletePhotoFromTrip(ctx context.Context, tripID uint, photoPath string) error {
	err := u.tripRepo.DeletePhotoFromTrip(ctx, tripID, photoPath)
	if err != nil {
		return fmt.Errorf("failed to delete photo from database: %w", err)
	}

	return nil
}

func (u *TripsUsecaseImpl) CreateSharingLink(ctx context.Context, tripID uint, token string, sharingOption string) error {
	err := u.tripRepo.CreateSharingLink(ctx, tripID, token, sharingOption)
	if err != nil {
		return fmt.Errorf("failed to delete photo from database: %w", err)
	}
	return nil
}

func (u *TripsUsecaseImpl) GetSharingToken(ctx context.Context, tripID uint) (models.SharingToken, error) {
	token, err := u.tripRepo.GetSharingToken(ctx, tripID)
	if err != nil {
		return models.SharingToken{}, fmt.Errorf("failed to retrieve token from database: %w", err)
	}
	return token, nil
}

func (u *TripsUsecaseImpl) GetTripBySharingToken(ctx context.Context, token string) (models.Trip, []models.UserProfile, error) {
	trip, users, err := u.tripRepo.GetTripBySharingToken(ctx, token)
	if err != nil {
		return models.Trip{}, nil, fmt.Errorf("failed to retrieve trip by sharing token from database: %w", err)
	}
	return trip, users, nil
}

func (u *TripsUsecaseImpl) AddUserToTrip(ctx context.Context, tripId, userId uint) error {
	err := u.tripRepo.AddUserToTrip(ctx, tripId, userId)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return fmt.Errorf("invalid request: %w", models.ErrNotFound)
		} else {
			return fmt.Errorf("internal error: %w", models.ErrInternal)
		}
	}

	return nil
}

func (u *TripsUsecaseImpl) GetSharingOption(ctx context.Context, userId, tripId uint) (string, error) {
	sharingOption, err := u.tripRepo.GetSharingOption(ctx, userId, tripId)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return "", fmt.Errorf("invalid request: %w", models.ErrNotFound)
		} else {
			return "", fmt.Errorf("internal error: %w", models.ErrInternal)
		}
	}
	return sharingOption, nil
}
