package categories

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

type CategoryUsecase interface {
	CreateCategory(ctx context.Context, category models.Category) error
	GetCategory(ctx context.Context, id int) (models.Category, error)
	GetCategories(ctx context.Context, limit, offset int) ([]models.Category, error)
	UpdateCategory(ctx context.Context, category models.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type CategoryREpo interface {
	CreateCategory(ctx context.Context, category models.Category) error
	GetCategory(ctx context.Context, id int) (models.Category, error)
	GetCategories(ctx context.Context, limit, offset int) ([]models.Category, error)
	UpdateCategory(ctx context.Context, category models.Category) error
	DeleteCategory(ctx context.Context, id int) error
}
