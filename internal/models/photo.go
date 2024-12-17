package models

type Photo struct {
	Path string `json:"path"`
}

//easyjson:json
type PhotoList []Photo
