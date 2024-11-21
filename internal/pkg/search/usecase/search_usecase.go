package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	search "2024_2_ThereWillBeName/internal/pkg/search"
	"context"
)

type SearchUsecaseImpl struct {
	repo search.SearchRepo
}

func NewSearchUsecase(repo search.SearchRepo) *SearchUsecaseImpl {
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
