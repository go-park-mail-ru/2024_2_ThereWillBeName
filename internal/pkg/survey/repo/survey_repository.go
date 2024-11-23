package repo

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	"github.com/lib/pq"
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
	query := `INSERT INTO user_survey (survey_id, user_id, rating) VALUES  (&1, &2, &3) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, response.SurveyId, response.UserId, response.Rating)
	if err != nil {
		return fmt.Errorf("could not create survey: %w", err)
	}
	return nil
}

func (r *SurveyRepository) GetSurveyStatsBySurveyId(ctx context.Context, surveyId uint) (models.SurveyStatsBySurvey, error) {
	query = `SELECT `
}
