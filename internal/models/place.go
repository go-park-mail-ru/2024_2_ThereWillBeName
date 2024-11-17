package models

import "2024_2_ThereWillBeName/internal/validator"

type CreatePlace struct {
	Name            string `json:"name"`
	ImagePath       string `json:"imagePath"`
	Description     string `json:"description"`
	Rating          int    `json:"rating"`
	NumberOfReviews int    `json:"numberOfReviews"`
	Address         string `json:"address"`
	CityId          int    `json:"cityId"`
	PhoneNumber     string `json:"phoneNumber"`
	CategoriesId    []int  `json:"categoriesId"`
}

type GetPlace struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	ImagePath       string   `json:"imagePath"`
	Description     string   `json:"description"`
	Rating          int      `json:"rating"`
	NumberOfReviews int      `json:"numberOfReviews"`
	Address         string   `json:"address"`
	City            string   `json:"city"`
	PhoneNumber     string   `json:"phoneNumber"`
	Categories      []string `json:"categories"`
}

type UpdatePlace struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	ImagePath       string `json:"imagePath"`
	Description     string `json:"description"`
	Rating          int    `json:"rating"`
	NumberOfReviews int    `json:"numberOfReviews"`
	Address         string `json:"address"`
	CityId          int    `json:"cityId"`
	PhoneNumber     string `json:"phoneNumber"`
	CategoriesId    []int  `json:"categoriesId"`
}

func ValidateCreatePlace(v *validator.Validator, place *CreatePlace) {
	v.Check(place.Name != "", "name", "must be provided")
	v.Check(len(place.Name) <= 255, "name", "must not be more than 255 symbols")
	v.Check(place.Description != "", "descriprion", "must be provided")
	v.Check(len(place.Description) <= 255, "description", "must not be more than 255 symbols")
	v.Check(place.ImagePath != "", "image path", "must be provided")
	v.Check(len(place.ImagePath) <= 255, "image path", "must not be more than 255 symbols")
	v.Check(place.Address != "", "address", "must be provided")
	v.Check(len(place.Address) <= 255, "address", "must not be more than 255 symbols")
	v.Check(place.CityId != 0, "city id", "must be provided")
}

func ValidateUpdatePlace(v *validator.Validator, place *UpdatePlace) {
	v.Check(place.Name != "", "name", "must be provided")
	v.Check(len(place.Name) <= 255, "name", "must not be more than 255 symbols")
	v.Check(place.Description != "", "descriprion", "must be provided")
	v.Check(len(place.Description) <= 255, "description", "must not be more than 255 symbols")
	v.Check(place.ImagePath != "", "image path", "must be provided")
	v.Check(len(place.ImagePath) <= 255, "image path", "must not be more than 255 symbols")
	v.Check(place.Address != "", "address", "must be provided")
	v.Check(len(place.Address) <= 255, "address", "must not be more than 255 symbols")
	v.Check(place.CityId != 0, "city id", "must be provided")
}
