package auth

import "database/sql"

type Repo interface {
	getPlaces() ([]Place, error)
}

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) getPlaces() ([]Place, error) {
	connStr := "user=postgres password=mypassword host=localhost port=5432 dbname=landmarks sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM places")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	places := []Place{}
	for rows.Next() {
		var place Place
		err := rows.Scan(&place.ID, &place.Name, &place.Image)
		if err != nil {
			return nil, err
		}
		places = append(places, place)
	}
	return places, nil
}
