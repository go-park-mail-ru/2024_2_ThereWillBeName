package survey

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

type SurveysUsecase interface {
	GetSurveyById(ctx context.Context, surveyID uint) (models.Survey, error)
	CreateSurveyResponse(ctx context.Context, response models.SurveyResponse) error
	GetSurveyStatsBySurveyId(ctx context.Context, surveyID uint) (models.SurveyStatsBySurvey, error)
	GetSurveyStatsByUserId(ctx context.Context, userID uint) ([]models.UserSurveyStats, error)
}

type SurveysRepo interface {
	GetSurveyById(ctx context.Context, surveyID uint) (models.Survey, error)
	CreateSurveyResponse(ctx context.Context, response models.SurveyResponse) error
	GetSurveyStatsBySurveyId(ctx context.Context, surveyID uint) (models.SurveyStatsBySurvey, error)
	GetSurveyStatsByUserId(ctx context.Context, userID uint) ([]models.UserSurveyStats, error)
}
