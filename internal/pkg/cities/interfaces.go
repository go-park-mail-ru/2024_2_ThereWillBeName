package cities

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

type CityUsecase interface {
	CreateCity(ctx context.Context, city models.City) error
	GetCity(ctx context.Context, id int) (models.City, error)
	GetCities(ctx context.Context, limit, offset int) ([]models.City, error)
	UpdateCity(ctx context.Context, city models.City) error
	DeleteCity(ctx context.Context, id int) error
}

type CityREpo interface {
	CreateCity(ctx context.Context, city models.City) error
	GetCity(ctx context.Context, id int) (models.City, error)
	GetCities(ctx context.Context, limit, offset int) ([]models.City, error)
	UpdateCity(ctx context.Context, city models.City) error
	DeleteCity(ctx context.Context, id int) error
}
