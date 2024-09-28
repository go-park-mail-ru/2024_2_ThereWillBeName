package places

import (
	"TripAdvisor/internal/models"
)

type PlaceRepo interface {
	GetPlaces() ([]models.Place, error)
}

type PlaceUsecase interface {
	GetPlaces() ([]models.Place, error)
}
