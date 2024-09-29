package places

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

type PlaceRepo interface {
	GetPlaces(ctx context.Context) ([]models.Place, error)
}

type PlaceUsecase interface {
	GetPlaces(ctx context.Context) ([]models.Place, error)
}
