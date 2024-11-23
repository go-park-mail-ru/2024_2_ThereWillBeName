package usecase

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/survey"
	"context"
	"errors"
	"fmt"
	"log"
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
	survey, err := u.surveyRepo.GetSurveyById(ctx, surveyId)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return models.Survey{}, fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		return models.Survey{}, fmt.Errorf("internal error: %w", models.ErrInternal)
	}
	return survey, nil
}

func (u *SurveysUseCaseImpl) CreateSurveyResponse(ctx context.Context, response models.SurveyResponse) error {
	survey, err := u.surveyRepo.GetSurveyById(ctx, response.SurveyId)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			log.Println(err)
			return fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		log.Println(err)
		return fmt.Errorf("internal error: %w", models.ErrInternal)
	}

	if response.Rating > survey.MaxRating {
		log.Println(err)
		return fmt.Errorf("invalid rating: cannot be higher than the maximum rating of %d", survey.MaxRating)
	}

	err = u.surveyRepo.CreateSurveyResponse(ctx, response)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			log.Println(err)
			return fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		log.Println(err)
		return fmt.Errorf("internal error: %w", models.ErrInternal)
	}
	return nil
}

func (u *SurveysUseCaseImpl) GetSurveyStatsBySurveyId(ctx context.Context, surveyId uint) (models.SurveyStatsBySurvey, error) {
	survey, err := u.surveyRepo.GetSurveyStatsBySurveyId(ctx, surveyId)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			log.Println(err)
			return models.SurveyStatsBySurvey{}, fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		log.Println(err)
		return models.SurveyStatsBySurvey{}, fmt.Errorf("internal error: %w", models.ErrInternal)
	}
	return survey, nil
}

func (u *SurveysUseCaseImpl) GetSurveyStatsByUserId(ctx context.Context, userId uint) ([]models.UserSurveyStats, error) {
	stats, err := u.surveyRepo.GetSurveyStatsByUserId(ctx, userId)
	if err != nil {
		log.Println(err)
		if errors.Is(err, models.ErrNotFound) {
			return nil, fmt.Errorf("invalid request: %w", models.ErrNotFound)
		}
		return nil, fmt.Errorf("internal error: %w", models.ErrInternal)
	}
	return stats, nil
}
