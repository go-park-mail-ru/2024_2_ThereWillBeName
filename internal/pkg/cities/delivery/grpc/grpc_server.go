package grpc

import (
	"2024_2_ThereWillBeName/internal/pkg/cities"
	"2024_2_ThereWillBeName/internal/pkg/cities/delivery/grpc/gen"
	"context"
)

//go:generate protoc -I . proto/cities.proto --go_out=./gen --go-grpc_out=./gen

type GrpcCitiesHandler struct {
	gen.UnimplementedCitiesServer
	uc cities.CitiesUsecase
}

func NewGrpcCitiesHandler(uc cities.CitiesUsecase) *GrpcCitiesHandler {
	return &GrpcCitiesHandler{uc: uc}
}

func (s *GrpcCitiesHandler) SearchCitiesByName(ctx context.Context, req *gen.SearchCitiesByNameRequest) (*gen.SearchCitiesByNameResponse, error) {
	cities, err := s.uc.SearchCitiesByName(ctx, req.Query)
	if err != nil {
		return nil, err
	}

	citiesResponse := make([]*gen.City, len(cities))
	for i, city := range cities {
		citiesResponse[i] = &gen.City{
			Id:   uint32(city.ID),
			Name: city.Name,
		}
	}
	return &gen.SearchCitiesByNameResponse{Cities: citiesResponse}, nil
}

func (s *GrpcCitiesHandler) SearchCityByID(ctx context.Context, req *gen.SearchCityByIDRequest) (*gen.SearchCityByIDResponse, error) {
	city, err := s.uc.SearchCityByID(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &gen.SearchCityByIDResponse{City: &gen.City{Id: uint32(city.ID), Name: city.Name}}, nil
}
