package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/dblogger"
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type PlaceRepository struct {
	db *dblogger.DB
}

func NewPLaceRepository(db *dblogger.DB) *PlaceRepository {
	return &PlaceRepository{db: db}
}

func (r *PlaceRepository) GetPlaces(ctx context.Context, limit, offset int) ([]models.GetPlace, error) {
	query := "SELECT p.id, p.name, p.image_path, p.description, p.rating, p.number_of_reviews, p.address, p.phone_number, p.latitude, p.longitude, c.name AS city_name, ARRAY_AGG(ca.name) AS categories FROM place p JOIN city c ON p.city_id = c.id JOIN place_category pc ON p.id = pc.place_id JOIN category ca ON pc.category_id = ca.id GROUP BY p.id, c.name ORDER BY p.id LIMIT $1 OFFSET $2"
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("couldn't get attractions: %w", err)
	}
	defer rows.Close()
	var places []models.GetPlace
	for rows.Next() {
		var place models.GetPlace
		err := rows.Scan(&place.ID, &place.Name, &place.ImagePath, &place.Description, &place.Rating, &place.NumberOfReviews, &place.Address, &place.PhoneNumber, &place.Latitude, &place.Longitude, &place.City, pq.Array(&place.Categories))
		if err != nil {
			return nil, fmt.Errorf("couldn't unmarshal list of attractions: %w", err)
		}
		places = append(places, place)
	}
	return places, nil
}

func (r *PlaceRepository) GetPlace(ctx context.Context, id uint) (models.GetPlace, error) {
	var place models.GetPlace
	query := "SELECT p.id, p.name, p.image_path, p.description, p.rating, p.number_of_reviews, p.address, p.phone_number, p.latitude, p.longitude, c.name AS city_name, ARRAY_AGG(ca.name) AS categories FROM place p JOIN city c ON p.city_id = c.id JOIN place_category pc ON p.id = pc.place_id JOIN category ca ON pc.category_id = ca.id WHERE p.id = $1 GROUP BY p.id, c.name ORDER BY p.id"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&place.ID, &place.Name, &place.ImagePath, &place.Description, &place.Rating, &place.NumberOfReviews, &place.Address, &place.PhoneNumber, &place.Latitude, &place.Longitude, &place.City, pq.Array(&place.Categories))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.GetPlace{}, fmt.Errorf("place not found")
		}
		return models.GetPlace{}, fmt.Errorf("couldn't get place by name: %w", err)
	}
	return place, nil
}

func (r *PlaceRepository) SearchPlaces(ctx context.Context, name string, category, city, filterType, limit, offset int) ([]models.GetPlace, error) {
	var places []models.GetPlace

	query := `
		SELECT 
			p.id, p.name, p.image_path, p.description, p.rating, p.number_of_reviews,
			p.address, p.phone_number, c.name AS city_name, 
			ARRAY_AGG(ca.name) AS categories 
		FROM 
			place p 
		JOIN 
			city c ON p.city_id = c.id 
		JOIN 
			place_category pc ON p.id = pc.place_id 
		JOIN 
			category ca ON pc.category_id = ca.id 
		WHERE 
			p.name LIKE '%' || $1 || '%'`

	args := []interface{}{name}

	if category > 0 {
		query += " AND pc.category_id = $2"
		args = append(args, category)
	}

	if city > 0 {
		query += " AND p.city_id = $" + fmt.Sprint(len(args)+1)
		args = append(args, city)
	}

	query += `
		GROUP BY 
			p.id, c.name`

	switch filterType {
	case 1:
		query += " ORDER BY p.rating DESC"
	case 2:
		query += " ORDER BY p.number_of_reviews DESC"
	default:
		query += " ORDER BY p.id"
	}

	query += `
		LIMIT $` + fmt.Sprint(len(args)+1) + ` OFFSET $` + fmt.Sprint(len(args)+2)
	args = append(args, limit, offset)

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("couldn't prepare query: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("couldn't get places: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var place models.GetPlace
		err := rows.Scan(&place.ID, &place.Name, &place.ImagePath, &place.Description, &place.Rating, &place.NumberOfReviews, &place.Address, &place.PhoneNumber, &place.City, pq.Array(&place.Categories))
		if err != nil {
			return nil, fmt.Errorf("couldn't unmarshal list of places: %w", err)
		}
		places = append(places, place)
	}

	return places, nil
}

func (r *PlaceRepository) GetPlacesByCategory(ctx context.Context, category string, limit, offset int) ([]models.GetPlace, error) {
	query := `SELECT 
    		p.id, p.name, p.image_path, p.description, p.rating, p.number_of_reviews, p.address, p.phone_number, p.latitude, p.longitude,
    		c.name AS city_name,
    		ARRAY_AGG(ca.name) AS categories
			FROM place p 
			JOIN city c 
			ON p.city_id = c.id 
			JOIN place_category pc
			ON p.id = pc.place_id
			JOIN category ca 
			ON pc.category_id = ca.id 
			WHERE p.id IN (
			               	SELECT p.id 
            				FROM place p 
            				JOIN place_category pc 
            				ON p.id = pc.place_id 
            				JOIN category ca 
            				ON pc.category_id = ca.id 
            				WHERE ca.name = $1)
			GROUP BY p.id, c.name
			ORDER BY p.id 
			LIMIT $2 
			OFFSET $3`
	rows, err := r.db.QueryContext(ctx, query, category, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("couldn't get attractions by category: %w", err)
	}
	defer rows.Close()
	var places []models.GetPlace
	for rows.Next() {
		var place models.GetPlace
		err := rows.Scan(&place.ID, &place.Name, &place.ImagePath, &place.Description, &place.Rating, &place.NumberOfReviews, &place.Address, &place.PhoneNumber, &place.Latitude, &place.Longitude, &place.City, pq.Array(&place.Categories))
		if err != nil {
			return nil, fmt.Errorf("couldn't unmarshal list of attractions: %w", err)
		}
		places = append(places, place)
	}
	return places, nil
}
