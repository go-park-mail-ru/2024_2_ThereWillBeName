package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	search "2024_2_ThereWillBeName/internal/pkg/search/repo"
	"context"
)

type SearchUsecaseImpl struct {
	repo search.SearchRepository
}

func NewSearchUsecase(repo search.SearchRepository) *SearchUsecaseImpl {
	return &SearchUsecaseImpl{repo: repo}
}

func (uc *SearchUsecaseImpl) Search(ctx context.Context, query string) ([]models.SearchResult, error) {
	var results []models.SearchResult

	results, err := uc.repo.SearchCitiesAndPlacesBySubString(ctx, query)
	if err != nil {
		return nil, err
	}

	return results, nil
}
