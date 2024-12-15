package models

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"Name"`
}

//easyjson:json
type CategoryList []Category
