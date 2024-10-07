package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/places"
	"context"
)

type PlaceUsecaseImpl struct {
	repo places.PlaceRepo
}

func NewPlaceUsecase(repo places.PlaceRepo) *PlaceUsecaseImpl {
	return &PlaceUsecaseImpl{repo: repo}
}

func (i *PlaceUsecaseImpl) GetPlaces(ctx context.Context) ([]models.Place, error) {
	places, err := i.repo.GetPlaces(ctx)
	if err != nil {
		return nil, err
	}
	return places, nil
}

func (i *PlaceUsecaseImpl) CreatePlace(ctx context.Context, place models.Place) error {
	return i.repo.CreatePlace(ctx, place)
}

func (i *PlaceUsecaseImpl) UpdatePlace(ctx context.Context, place models.Place) error {
	return i.repo.UpdatePlace(ctx, place)
}

func (i *PlaceUsecaseImpl) DeletePlace(ctx context.Context, name string) error {
	return i.repo.DeletePlace(ctx, name)
}

func (i *PlaceUsecaseImpl) GetPlace(ctx context.Context, name string) (models.Place, error) {
	place, err := i.repo.GetPlace(ctx, name)
	if err != nil {
		return models.Place{}, err
	}
	return place, nil
}

func (i *PlaceUsecaseImpl) GetPlacesBySearch(ctx context.Context, name string) ([]models.Place, error) {
	return i.repo.GetPlacesBySearch(ctx, name)
}
