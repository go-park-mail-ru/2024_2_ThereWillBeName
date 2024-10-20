package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type CategoriesRepository struct {
	db *sql.DB
}

func NewCategoriesRepository(db *sql.DB) *CategoriesRepository { return &CategoriesRepository{db} }

func (c *CategoriesRepository) CreateCategory(ctx context.Context, category models.Category) error {
	query := "INSERT INTO categories (name) VALUES ($1)"
	log.Println(query, category.Name)
	_, err := c.db.ExecContext(ctx, query, category.Name)
	return err
}

func (c *CategoriesRepository) GetCategory(ctx context.Context, id int) (models.Category, error) {
	var category models.Category
	query := "SELECT id, name FROM categories WHERE id = $1"
	row := c.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&category.ID, &category.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Category{}, fmt.Errorf("category not found: %w", models.ErrNotFound)
		}
		return models.Category{}, fmt.Errorf("couldn't get category by name: %w", err)
	}
	return category, nil
}

func (c *CategoriesRepository) UpdateCategory(ctx context.Context, category models.Category) error {
	query := "UPDATE categories SET name = $1 WHERE id = $2"

	result, err := c.db.ExecContext(ctx, query, category.Name, category.ID)
	if err != nil {
		return fmt.Errorf("couldn't update category: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("couldn't get number of rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated: %w", models.ErrNotFound)
	}
	return nil
}

func (c *CategoriesRepository) DeleteCategory(ctx context.Context, id int) error {
	query := "DELETE FROM categories WHERE id = $1"
	result, err := c.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("couldn't delete category: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("couldn't get number of rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted: %w", models.ErrNotFound)
	}
	return nil
}

func (c *CategoriesRepository) GetCategories(ctx context.Context, limit, offset int) ([]models.Category, error) {
	var categories []models.Category
	query := "SELECT id, name FROM categories LIMIT $1 OFFSET $2"
	rows, err := c.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("couldn't get categories: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var category models.Category
		log.Println(rows)
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, fmt.Errorf("couldn't unmarshal category by name: %w", err)
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("couldn't  unmarshal categories: %w", err)
	}
	log.Println(categories)
	return categories, nil
}
