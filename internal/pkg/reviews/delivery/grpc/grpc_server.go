package grpc

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/reviews"
	"2024_2_ThereWillBeName/internal/pkg/reviews/delivery/grpc/gen"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//go:generate protoc -I . proto/reviews.proto --go_out=./gen --go-grpc_out=./gen

type GrpcReviewsHandler struct {
	gen.UnimplementedReviewsServer
	uc reviews.ReviewsUsecase
}

func NewGrpcReviewsHandler(uc reviews.ReviewsUsecase) *GrpcReviewsHandler {
	return &GrpcReviewsHandler{uc: uc}
}

func (s *GrpcReviewsHandler) CreateReview(ctx context.Context, req *gen.CreateReviewRequest) (*gen.CreateReviewResponse, error) {
	review := models.Review{
		ID:         uint(req.Review.Id),
		UserID:     uint(req.Review.UserId),
		PlaceID:    uint(req.Review.PlaceId),
		Rating:     int(req.Review.Rating),
		ReviewText: req.Review.ReviewText,
	}

	res, err := s.uc.CreateReview(ctx, review)
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

func (s *GrpcReviewsHandler) UpdateReview(ctx context.Context, req *gen.UpdateReviewRequest) (*gen.UpdateReviewResponse, error) {
	review := models.Review{
		ID:         uint(req.Review.Id),
		UserID:     uint(req.Review.UserId),
		PlaceID:    uint(req.Review.PlaceId),
		Rating:     int(req.Review.Rating),
		ReviewText: req.Review.ReviewText,
	}

	err := s.uc.UpdateReview(ctx, review)
	if err != nil {
		return &gen.UpdateReviewResponse{Success: false}, err
	}
	return &gen.UpdateReviewResponse{Success: true}, nil
}

func (s *GrpcReviewsHandler) DeleteReview(ctx context.Context, req *gen.DeleteReviewRequest) (*gen.DeleteReviewResponse, error) {
	err := s.uc.DeleteReview(ctx, uint(req.Id))
	if err != nil {
		return &gen.DeleteReviewResponse{Success: false}, err
	}
	return &gen.DeleteReviewResponse{Success: true}, nil
}

func (s *GrpcReviewsHandler) GetReviewsByPlaceID(ctx context.Context, req *gen.GetReviewsByPlaceIDRequest) (*gen.GetReviewsByPlaceIDResponse, error) {
	reviews, err := s.uc.GetReviewsByPlaceID(ctx, uint(req.PlaceId), int(req.Limit), int(req.Offset))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "reviews for place ID %d not found", req.PlaceId)
		}
		return nil, status.Errorf(codes.Internal, "failed to get reviews: %v", err)
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

func (s *GrpcReviewsHandler) GetReviewsByUserID(ctx context.Context, req *gen.GetReviewsByUserIDRequest) (*gen.GetReviewsByUserIDResponse, error) {
	reviews, err := s.uc.GetReviewsByUserID(ctx, uint(req.UserId), int(req.Limit), int(req.Offset))
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, "reviews for place ID %d not found", req.UserId)
		}
		return nil, status.Errorf(codes.Internal, "failed to get reviews: %v", err)
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

func (s *GrpcReviewsHandler) GetReview(ctx context.Context, req *gen.GetReviewRequest) (*gen.GetReviewResponse, error) {
	review, err := s.uc.GetReview(ctx, uint(req.Id))
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
