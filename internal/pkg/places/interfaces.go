package places

import (
	"TripAdvisor/internal/models"
	"context"
)

type PlaceRepo interface {
	GetPlaces(ctx context.Context) ([]models.Place, error)
}

type PlaceUsecase interface {
	GetPlaces(ctx context.Context) ([]models.Place, error)
}
