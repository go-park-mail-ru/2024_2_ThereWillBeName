package search

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	"fmt"
)

type SearchRepository struct {
	db *sql.DB
}

func NewSearchRepository(db *sql.DB) *SearchRepository {
	return &SearchRepository{db: db}
}

func (r *SearchRepository) SearchCitiesAndPlacesBySubString(ctx context.Context, query string) ([]models.SearchResult, error) {
	queryStr := `
        SELECT id, name, 'city' AS type
        FROM city
        WHERE name ILIKE '%' || $1 || '%'
        UNION ALL
        SELECT id, name, 'place' AS type
        FROM place
        WHERE name ILIKE '%' || $1 || '%'
    `

	rows, err := r.db.QueryContext(ctx, queryStr, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute search query: %w", models.ErrInternal)
	}
	defer rows.Close()

	var searchResults []models.SearchResult
	for rows.Next() {
		var item models.SearchResult
		if err := rows.Scan(&item.ID, &item.Name, &item.Type); err != nil {
			return nil, fmt.Errorf("failed to scan search row: %w", models.ErrInternal)
		}
		searchResults = append(searchResults, item)
	}

	if len(searchResults) == 0 {
		return nil, fmt.Errorf("no results found matching query %q: %w", query, models.ErrNotFound)
	}

	return searchResults, nil
}
