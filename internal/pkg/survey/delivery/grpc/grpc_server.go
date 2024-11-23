package grpc

import (
	"2024_2_ThereWillBeName/internal/models"
	surveyPkg "2024_2_ThereWillBeName/internal/pkg/survey"
	"2024_2_ThereWillBeName/internal/pkg/survey/delivery/grpc/gen"
	"context"
)

//go:generate protoc -I . proto/survey.proto --go_out=./gen --go-grpc_out=./gen

type GrpcSurveyHandler struct {
	gen.UnimplementedSurveyServiceServer
	uc surveyPkg.SurveysUsecase
}

func NewGrpcSurveyHandler(uc *surveyPkg.SurveysUsecase) *GrpcSurveyHandler {
	return &GrpcSurveyHandler{uc: *uc}
}

func (s *GrpcSurveyHandler) GetSurveyById(ctx context.Context, req *gen.GetSurveyByIdRequest) (*gen.GetSurveyByIdResponce, error) {
	survey, err := s.uc.GetSurveyById(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}
	surveyResponce := &gen.Survey{
		Id:         uint32(survey.Id),
		SurveyText: survey.SurveyText,
		MaxRating:  uint32(survey.MaxRating),
	}
	return &gen.GetSurveyByIdResponce{Survey: surveyResponce}, nil
}

func (s *GrpcSurveyHandler) CreateSurvey(ctx context.Context, req *gen.CreateSurveyRequest) (*gen.CreateSurveyResponce, error) {
	survey := models.SurveyResponse{
		SurveyId: uint(req.ServeyResponce.SurveyId),
		UserId:   uint(req.ServeyResponce.UserId),
		Rating:   int(req.ServeyResponce.Rating),
	}
	err := s.uc.CreateSurveyResponse(ctx, survey)
	if err != nil {
		return nil, err
	}
	return &gen.CreateSurveyResponce{Success: true}, nil
}

func (s *GrpcSurveyHandler) GetSurveyStatsBySurveyId(ctx context.Context, req *gen.GetSurveyStatsBySurveyIdRequest) (*gen.GetSurveyStatsBySurveyIdResponce, error) {
	res, err := s.uc.GetSurveyStatsBySurveyId(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}
	ratingsCount := make(map[int32]int32)
	for k, v := range res.RatingsCount {
		ratingsCount[int32(k)] = int32(v)
	}
	statsResponce := &gen.SurveyStatsBySurvey{
		ServeyId:     uint32(res.SurveyId),
		ServeyText:   res.SurveyText,
		AvgRating:    float32(res.AvgRating),
		RatingsCount: ratingsCount,
	}
	return &gen.GetSurveyStatsBySurveyIdResponce{SurveyStatsBySurvey: statsResponce}, nil
}

func (s *GrpcSurveyHandler) GetSurveyStatsByUserId(ctx context.Context, req *gen.GetSurveyStatsByUserIdRequest) (*gen.GetSurveyStatsByUserIdResponce, error) {
	res, err := s.uc.GetSurveyStatsByUserId(ctx, uint(req.UserId))
	if err != nil {
		return nil, err
	}
	serveyStatsResponce := make([]*gen.UserSurveyStats, len(res))
	for k, v := range res {
		serveyStatsResponce[k] = &gen.UserSurveyStats{
			ServeyId:   uint32(v.SurveyId),
			ServeyText: v.SurveyText,
			Answered:   v.Answered,
		}
	}
	return &gen.GetSurveyStatsByUserIdResponce{UserServeyStats: serveyStatsResponce}, nil
}
