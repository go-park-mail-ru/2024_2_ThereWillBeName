package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"context"
	"fmt"
)

type CategoriesRepo struct {
	db *dblogger.DB
}

func NewCategoriesRepo(db *dblogger.DB) *CategoriesRepo {
	return &CategoriesRepo{db: db}
}

func (r *CategoriesRepo) GetCategories(ctx context.Context, limit, offset int) ([]models.Category, error) {
	query := "SELECT id, name FROM category ORDER BY id LIMIT $1 OFFSET $2"
	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("couldn't get categories: %w", err)
	}
	defer rows.Close()
	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, fmt.Errorf("Couldn't unmarshal list of categories: %w", err)
		}
		categories = append(categories, category)
	}
	return categories, nil
}
