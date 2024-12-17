package models

import "time"

type City struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

//easyjson:json
type CityList []City
