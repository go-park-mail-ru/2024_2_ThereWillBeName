package models

import (
	"2024_2_ThereWillBeName/internal/validator"
	"time"
)

type Trip struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CityID      uint      `json:"city_id"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
	Private     bool      `json:"private_trip"`
	Photos      []string  `json:"photos"`
	CreatedAt   time.Time `json:"created_at"`
}

//easyjson:json
type TripList []Trip

type AddPlaceRequest struct {
	PlaceID uint `json:"place_id"`
}

type CreateSharingLinkResponse struct {
	URL string `json:"url"`
}

type TripResponse struct {
	Trip  `json:"trip"`
	Users []UserProfile `json:"users"`
}

func ValidateTrip(v *validator.Validator, trip *Trip) {
	v.Check(trip.Name != "", "name", "must be provided")
	v.Check(len(trip.Name) <= 255, "name", "must not be more than 255 symbols")
	v.Check(trip.UserID != 0, "user id", "must be provided")
	v.Check(len(trip.Description) <= 255, "description", "must not be more than 255 symbols")
	v.Check(trip.CityID != 0, "city id", "must be provided")

	startDate, _ := time.Parse("2006-01-02", trip.StartDate)
	endDate, _ := time.Parse("2006-01-02", trip.EndDate)

	now := time.Now()
	oneMonthAgo := now.AddDate(0, -1, 0)
	twoYearsAhead := now.AddDate(2, 0, 0)
	v.Check(startDate.Before(twoYearsAhead), "start date", "must not be too late")
	v.Check(startDate.After(oneMonthAgo), "start date", "must not be too early")
	v.Check(endDate.Before(twoYearsAhead), "end date", "must not be too late")
	v.Check(endDate.After(oneMonthAgo), "end date", "must not be too early")
	v.Check(endDate.After(startDate), "dates", "start must not be later than end")
}
