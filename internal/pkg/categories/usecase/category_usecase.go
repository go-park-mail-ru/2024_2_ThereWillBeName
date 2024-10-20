package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/categories/repo"
	"context"
)

type CategoriesUsecase struct {
	repo *repo.CategoriesRepository
}

func NewCategoriesUsecase(repo *repo.CategoriesRepository) *CategoriesUsecase {
	return &CategoriesUsecase{repo}
}

func (c *CategoriesUsecase) CreateCategory(ctx context.Context, category models.Category) error {
	return c.repo.CreateCategory(ctx, category)
}

func (c *CategoriesUsecase) GetCategory(ctx context.Context, id int) (models.Category, error) {
	return c.repo.GetCategory(ctx, id)
}

func (c *CategoriesUsecase) DeleteCategory(ctx context.Context, id int) error {
	return c.repo.DeleteCategory(ctx, id)
}

func (c *CategoriesUsecase) UpdateCategory(ctx context.Context, category models.Category) error {
	return c.repo.UpdateCategory(ctx, category)
}

func (c *CategoriesUsecase) GetCategories(ctx context.Context, limit, offset int) ([]models.Category, error) {
	return c.repo.GetCategories(ctx, limit, offset)
}
