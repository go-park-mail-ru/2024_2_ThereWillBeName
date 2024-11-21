package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/categories"
	"context"
)

type CategoriesUsecase struct {
	repo categories.CategoriesRepository
}

func NewCategoriesUsecase(repo categories.CategoriesRepository) *CategoriesUsecase {
	return &CategoriesUsecase{repo: repo}
}

func (u *CategoriesUsecase) GetCategories(ctx context.Context, limit, offset int) ([]models.Category, error) {
	return u.repo.GetCategories(ctx, limit, offset)
}
