package cities

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

type CitiesUsecase interface {
	SearchCitiesByName(ctx context.Context, query string) ([]models.City, error)
	SearchCityByID(ctx context.Context, id uint) (models.City, error)
}

type CitiesRepo interface {
	SearchCitiesByName(ctx context.Context, query string) ([]models.City, error)
	SearchCityByID(ctx context.Context, id uint) (models.City, error)
	// SearchCitiesBySubString(ctx context.Context, query string) ([]models.SearchItem, error)
}
