package survey

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

type SurveysUsecase interface {
	GetSurveyById(ctx context.Context, surveyId uint) (models.Survey, error)
	CreateSurveyResponse(ctx context.Context, response models.SurveyResponse) error
	GetSurveyStatsBySurveyId(ctx context.Context, surveyId uint) (models.SurveyStatsBySurvey, error)
	GetSurveyStatsByUserId(ctx context.Context, userId uint) ([]models.UserSurveyStats, error)
}

type SurveysRepo interface {
	GetSurveyById(ctx context.Context, surveyId uint) (models.Survey, error)
	CreateSurveyResponse(ctx context.Context, response models.SurveyResponse) error
	GetSurveyStatsBySurveyId(ctx context.Context, surveyId uint) (models.SurveyStatsBySurvey, error)
	GetSurveyStatsByUserId(ctx context.Context, userId uint) ([]models.UserSurveyStats, error)
}
