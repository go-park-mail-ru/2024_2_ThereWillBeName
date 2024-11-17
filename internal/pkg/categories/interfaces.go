package categories

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

type CategoriesUsecase interface {
	GetCategories(ctx context.Context, limit, offset int) ([]models.Category, error)
}

type CategoriesRepository interface {
	GetCategories(ctx context.Context, limit, offset int) ([]models.Category, error)
}
