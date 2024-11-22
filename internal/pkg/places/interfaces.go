package places

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type PlaceRepo interface {
	GetPlaces(ctx context.Context, limit, offset int) ([]models.GetPlace, error)
	CreatePlace(ctx context.Context, place models.CreatePlace) error
	GetPlace(ctx context.Context, id uint) (models.GetPlace, error)
	UpdatePlace(ctx context.Context, place models.UpdatePlace) error
	DeletePlace(ctx context.Context, id uint) error
	SearchPlaces(ctx context.Context, name string, limit, offset int) ([]models.GetPlace, error)
	GetPlacesByCategory(ctx context.Context, category string, limit, offset int) ([]models.GetPlace, error)
	// SearchPlacesBySubString(ctx context.Context, query string) ([]models.SearchItem, error)
}

type PlaceUsecase interface {
	GetPlaces(ctx context.Context, limit, offset int) ([]models.GetPlace, error)
	CreatePlace(ctx context.Context, place models.CreatePlace) error
	GetPlace(ctx context.Context, id uint) (models.GetPlace, error)
	UpdatePlace(ctx context.Context, place models.UpdatePlace) error
	DeletePlace(ctx context.Context, id uint) error
	SearchPlaces(ctx context.Context, name string, limit, offset int) ([]models.GetPlace, error)
	GetPlacesByCategory(ctx context.Context, category string, limit, offset int) ([]models.GetPlace, error)
}
