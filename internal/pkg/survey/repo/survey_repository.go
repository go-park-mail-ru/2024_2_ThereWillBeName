package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"log"
)

type SurveyRepository struct {
	db *sql.DB
}

func NewPLaceRepository(db *sql.DB) *SurveyRepository {
	return &SurveyRepository{db: db}
}

func (r *SurveyRepository) GetSurveyById(ctx context.Context, surveyId uint) (models.Survey, error) {
	query := `SELECT id, survey_text, max_rating FROM survey WHERE id=$1`
	row := r.db.QueryRowContext(ctx, query, surveyId)

	var survey models.Survey
	err := row.Scan(&survey.Id, &survey.SurveyText, &survey.MaxRating)
	if err != nil {
		return models.Survey{}, fmt.Errorf("could not retrieve survey: %w", err)
	}
	return survey, nil
}

func (r *SurveyRepository) CreateSurveyResponse(ctx context.Context, response models.SurveyResponse) error {
	query := `INSERT INTO user_survey (survey_id, user_id, rating) VALUES  ($1, $2, $3)`
	log.Println(response.SurveyId, response.UserId, response.Rating)
	_, err := r.db.ExecContext(ctx, query, response.SurveyId, response.UserId, response.Rating)
	if err != nil {
		return fmt.Errorf("could not create survey: %w", err)
	}
	return nil
}

func (r *SurveyRepository) GetSurveyStatsBySurveyId(ctx context.Context, surveyId uint) (models.SurveyStatsBySurvey, error) {
	query := `SELECT 
    	s.id AS survey_id,
    	s.survey_text,
    	us.rating,
    	COUNT(us.user_id) AS count_of_users
	FROM 
    	survey s
	LEFT JOIN 
    	user_survey us ON s.id = us.survey_id
	WHERE 
		s.id = $1
	GROUP BY 
		s.id, s.survey_text, us.rating
	ORDER BY 
		s.id, us.rating; `

	row := r.db.QueryRowContext(ctx, query, surveyId)
	var surveyStats models.SurveyStatsBySurvey
	err := row.Scan(&surveyStats.SurveyId, &surveyStats.SurveyText, &surveyStats.AvgRating, &surveyStats.RatingsCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.SurveyStatsBySurvey{}, fmt.Errorf("trip not found: %w", models.ErrNotFound)
		}
		return models.SurveyStatsBySurvey{}, fmt.Errorf("failed to scan survey stats row: %w", models.ErrInternal)
	}
	return surveyStats, nil
}

func (r *SurveyRepository) GetSurveyStatsByUserId(ctx context.Context, userId uint) ([]models.UserSurveyStats, error) {
	query := `SELECT 
    s.id AS survey_id,
    s.survey_text,
    CASE 
        WHEN us.user_id IS NOT NULL THEN TRUE
        ELSE FALSE
    END AS participated
FROM 
    survey s
LEFT JOIN 
    user_survey us ON s.id = us.survey_id AND us.user_id = 1
ORDER BY 
    s.id;`
	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user stats: %w", models.ErrInternal)
	}
	defer rows.Close()
	var userStats []models.UserSurveyStats
	for rows.Next() {
		var userStat models.UserSurveyStats
		if err := rows.Scan(&userStat.SurveyId, &userStat.SurveyText, &userStat.Answered); err != nil {
			return nil, fmt.Errorf("failed to scan user stat row: %w", models.ErrInternal)
		}
		userStats = append(userStats, userStat)
	}
	if len(userStats) == 0 {
		return nil, fmt.Errorf("no user stats found: %w", models.ErrNotFound)
	}

	return userStats, nil
}
