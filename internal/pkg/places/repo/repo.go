package repo

import (
	"TripAdvisor/internal/models"
	"database/sql"
)

type PlaceRepository struct {
	db *sql.DB
}

func NewRepository() *PlaceRepository {
	return &PlaceRepository{}
}

func (r *PlaceRepository) GetPlaces() ([]models.Place, error) {
	//rows, err := r.db.Query("SELECT * FROM places")
	//if err != nil {
	//	return nil, err
	//}
	//defer rows.Close()
	//places := []models.Place{}
	//for rows.Next() {
	//	var place models.Place
	//	err := rows.Scan(&place.ID, &place.Name, &place.Image)
	//	if err != nil {
	//		return nil, err
	//	}
	//	places = append(places, place)
	//}
	//return places, nil

	//file, err := os.Open("./places.json")
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//data, err := io.ReadAll(file)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//defer file.Close()
	//images := make([]models.Place, 0)
	//json.Unmarshal(data, &images)
	var places []models.Place
	places = append(places, models.Place{ID: 1, Name: "cdcs", Image: "cdcs"})
	return places, nil
}
