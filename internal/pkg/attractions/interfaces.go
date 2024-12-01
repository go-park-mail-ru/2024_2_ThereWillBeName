package attractions

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type PlaceRepo interface {
	GetPlaces(ctx context.Context, limit, offset int) ([]models.GetPlace, error)
	GetPlace(ctx context.Context, id uint) (models.GetPlace, error)
	SearchPlaces(ctx context.Context, name string, category, city, limit, offset int) ([]models.GetPlace, error)
	GetPlacesByCategory(ctx context.Context, category string, limit, offset int) ([]models.GetPlace, error)
}

type PlaceUsecase interface {
	GetPlaces(ctx context.Context, limit, offset int) ([]models.GetPlace, error)
	GetPlace(ctx context.Context, id uint) (models.GetPlace, error)
	SearchPlaces(ctx context.Context, name string, category, city, limit, offset int) ([]models.GetPlace, error)
	GetPlacesByCategory(ctx context.Context, category string, limit, offset int) ([]models.GetPlace, error)
}
