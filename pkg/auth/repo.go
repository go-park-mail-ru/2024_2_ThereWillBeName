package auth

import "database/sql"

type Repo interface {
	getPlaces() ([]Place, error)
}

type Repository struct {
	connectStr string
}

func NewRepository(c string) *Repository {
	return &Repository{connectStr: c}
}

func (r *Repository) getPlaces() ([]Place, error) {

	db, err := sql.Open("postgres", r.connectStr)
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
