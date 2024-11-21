package grpc

//go:generate protoc -I . proto/attractions.proto --go_out=./gen --go-grpc_out=./gen

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/attractions"
	"2024_2_ThereWillBeName/internal/pkg/attractions/delivery/grpc/gen"
	categoriespkg "2024_2_ThereWillBeName/internal/pkg/categories"
	citiesPkg "2024_2_ThereWillBeName/internal/pkg/cities"
	reviewsPkg "2024_2_ThereWillBeName/internal/pkg/reviews"
	"context"
)

type GrpcAttractionsHandler struct {
	gen.UnimplementedAttractionsServer
	placeUsecase      attractions.PlaceUsecase
	citiesUsecase     citiesPkg.CitiesUsecase
	reviewUsecase     reviewsPkg.ReviewsUsecase
	categoriesUsecase categoriespkg.CategoriesUsecase
}

func NewGrpcAttractionsHandler(
	placeUsecase attractions.PlaceUsecase,
	citiesUsecase citiesPkg.CitiesUsecase,
	reviewUsecase reviewsPkg.ReviewsUsecase,
	categoriesUsecase categoriespkg.CategoriesUsecase,
) *GrpcAttractionsHandler {
	return &GrpcAttractionsHandler{
		placeUsecase:      placeUsecase,
		citiesUsecase:     citiesUsecase,
		reviewUsecase:     reviewUsecase,
		categoriesUsecase: categoriesUsecase,
	}
}

func (s *GrpcAttractionsHandler) GetPlaces(ctx context.Context, req *gen.GetPlacesRequest) (*gen.GetPlacesResponse, error) {
	places, err := s.placeUsecase.GetPlaces(ctx, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}
	placesResponse := make([]*gen.Place, len(places))
	for i, place := range places {
		placesResponse[i] = &gen.Place{
			Id:          uint32(place.ID),
			Name:        place.Name,
			ImagePath:   place.ImagePath,
			Description: place.Description,
			Rating:      int32(place.Rating),
			Address:     place.Address,
			City:        place.City,
			PhoneNumber: place.PhoneNumber,
			Categories:  place.Categories,
			Latitude:    place.Latitude,
			Longitude:   place.Longitude,
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
		Id:          uint32(place.ID),
		Name:        place.Name,
		Description: place.Description,
		Rating:      int32(place.Rating),
		Address:     place.Address,
		City:        place.City,
		PhoneNumber: place.PhoneNumber,
		Categories:  place.Categories,
		Latitude:    place.Latitude,
		Longitude:   place.Longitude,
	}
	return &gen.GetPlaceResponse{Place: placeResponse}, nil
}

func (s *GrpcAttractionsHandler) SearchPlaces(ctx context.Context, req *gen.SearchPlacesRequest) (*gen.SearchPlacesResponse, error) {
	places, err := s.placeUsecase.SearchPlaces(ctx, req.Name, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	placesResponse := make([]*gen.Place, len(places))
	for i, place := range places {
		placesResponse[i] = &gen.Place{
			Id:          uint32(place.ID),
			Name:        place.Name,
			Description: place.Description,
			Rating:      int32(place.Rating),
			Address:     place.Address,
			City:        place.City,
			PhoneNumber: place.PhoneNumber,
			Categories:  place.Categories,
			Latitude:    place.Latitude,
			Longitude:   place.Longitude,
		}
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
			Id:          uint32(place.ID),
			Name:        place.Name,
			Description: place.Description,
			Rating:      int32(place.Rating),
			Address:     place.Address,
			City:        place.City,
			PhoneNumber: place.PhoneNumber,
			Categories:  place.Categories,
			Latitude:    place.Latitude,
			Longitude:   place.Longitude,
		}
	}
	return &gen.GetPlacesByCategoryResponse{Places: placesResponse}, nil
}

func (s *GrpcAttractionsHandler) GetCategories(ctx context.Context, req *gen.GetCategoriesRequest) (*gen.GetCategoriesResponse, error) {
	categories, err := s.categoriesUsecase.GetCategories(ctx, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	categoriesResponse := make([]*gen.Category, len(categories))
	for i, category := range categories {
		categoriesResponse[i] = &gen.Category{
			Id:   uint32(category.ID),
			Name: category.Name,
		}
	}
	return &gen.GetCategoriesResponse{Categories: categoriesResponse}, nil
}

func (s *GrpcAttractionsHandler) SearchCitiesByName(ctx context.Context, req *gen.SearchCitiesByNameRequest) (*gen.SearchCitiesByNameResponse, error) {
	cities, err := s.citiesUsecase.SearchCitiesByName(ctx, req.Query)
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

func (s *GrpcAttractionsHandler) SearchCityByID(ctx context.Context, req *gen.SearchCityByIDRequest) (*gen.SearchCityByIDResponse, error) {
	city, err := s.citiesUsecase.SearchCityByID(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &gen.SearchCityByIDResponse{City: &gen.City{Id: uint32(city.ID), Name: city.Name}}, nil
}

func (s *GrpcAttractionsHandler) CreateReview(ctx context.Context, req *gen.CreateReviewRequest) (*gen.CreateReviewResponse, error) {
	review := models.Review{
		ID:         uint(req.Review.Id),
		UserID:     uint(req.Review.UserId),
		PlaceID:    uint(req.Review.PlaceId),
		Rating:     int(req.Review.Rating),
		ReviewText: req.Review.ReviewText,
	}

	res, err := s.reviewUsecase.CreateReview(ctx, review)
	if err != nil {
		return nil, err
	}
	reviewResponse := &gen.GetReview{
		Id:         uint32(res.ID),
		UserLogin:  res.UserLogin,
		AvatarPath: res.AvatarPath,
		Rating:     int32(res.Rating),
		ReviewText: res.ReviewText,
	}
	return &gen.CreateReviewResponse{Review: reviewResponse}, nil
}

func (s *GrpcAttractionsHandler) UpdateReview(ctx context.Context, req *gen.UpdateReviewRequest) (*gen.UpdateReviewResponse, error) {
	review := models.Review{
		ID:         uint(req.Review.Id),
		UserID:     uint(req.Review.UserId),
		PlaceID:    uint(req.Review.PlaceId),
		Rating:     int(req.Review.Rating),
		ReviewText: req.Review.ReviewText,
	}

	err := s.reviewUsecase.UpdateReview(ctx, review)
	if err != nil {
		return &gen.UpdateReviewResponse{Success: false}, err
	}
	return &gen.UpdateReviewResponse{Success: true}, nil
}

func (s *GrpcAttractionsHandler) DeleteReview(ctx context.Context, req *gen.DeleteReviewRequest) (*gen.DeleteReviewResponse, error) {
	err := s.reviewUsecase.DeleteReview(ctx, uint(req.Id))
	if err != nil {
		return &gen.DeleteReviewResponse{Success: false}, err
	}
	return &gen.DeleteReviewResponse{Success: true}, nil
}

func (s *GrpcAttractionsHandler) GetReviewsByPlaceID(ctx context.Context, req *gen.GetReviewsByPlaceIDRequest) (*gen.GetReviewsByPlaceIDResponse, error) {
	reviews, err := s.reviewUsecase.GetReviewsByPlaceID(ctx, uint(req.PlaceId), int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	reviewsResponse := make([]*gen.GetReview, len(reviews))
	for i, review := range reviews {
		reviewsResponse[i] = &gen.GetReview{
			Id:         uint32(review.ID),
			UserLogin:  review.UserLogin,
			AvatarPath: review.AvatarPath,
			Rating:     int32(review.Rating),
			ReviewText: review.ReviewText,
		}
	}
	return &gen.GetReviewsByPlaceIDResponse{Reviews: reviewsResponse}, nil
}

func (s *GrpcAttractionsHandler) GetReviewsByUserID(ctx context.Context, req *gen.GetReviewsByUserIDRequest) (*gen.GetReviewsByUserIDResponse, error) {
	reviews, err := s.reviewUsecase.GetReviewsByUserID(ctx, uint(req.UserId), int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, err
	}

	reviewsResponse := make([]*gen.GetReviewByUserID, len(reviews))
	for i, review := range reviews {
		reviewsResponse[i] = &gen.GetReviewByUserID{
			Id:         uint32(review.ID),
			PlaceName:  review.PlaceName,
			Rating:     int32(review.Rating),
			ReviewText: review.ReviewText,
		}
	}
	return &gen.GetReviewsByUserIDResponse{Reviews: reviewsResponse}, nil
}

func (s *GrpcAttractionsHandler) GetReview(ctx context.Context, req *gen.GetReviewRequest) (*gen.GetReviewResponse, error) {
	review, err := s.reviewUsecase.GetReview(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	reviewResponse := &gen.GetReview{
		Id:         uint32(review.ID),
		UserLogin:  review.UserLogin,
		AvatarPath: review.AvatarPath,
		Rating:     int32(review.Rating),
		ReviewText: review.ReviewText,
	}
	return &gen.GetReviewResponse{Review: reviewResponse}, nil
}
