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

	"google.golang.org/protobuf/types/known/timestamppb"
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
		ID:          uint(in.Trip.Id),
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
func (h *GrpcTripsHandler) CreateSharingLink(ctx context.Context, in *tripsGen.CreateSharingLinkRequest) (*tripsGen.CreateSharingLinkResponse, error) {
	token := in.Token
	sharingOption := in.SharingOption
	err := h.uc.CreateSharingLink(ctx, uint(in.TripId), token, sharingOption)
	if err != nil {
		h.logger.Error("Failed to dcreate a sharing link for a trip", slog.Any("error", err))
		return nil, err
	}

	return &tripsGen.CreateSharingLinkResponse{Token: token}, nil
}

func (h *GrpcTripsHandler) GetSharingToken(ctx context.Context, in *tripsGen.GetSharingTokenRequest) (*tripsGen.GetSharingTokenResponse, error) {
	tripID := in.TripId
	token, err := h.uc.GetSharingToken(ctx, uint(tripID))
	if err != nil {
		h.logger.Error("Failed to get sharing token for a trip", slog.Any("error", err))
		return nil, err
	}
	return &tripsGen.GetSharingTokenResponse{
		Token: &tripsGen.Token{
			Id:            uint32(token.ID),
			Token:         token.Token,
			SharingOption: token.SharingOption,
			ExpiresAt:     timestamppb.New(token.ExpiresAt),
		},
	}, nil
}

func (h *GrpcTripsHandler) GetTripBySharingToken(ctx context.Context, in *tripsGen.GetTripBySharingTokenRequest) (*tripsGen.GetTripBySharingTokenResponse, error) {
	token := in.Token
	trip, err := h.uc.GetTripBySharingToken(ctx, token)
	if err != nil {
		h.logger.Error("Failed to get trip by sharing token", slog.Any("error", err))
		return nil, err
	}
	return &tripsGen.GetTripBySharingTokenResponse{
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

func (h *GrpcTripsHandler) AddUserToTrip(ctx context.Context, in *tripsGen.AddUserToTripRequest) (*tripsGen.EmptyResponse, error) {
	tripId := in.TripId
	userId := in.UserId

	err := h.uc.AddUserToTrip(ctx, uint(tripId), uint(userId))
	if err != nil {
		h.logger.Error("Failed to add user to trip", slog.Any("error", err))
		return nil, err
	}
	return &tripsGen.EmptyResponse{}, nil
}

func (h *GrpcTripsHandler) GetSharingOption(ctx context.Context, in *tripsGen.GetSharingOptionRequest) (*tripsGen.GetSharingOptionResponse, error) {
	tripId := in.TripId
	userId := in.UserId

	sharingOption, err := h.uc.GetSharingOption(ctx, uint(userId), uint(tripId))
	if err != nil {
		h.logger.Error("Failed to rettrieve sharing option", slog.Any("error", err))
		return nil, err
	}
	return &tripsGen.GetSharingOptionResponse{
		SharingOption: sharingOption,
	}, nil
}
