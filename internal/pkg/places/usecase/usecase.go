package usecase

import (
	"TripAdvisor/internal/models"
	"TripAdvisor/internal/pkg/places"
)

type PlaceUsecaseImpl struct {
	repo places.PlaceRepo
}

func NewPlaceUsecase(repo places.PlaceRepo) *PlaceUsecaseImpl {
	return &PlaceUsecaseImpl{repo: repo}
}

func (i *PlaceUsecaseImpl) GetPlaces() ([]models.Place, error) {
	places, err := i.repo.GetPlaces()
	if err != nil {
		return nil, err
	}
	return places, nil
}
