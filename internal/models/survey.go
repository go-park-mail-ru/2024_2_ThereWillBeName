package models

type Survey struct {
	Id         uint   `json:"id"`
	SurveyText string `json:"survey_text"`
	MaxRating  int    `json:"max_rating"`
}

type SurveyResponse struct {
	SurveyId uint `json:"survey_id"`
	UserId   uint `json:"user_id"`
	Rating   int  `json:"rating"`
}

type SurveyStatsBySurvey struct {
	SurveyId     uint        `json:"survey_id"`
	SurveyText   string      `json:"survey_text"`
	AvgRating    float64     `json:"avg_rating"`
	RatingsCount map[int]int `json:"ratings_count"`
}

type UserSurveyStats struct {
	SurveyId   uint   `json:"survey_id"`
	SurveyText string `json:"survey_text"`
	Answered   bool   `json:"answered"`
}
