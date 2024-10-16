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

func (i *PlaceUsecaseImpl) GetPlaces(ctx context.Context, limit, offset int) ([]models.Place, error) {
	places, err := i.repo.GetPlaces(ctx, limit, offset)
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

func (i *PlaceUsecaseImpl) DeletePlace(ctx context.Context, id int) error {
	return i.repo.DeletePlace(ctx, id)
}

func (i *PlaceUsecaseImpl) GetPlace(ctx context.Context, id int) (models.Place, error) {
	return i.repo.GetPlace(ctx, id)
}

func (i *PlaceUsecaseImpl) SearchPlaces(ctx context.Context, name string, limit, offset int) ([]models.Place, error) {
	return i.repo.SearchPlaces(ctx, name, limit, offset)
}
