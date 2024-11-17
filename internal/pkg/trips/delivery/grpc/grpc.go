package grpc

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/trips"
	tripsGen "2024_2_ThereWillBeName/internal/pkg/trips/delivery/grpc/gen"
	"context"
	"log/slog"
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
		},
	}, nil
}

func (h *GrpcTripsHandler) AddPlaceToTrip(ctx context.Context, in *tripsGen.AddPlaceToTripRequest) (*tripsGen.EmptyResponse, error) {
	err := h.uc.AddPlaceToTrip(context.Background(), uint(in.TripId), uint(in.PlaceId))
	if err != nil {
		return nil, err
	}
	return &tripsGen.EmptyResponse{}, nil
}
