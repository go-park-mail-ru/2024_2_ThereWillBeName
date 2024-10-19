package models

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
