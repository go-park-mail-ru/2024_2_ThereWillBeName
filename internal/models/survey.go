package models

type Survey struct {
	ID         uint   `json:"id"`
	SurveyText string `json:"survey_text"`
	MaxRating  int    `json:"max_rating"`
}

type SurveyResponse struct {
	SurveyID uint `json:"survey_id"`
	UserID   uint `json:"user_id"`
	Rating   int  `json:"rating"`
}

type SurveyStatsBySurvey struct {
	SurveyID     uint        `json:"survey_id"`
	SurveyText   string      `json:"survey_text"`
	AvgRating    float64     `json:"avg_rating"`
	RatingsCount map[int]int `json:"ratings_count"`
}

// Статистика по ID пользователя
type UserSurveyStats struct {
	SurveyID   uint   `json:"survey_id"`
	SurveyText string `json:"survey_text"`
	Answered   bool   `json:"answered"`
}
