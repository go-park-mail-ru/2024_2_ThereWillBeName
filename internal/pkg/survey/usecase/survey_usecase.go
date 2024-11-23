package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/survey"
	"context"
)

type SurveysUseCaseImpl struct {
	surveyRepo survey.SurveysRepo
}

func NewSurveysUsecase(repo survey.SurveysRepo) *SurveysUseCaseImpl {
	return &SurveysUseCaseImpl{
		surveyRepo: repo,
	}
}

func (u *SurveysUseCaseImpl) GetSurveyById(ctx context.Context, surveyId uint) (models.Survey, error) {

}

func (u *SurveysUseCaseImpl) CreateSurveyResponse(ctx context.Context, surveyId uint) (models.Survey, error) {

}

func (u *SurveysUseCaseImpl) GetSurveyStatsBySurveyId(ctx context.Context, surveyId uint) (models.Survey, error) {

}

func (u *SurveysUseCaseImpl) GetSurveyStatsByUserId(ctx context.Context, surveyId uint) (models.Survey, error) {

}
