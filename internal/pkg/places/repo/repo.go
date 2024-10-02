package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	_ "embed"
	"encoding/json"
	"log"
)

type PlaceRepository struct {
}

func NewPLaceRepository() *PlaceRepository {
	return &PlaceRepository{}
}

//go:embed places.json
var jsonFileData []byte

func (r *PlaceRepository) GetPlaces(ctx context.Context) ([]models.Place, error) {
	images := make([]models.Place, 0)
	err := json.Unmarshal(jsonFileData, &images)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return images, nil
}
