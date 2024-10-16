package places

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type PlaceRepo interface {
	GetPlaces(ctx context.Context, limit, offset int) ([]models.Place, error)
	CreatePlace(ctx context.Context, place models.Place) error
	GetPlace(ctx context.Context, id int) (models.Place, error)
	UpdatePlace(ctx context.Context, place models.Place) error
	DeletePlace(ctx context.Context, id int) error
	SearchPlaces(ctx context.Context, name string, limit, offset int) ([]models.Place, error)
}

type PlaceUsecase interface {
	GetPlaces(ctx context.Context, limit, offset int) ([]models.Place, error)
	CreatePlace(ctx context.Context, place models.Place) error
	GetPlace(ctx context.Context, id int) (models.Place, error)
	UpdatePlace(ctx context.Context, place models.Place) error
	DeletePlace(ctx context.Context, id int) error
	SearchPlaces(ctx context.Context, name string, limit, offset int) ([]models.Place, error)
}
