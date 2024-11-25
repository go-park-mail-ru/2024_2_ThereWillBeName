package search

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

type SearchUsecase interface {
	Search(ctx context.Context, query string) ([]models.SearchResult, error)
}

type SearchRepo interface {
	SearchCitiesAndPlacesBySubString(ctx context.Context, query string) ([]models.SearchResult, error)
}
