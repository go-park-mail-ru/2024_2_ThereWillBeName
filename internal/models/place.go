package models

type Place struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	ImagePath       string `json:"imagePath"`
	Description     string `json:"description"`
	Rating          int    `json:"rating"`
	NumberOfReviews int    `json:"numberOfReviews"`
	Address         string `json:"address"`
	City            string `json:"city"`
	PhoneNumber     string `json:"phoneNumber"`
	Category        string `json:"category"`
}
