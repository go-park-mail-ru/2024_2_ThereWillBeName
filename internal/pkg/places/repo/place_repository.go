package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	_ "embed"
)

type PlaceRepository struct {
	db *sql.DB
}

func NewPLaceRepository(db *sql.DB) *PlaceRepository {
	return &PlaceRepository{db: db}
}

//go:embed places.json
var jsonFileData []byte

func (r *PlaceRepository) GetPlaces(ctx context.Context) ([]models.Place, error) {
	rows, err := r.db.Query("SELECT * FROM places")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	places := []models.Place{}
	for rows.Next() {
		var place models.Place
		err := rows.Scan(&place.ID, &place.Name, &place.Image)
		if err != nil {
			return nil, err
		}
		places = append(places, place)
	}
	return places, nil
}

func (r *PlaceRepository) CreatePlace(ctx context.Context, place models.Place) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO places (name, image, description) VALUES ($1, $2, $3)", place.Name, place.Image, place.Description)
	return err
}

func (r *PlaceRepository) ReadPlace(ctx context.Context, name string) (models.Place, error) {
	var dataPlace models.Place
	row := r.db.QueryRowContext(ctx, "SELECT * FROM places where name = $1", name)
	err := row.Scan(&dataPlace.ID, &dataPlace.Name, &dataPlace.Image, &dataPlace.Description)
	return dataPlace, err
}

func (r *PlaceRepository) UpdatePlace(ctx context.Context, place models.Place) error {
	_, err := r.db.ExecContext(ctx, "UPDATE places SET name = $1, image = $2, description = $3", place.Name, place.Image, place.Description)
	return err
}

func (r *PlaceRepository) DeletePlace(ctx context.Context, name string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM places WHERE name=$1", name)
	return err
}
