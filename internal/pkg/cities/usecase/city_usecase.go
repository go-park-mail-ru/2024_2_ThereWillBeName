package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/cities/repo"
	"context"
)

type CitiesUsecase struct {
	repo *repo.CitiesRepository
}

func NewCitiesUsecase(repo *repo.CitiesRepository) *CitiesUsecase { return &CitiesUsecase{repo} }

func (c *CitiesUsecase) CreateCity(ctx context.Context, city models.City) error {
	return c.repo.CreateCity(ctx, city)
}

func (c *CitiesUsecase) GetCity(ctx context.Context, id int) (models.City, error) {
	return c.repo.GetCity(ctx, id)
}

func (c *CitiesUsecase) DeleteCity(ctx context.Context, id int) error {
	return c.repo.DeleteCity(ctx, id)
}

func (c *CitiesUsecase) UpdateCity(ctx context.Context, city models.City) error {
	return c.repo.UpdateCity(ctx, city)
}

func (c *CitiesUsecase) GetCities(ctx context.Context, limit, offset int) ([]models.City, error) {
	return c.repo.GetCities(ctx, limit, offset)
}
