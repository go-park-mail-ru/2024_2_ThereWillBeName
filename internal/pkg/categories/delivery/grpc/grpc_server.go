package grpc

import (
	"2024_2_ThereWillBeName/internal/pkg/categories"
	"2024_2_ThereWillBeName/internal/pkg/categories/delivery/grpc/gen"
	"context"
)

//go:generate protoc -I . proto/categories.proto --go_out=./gen --go-grpc_out=./gen

type GrpcCategoriesHandler struct {
	gen.UnimplementedCategoriesServer
	uc categories.CategoriesUsecase
}

func NewGrpcCategoriesHandler(uc categories.CategoriesUsecase) *GrpcCategoriesHandler {
	return &GrpcCategoriesHandler{uc: uc}
}

func (s *GrpcCategoriesHandler) GetCategories(ctx context.Context, req *gen.GetCategoriesRequest) (*gen.GetCategoriesResponse, error) {
	categories, err := s.uc.GetCategories(ctx, int(req.Limit), int(req.Offset))
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
