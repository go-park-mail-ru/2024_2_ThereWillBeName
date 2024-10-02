package places

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock.go

type PlaceRepo interface {
	GetPlaces(ctx context.Context) ([]models.Place, error)
}

type PlaceUsecase interface {
	GetPlaces(ctx context.Context) ([]models.Place, error)
}
