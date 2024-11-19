package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type PlaceRepository struct {
	db *sql.DB
}

func NewPLaceRepository(db *sql.DB) *PlaceRepository {
	return &PlaceRepository{db: db}
}

func (r *PlaceRepository) GetPlaces(ctx context.Context, limit, offset int) ([]models.GetPlace, error) {
	query := "SELECT p.id, p.name, p.image_path, p.description, p.rating, p.address, p.phone_number, p.latitude, p.longitude, c.name AS city_name, ARRAY_AGG(ca.name) AS categories FROM place p JOIN city c ON p.city_id = c.id JOIN place_category pc ON p.id = pc.place_id JOIN category ca ON pc.category_id = ca.id GROUP BY p.id, c.name ORDER BY p.id LIMIT $1 OFFSET $2"
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("couldn't get places: %w", err)
	}
	defer rows.Close()
	var places []models.GetPlace
	for rows.Next() {
		var place models.GetPlace
		err := rows.Scan(&place.ID, &place.Name, &place.ImagePath, &place.Description, &place.Rating, &place.Address, &place.PhoneNumber, &place.Latitude, &place.Longitude, &place.City, pq.Array(&place.Categories))
		if err != nil {
			return nil, fmt.Errorf("couldn't unmarshal list of places: %w", err)
		}
		places = append(places, place)
	}
	return places, nil
}

func (r *PlaceRepository) CreatePlace(ctx context.Context, place models.CreatePlace) error {
	query := "INSERT INTO place (name, image_path, description, rating, address, city_id, phone_number, latitude, longitude) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
	var id int
	err := r.db.QueryRowContext(ctx, query, place.Name, place.ImagePath, place.Description, place.Rating, place.Address, place.CityId, place.PhoneNumber, place.Latitude, place.Longitude).Scan(&id)
	if err != nil {
		return fmt.Errorf("coldn't create place: %w", err)
	}
	for _, categoryId := range place.CategoriesId {
		query = "INSERT INTO place_category (place_id, category_id) VALUES ($1, $2)"
		result, err := r.db.ExecContext(ctx, query, id, categoryId)
		if err != nil {
			return fmt.Errorf("coldn't create place_category: %w", err)
		}
		if _, err = result.RowsAffected(); err != nil {
			return fmt.Errorf("couldn't get number of rows affected: %w", err)
		}
	}
	return nil
}

func (r *PlaceRepository) GetPlace(ctx context.Context, id uint) (models.GetPlace, error) {
	var place models.GetPlace
	query := "SELECT p.id, p.name, p.image_path, p.description, p.rating, p.address, p.phone_number, p.latitude, p.longitude, c.name AS city_name, ARRAY_AGG(ca.name) AS categories FROM place p JOIN city c ON p.city_id = c.id JOIN place_category pc ON p.id = pc.place_id JOIN category ca ON pc.category_id = ca.id WHERE p.id = $1 GROUP BY p.id, c.name ORDER BY p.id"
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&place.ID, &place.Name, &place.ImagePath, &place.Description, &place.Rating, &place.Address, &place.PhoneNumber, &place.Latitude, &place.Longitude, &place.City, pq.Array(&place.Categories))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.GetPlace{}, fmt.Errorf("place not found")
		}
		return models.GetPlace{}, fmt.Errorf("couldn't get place by name: %w", err)
	}
	return place, nil
}

func (r *PlaceRepository) UpdatePlace(ctx context.Context, place models.UpdatePlace) error {
	query := "UPDATE place SET name = $1, image_path = $2, description = $3, rating = $4, address = $5, phone_number = $6, latitude = $7, longitude = $8, updated_at = NOW() WHERE id=$7"
	result, err := r.db.ExecContext(ctx, query, place.Name, place.ImagePath, place.Description, place.Rating, place.Address, place.PhoneNumber, place.Latitude, place.Longitude, place.ID)
	if err != nil {
		return fmt.Errorf("couldn't update place: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("couldn't get number of rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	query = "DELETE FROM place_category WHERE place_id=$1"
	result, err = r.db.ExecContext(ctx, query, place.ID)
	if err != nil {
		return fmt.Errorf("couldn't delete place_category: %w", err)
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("couldn't get number of rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted")
	}

	for _, categoryId := range place.CategoriesId {
		query = "INSERT INTO place_category (place_id, category_id) VALUES ($1, $2)"
		result, err := r.db.ExecContext(ctx, query, place.ID, categoryId)
		if err != nil {
			return fmt.Errorf("coldn't create place_category: %w", err)
		}
		if _, err = result.RowsAffected(); err != nil {
			return fmt.Errorf("couldn't get number of rows affected: %w", err)
		}
	}
	return nil
}

func (r *PlaceRepository) DeletePlace(ctx context.Context, id uint) error {
	query := "DELETE FROM place WHERE id=$1"
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("couldn't delete place: %w", err)
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

func (r *PlaceRepository) SearchPlaces(ctx context.Context, name string, limit, offset int) ([]models.GetPlace, error) {
	var places []models.GetPlace
	query := "SELECT p.id, p.name, p.image_path, p.description, p.rating, p.address, p.phone_number, p.latitude, p.longitude, c.name AS city_name, ARRAY_AGG(ca.name) AS categories FROM place p JOIN city c ON p.city_id = c.id JOIN place_category pc ON p.id = pc.place_id JOIN category ca ON pc.category_id = ca.id WHERE p.name LIKE '%' || $1 || '%' GROUP BY p.id, c.name ORDER BY p.id LIMIT $2 OFFSET $3"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("couldn't prepare query: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, name, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("couldn't get places: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var place models.GetPlace
		err := rows.Scan(&place.ID, &place.Name, &place.ImagePath, &place.Description, &place.Rating, &place.Address, &place.PhoneNumber, &place.Latitude, &place.Longitude, &place.City, pq.Array(&place.Categories))
		if err != nil {
			return nil, fmt.Errorf("couldn't unmarshal list of places: %w", err)
		}
		places = append(places, place)
	}
	return places, nil
}

func (r *PlaceRepository) GetPlacesByCategory(ctx context.Context, category string, limit, offset int) ([]models.GetPlace, error) {
	query := `SELECT 
    		p.id, p.name, p.image_path, p.description, p.rating, p.address, p.phone_number, p.latitude, p.longitude,
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
		return nil, fmt.Errorf("couldn't get places by category: %w", err)
	}
	defer rows.Close()
	var places []models.GetPlace
	for rows.Next() {
		var place models.GetPlace
		err := rows.Scan(&place.ID, &place.Name, &place.ImagePath, &place.Description, &place.Rating, &place.Address, &place.PhoneNumber, &place.Latitude, &place.Longitude, &place.City, pq.Array(&place.Categories))
		if err != nil {
			return nil, fmt.Errorf("couldn't unmarshal list of places: %w", err)
		}
		places = append(places, place)
	}
	return places, nil
}
