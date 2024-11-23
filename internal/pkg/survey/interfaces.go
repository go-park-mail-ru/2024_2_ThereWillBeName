package survey

import (
	"2024_2_ThereWillBeName/internal/models"
	"context"
)

type SurveysUsecase interface {
	GetSurveyByID(ctx context.Context, surveyID uint) (models.Survey, error)
	SubmitResponse(ctx context.Context, response models.SurveyResponse) error
	GetSurveyStatsBySurveyID(ctx context.Context, surveyID uint) (models.SurveyStatsBySurvey, error)
	GetSurveyStatsByUserID(ctx context.Context, userID uint) ([]models.UserSurveyStats, error)
}

type SurveysRepo interface {
	GetSurveyByID(ctx context.Context, surveyID uint) (models.Survey, error)
	SubmitResponse(ctx context.Context, response models.SurveyResponse) error
	GetSurveyStatsBySurveyID(ctx context.Context, surveyID uint) (models.SurveyStatsBySurvey, error)
	GetSurveyStatsByUserID(ctx context.Context, userID uint) ([]models.UserSurveyStats, error)
}
