package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/cities"
	"context"
	"errors"
	"fmt"
)

type CitiesUsecaseImpl struct {
	citiesRepo cities.CitiesRepo
}

func NewCitiesUsecase(repo cities.CitiesRepo) *CitiesUsecaseImpl {
	return &CitiesUsecaseImpl{
		citiesRepo: repo,
	}
}

func (u *CitiesUsecaseImpl) SearchCitiesByName(ctx context.Context, query string) ([]models.City, error) {
	citiesFound, err := u.citiesRepo.SearchCitiesByName(ctx, query)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, fmt.Errorf("could not find cities: %w", models.ErrNotFound)
		}
		return nil, fmt.Errorf("internal error: %w", models.ErrInternal)
	}
	return citiesFound, nil
}

func (u *CitiesUsecaseImpl) SearchCityByID(ctx context.Context, id uint) (models.City, error) {
	city, err := u.citiesRepo.SearchCityByID(ctx, id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return models.City{}, fmt.Errorf("could not find city: %w", models.ErrNotFound)
		}
		return models.City{}, fmt.Errorf("internal error^ %w", models.ErrInternal)

	}
	return city, nil
}
