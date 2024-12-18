package grpc

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/user"
	"2024_2_ThereWillBeName/internal/pkg/user/delivery/grpc/gen"
	"context"
	"log"
	"log/slog"
)

type GrpcUserHandler struct {
	gen.UserServiceServer
	usecase user.UserUsecase
	logger  *slog.Logger
}

func NewGrpcUserHandler(usecase user.UserUsecase, logger *slog.Logger) *GrpcUserHandler {
	return &GrpcUserHandler{usecase: usecase, logger: logger}
}

func (h *GrpcUserHandler) SignUp(ctx context.Context, in *gen.SignUpRequest) (*gen.SignUpResponse, error) {
	user := models.User{
		Login:    in.Login,
		Email:    in.Email,
		Password: in.Password,
	}

	userID, err := h.usecase.SignUp(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return &gen.SignUpResponse{
		Id: uint32(userID),
	}, nil
}

func (h *GrpcUserHandler) Login(ctx context.Context, in *gen.LoginRequest) (*gen.LoginResponse, error) {
	user, err := h.usecase.Login(context.Background(), in.Email, in.Password)

	if err != nil {
		return nil, err
	}

	return &gen.LoginResponse{
		Id:         uint32(user.ID),
		Login:      user.Login,
		Email:      user.Email,
		AvatarPath: user.AvatarPath,
	}, nil
}

func (h *GrpcUserHandler) UploadAvatar(ctx context.Context, in *gen.UploadAvatarRequest) (*gen.UploadAvatarResponse, error) {
	avatarPath, err := h.usecase.UploadAvatar(context.Background(), uint(in.Id), in.AvatarData, in.AvatarFileName)

	if err != nil {
		return nil, err
	}

	return &gen.UploadAvatarResponse{
		AvatarPath: avatarPath,
	}, nil
}

func (h *GrpcUserHandler) GetProfile(ctx context.Context, in *gen.GetProfileRequest) (*gen.GetProfileResponse, error) {
	profile, err := h.usecase.GetProfile(ctx, uint(in.Id), uint(in.RequesterId))
	if err != nil {
		return nil, err
	}

	return &gen.GetProfileResponse{
		Login:      profile.Login,
		Email:      profile.Email,
		AvatarPath: profile.AvatarPath,
	}, nil
}

func (h *GrpcUserHandler) UpdatePassword(ctx context.Context, in *gen.UpdatePasswordRequest) (*gen.EmptyResponse, error) {
	user := models.User{
		ID:       uint(in.Id),
		Login:    in.Login,
		Email:    in.Email,
		Password: in.OldPassword,
	}

	err := h.usecase.UpdatePassword(ctx, user, in.NewPassword)
	if err != nil {
		return nil, err
	}

	return &gen.EmptyResponse{}, nil
}

func (h *GrpcUserHandler) UpdateProfile(ctx context.Context, in *gen.UpdateProfileRequest) (*gen.EmptyResponse, error) {

	err := h.usecase.UpdateProfile(ctx, uint(in.UserId), in.Username, in.Email)
	if err != nil {
		return nil, err
	}

	return &gen.EmptyResponse{}, nil
}

func (h *GrpcUserHandler) GetAchievements(ctx context.Context, in *gen.GetAchievementsRequest) (*gen.GetAchievementsResponse, error) {
	achievements, err := h.usecase.GetAchievements(ctx, uint(in.Id))
	if err != nil {
		return nil, err
	}
	log.Println(achievements)

	achievementsResponse := make([]*gen.Achievement, len(achievements))
	for i, achievement := range achievements {
		achievementsResponse[i] = &gen.Achievement{
			Id:       uint32(achievement.ID),
			Name:     achievement.Name,
			IconPath: achievement.IconPath,
		}
	}
	return &gen.GetAchievementsResponse{Achievements: achievementsResponse}, nil
}
