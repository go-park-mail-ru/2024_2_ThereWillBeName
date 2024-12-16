package grpc

//go:generate protoc -I . proto/attractions.proto --go_out=./gen --go-grpc_out=./gen

import (
	"2024_2_ThereWillBeName/internal/pkg/attractions"
	"2024_2_ThereWillBeName/internal/pkg/attractions/delivery/grpc/gen"
	"context"
	"log"
)

type GrpcAttractionsHandler struct {
	gen.UnimplementedAttractionsServer
	placeUsecase attractions.PlaceUsecase
}

func NewGrpcAttractionsHandler(placeUsecase attractions.PlaceUsecase) *GrpcAttractionsHandler {
	return &GrpcAttractionsHandler{placeUsecase: placeUsecase}
}

func (s *GrpcAttractionsHandler) GetPlaces(ctx context.Context, req *gen.GetPlacesRequest) (*gen.GetPlacesResponse, error) {
	places, err := s.placeUsecase.GetPlaces(ctx, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}
	placesResponse := make([]*gen.Place, len(places))
	for i, place := range places {
		placesResponse[i] = &gen.Place{
			Id:              uint32(place.ID),
			Name:            place.Name,
			ImagePath:       place.ImagePath,
			Description:     place.Description,
			Rating:          float32(place.Rating),
			NumberOfReviews: uint32(place.NumberOfReviews),
			Address:         place.Address,
			City:            place.City,
			PhoneNumber:     place.PhoneNumber,
			Categories:      place.Categories,
			Latitude:        place.Latitude,
			Longitude:       place.Longitude,
		}
	}
	return &gen.GetPlacesResponse{Places: placesResponse}, nil
}

func (s *GrpcAttractionsHandler) GetPlace(ctx context.Context, req *gen.GetPlaceRequest) (*gen.GetPlaceResponse, error) {
	place, err := s.placeUsecase.GetPlace(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}
	placeResponse := &gen.Place{
		Id:              uint32(place.ID),
		Name:            place.Name,
		ImagePath:       place.ImagePath,
		Description:     place.Description,
		Rating:          float32(place.Rating),
		NumberOfReviews: uint32(place.NumberOfReviews),
		Address:         place.Address,
		City:            place.City,
		PhoneNumber:     place.PhoneNumber,
		Categories:      place.Categories,
		Latitude:        place.Latitude,
		Longitude:       place.Longitude,
	}
	return &gen.GetPlaceResponse{Place: placeResponse}, nil
}

func (s *GrpcAttractionsHandler) SearchPlaces(ctx context.Context, req *gen.SearchPlacesRequest) (*gen.SearchPlacesResponse, error) {
	places, err := s.placeUsecase.SearchPlaces(ctx, req.Name, int(req.Category), int(req.City), int(req.FilterType), int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	placesResponse := make([]*gen.Place, len(places))
	for i, place := range places {
		placesResponse[i] = &gen.Place{
			Id:              uint32(place.ID),
			Name:            place.Name,
			ImagePath:       place.ImagePath,
			Description:     place.Description,
			Rating:          float32(place.Rating),
			NumberOfReviews: uint32(place.NumberOfReviews),
			Address:         place.Address,
			City:            place.City,
			PhoneNumber:     place.PhoneNumber,
			Categories:      place.Categories,
			Latitude:        float32(place.Latitude),
			Longitude:       float32(place.Longitude),
		}
		log.Println("logging, latitude", place.Latitude, " ", float32(place.Latitude))
	}
	return &gen.SearchPlacesResponse{Places: placesResponse}, nil
}

func (s *GrpcAttractionsHandler) GetPlacesByCategory(ctx context.Context, req *gen.GetPlacesByCategoryRequest) (*gen.GetPlacesByCategoryResponse, error) {
	places, err := s.placeUsecase.GetPlacesByCategory(ctx, req.Category, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}
	placesResponse := make([]*gen.Place, len(places))
	for i, place := range places {
		placesResponse[i] = &gen.Place{
			Id:              uint32(place.ID),
			Name:            place.Name,
			Description:     place.Description,
			Rating:          float32(place.Rating),
			NumberOfReviews: uint32(place.NumberOfReviews),
			Address:         place.Address,
			City:            place.City,
			PhoneNumber:     place.PhoneNumber,
			Categories:      place.Categories,
			Latitude:        place.Latitude,
			Longitude:       place.Longitude,
		}
	}
	return &gen.GetPlacesByCategoryResponse{Places: placesResponse}, nil
}
