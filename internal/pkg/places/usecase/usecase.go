package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/places"
	"context"
	"log"
)

type PlaceUsecaseImpl struct {
	repo places.PlaceRepo
}

func NewPlaceUsecase(repo places.PlaceRepo) *PlaceUsecaseImpl {
	return &PlaceUsecaseImpl{repo: repo}
}

func (i *PlaceUsecaseImpl) GetPlaces(ctx context.Context, limit, offset int) ([]models.GetPlace, error) {
	places, err := i.repo.GetPlaces(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return places, nil
}

func (i *PlaceUsecaseImpl) CreatePlace(ctx context.Context, place models.CreatePlace) error {

	return i.repo.CreatePlace(ctx, place)
}

func (i *PlaceUsecaseImpl) UpdatePlace(ctx context.Context, place models.UpdatePlace) error {
	return i.repo.UpdatePlace(ctx, place)
}

func (i *PlaceUsecaseImpl) DeletePlace(ctx context.Context, id uint) error {
	return i.repo.DeletePlace(ctx, id)
}

func (i *PlaceUsecaseImpl) GetPlace(ctx context.Context, id uint) (models.GetPlace, error) {
	return i.repo.GetPlace(ctx, id)
}

func (i *PlaceUsecaseImpl) SearchPlaces(ctx context.Context, name string, category, city, limit, offset int) ([]models.GetPlace, error) {
	log.Println("usecase place name, category, city", name, category, city)
	return i.repo.SearchPlaces(ctx, name, category, city, limit, offset)
}

func (i *PlaceUsecaseImpl) GetPlacesByCategory(ctx context.Context, category string, limit, offset int) ([]models.GetPlace, error) {
	return i.repo.GetPlacesByCategory(ctx, category, limit, offset)
}
