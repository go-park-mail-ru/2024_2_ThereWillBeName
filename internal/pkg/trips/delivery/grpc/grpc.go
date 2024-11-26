package grpc

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/trips"
	tripsGen "2024_2_ThereWillBeName/internal/pkg/trips/delivery/grpc/gen"
	"context"
	"encoding/base64"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type GrpcTripsHandler struct {
	tripsGen.TripsServer
	uc     trips.TripsUsecase
	logger *slog.Logger
}

func NewGrpcTripHandler(uc trips.TripsUsecase, logger *slog.Logger) *GrpcTripsHandler {
	return &GrpcTripsHandler{uc: uc, logger: logger}
}

func (h *GrpcTripsHandler) CreateTrip(ctx context.Context, in *tripsGen.CreateTripRequest) (*tripsGen.EmptyResponse, error) {

	trip := models.Trip{
		UserID:      uint(in.Trip.UserId),
		Name:        in.Trip.Name,
		Description: in.Trip.Description,
		CityID:      uint(in.Trip.CityId),
		StartDate:   in.Trip.StartDate,
		EndDate:     in.Trip.EndDate,
		Private:     in.Trip.Private,
	}

	err := h.uc.CreateTrip(context.Background(), trip)

	if err != nil {
		return nil, err
	}

	return &tripsGen.EmptyResponse{}, nil
}

func (h *GrpcTripsHandler) UpdateTrip(ctx context.Context, in *tripsGen.UpdateTripRequest) (*tripsGen.EmptyResponse, error) {

	trip := models.Trip{
		UserID:      uint(in.Trip.UserId),
		Name:        in.Trip.Name,
		Description: in.Trip.Description,
		CityID:      uint(in.Trip.CityId),
		StartDate:   in.Trip.StartDate,
		EndDate:     in.Trip.EndDate,
		Private:     in.Trip.Private,
	}

	err := h.uc.UpdateTrip(context.Background(), trip)

	if err != nil {
		return nil, err
	}

	return &tripsGen.EmptyResponse{}, nil
}

func (h *GrpcTripsHandler) DeleteTrip(ctx context.Context, in *tripsGen.DeleteTripRequest) (*tripsGen.EmptyResponse, error) {
	err := h.uc.DeleteTrip(context.Background(), uint(in.Id))

	if err != nil {
		return nil, err
	}

	return &tripsGen.EmptyResponse{}, nil
}

func (h *GrpcTripsHandler) GetTripsByUserID(ctx context.Context, in *tripsGen.GetTripsByUserIDRequest) (*tripsGen.GetTripsByUserIDResponse, error) {
	trips, err := h.uc.GetTripsByUserID(context.Background(), uint(in.UserId), int(in.Limit), int(in.Offset))
	if err != nil {
		return nil, err
	}
	grpcTrips := make([]*tripsGen.Trip, 0, len(trips))
	for _, trip := range trips {
		grpcTrips = append(grpcTrips, &tripsGen.Trip{
			Id:          uint32(trip.ID),
			UserId:      uint32(trip.UserID),
			Name:        trip.Name,
			Description: trip.Description,
			CityId:      uint32(trip.CityID),
			StartDate:   trip.StartDate,
			EndDate:     trip.EndDate,
			Photos:      trip.Photos,
			Private:     trip.Private,
		})
	}
	return &tripsGen.GetTripsByUserIDResponse{Trips: grpcTrips}, nil
}

func (h *GrpcTripsHandler) GetTrip(ctx context.Context, in *tripsGen.GetTripRequest) (*tripsGen.GetTripResponse, error) {
	trip, err := h.uc.GetTrip(context.Background(), uint(in.TripId))
	if err != nil {
		return nil, err
	}
	return &tripsGen.GetTripResponse{
		Trip: &tripsGen.Trip{
			Id:          uint32(trip.ID),
			UserId:      uint32(trip.UserID),
			Name:        trip.Name,
			Description: trip.Description,
			CityId:      uint32(trip.CityID),
			StartDate:   trip.StartDate,
			EndDate:     trip.EndDate,
			Private:     trip.Private,
			Photos:      trip.Photos,
		},
	}, nil
}

func (h *GrpcTripsHandler) AddPlaceToTrip(ctx context.Context, in *tripsGen.AddPlaceToTripRequest) (*tripsGen.EmptyResponse, error) {
	tripID := uint(in.TripId)
	placeID := uint(in.PlaceId)
	err := h.uc.AddPlaceToTrip(ctx, tripID, placeID)
	if err != nil {
		h.logger.Error("Failed to add place to trip", slog.Any("error", err))
		return nil, err
	}

	return &tripsGen.EmptyResponse{}, nil
}

func (h *GrpcTripsHandler) AddPhotosToTrip(ctx context.Context, in *tripsGen.AddPhotosToTripRequest) (*tripsGen.AddPhotosToTripResponse, error) {
	var savedPhotoPaths []string
	for _, base64Photo := range in.Photos {
		if strings.HasPrefix(base64Photo, "data:image/") {
			index := strings.Index(base64Photo, ",")
			if index != -1 {
				base64Photo = base64Photo[index+1:]
			} else {
				h.logger.Error("Invalid base64 photo format: missing ',' separator")
				return nil, fmt.Errorf("invalid base64 photo format: missing ',' separator")
			}
		}

		photoBytes, err := base64.StdEncoding.DecodeString(base64Photo)
		if err != nil {
			h.logger.Error("Failed to decode base64 photo", slog.Any("error", err))
			return nil, fmt.Errorf("invalid base64 data: %w", err)
		}

		photoName := fmt.Sprintf("trip_%d_%d.jpg", in.TripId, time.Now().UnixNano())
		photoPath := filepath.Join(os.Getenv("PHOTO_STORAGE_PATH"), photoName)
		if _, err := os.Stat(os.Getenv("PHOTO_STORAGE_PATH")); os.IsNotExist(err) {
			err := os.MkdirAll(os.Getenv("PHOTO_STORAGE_PATH"), 0755)
			if err != nil {
				h.logger.Error("Failed to create photo storage directory", slog.Any("error", err))
				return nil, fmt.Errorf("failed to create photo storage directory: %w", err)
			}
		}
		err = os.WriteFile(photoPath, photoBytes, 0644)
		if err != nil {
			h.logger.Error("Failed to save photo to disk", slog.String("path", photoPath), slog.Any("error", err))
			return nil, fmt.Errorf("failed to save photo: %w", err)
		}

		savedPhotoPaths = append(savedPhotoPaths, photoPath)
	}

	err := h.uc.AddPhotosToTrip(ctx, uint(in.TripId), savedPhotoPaths)
	if err != nil {
		h.logger.Error("Failed to save photo paths in database", slog.Any("error", err))
		return nil, fmt.Errorf("failed to save photo paths: %w", err)
	}

	var grpcPhotos []*tripsGen.Photo
	for _, path := range savedPhotoPaths {
		grpcPhotos = append(grpcPhotos, &tripsGen.Photo{PhotoPath: path})
	}

	return &tripsGen.AddPhotosToTripResponse{Photos: grpcPhotos}, nil
}

func (h *GrpcTripsHandler) DeletePhotoFromTrip(ctx context.Context, in *tripsGen.DeletePhotoRequest) (*tripsGen.EmptyResponse, error) {
	photoPath := in.PhotoPath
	err := h.uc.DeletePhotoFromTrip(ctx, uint(in.TripId), photoPath)
	if err != nil {
		h.logger.Error("Failed to delete photo path from database", slog.Any("error", err))
		return nil, fmt.Errorf("failed to delete photo path: %w", err)
	}

	err = os.Remove(photoPath)
	if err != nil && !os.IsNotExist(err) {
		h.logger.Error("Failed to delete photo file", slog.String("path", photoPath), slog.Any("error", err))
		return nil, fmt.Errorf("failed to delete photo file: %w", err)
	}

	h.logger.Info("Photo successfully deleted", slog.String("path", photoPath))
	return &tripsGen.EmptyResponse{}, nil
}
