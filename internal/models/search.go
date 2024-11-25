package models

type SearchResult struct {
	Name string `json:"name"`
	Id   uint   `json:"id"`
	Type string `json:"type"` //"city" или "place"
}
