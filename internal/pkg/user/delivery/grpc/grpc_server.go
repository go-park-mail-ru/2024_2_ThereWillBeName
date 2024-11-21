package grpc

import (
	"2024_2_ThereWillBeName/internal/pkg/user"
	"2024_2_ThereWillBeName/internal/pkg/user/delivery/grpc/gen"
)

type GrpcUserHandler struct {
	gen.UnimplementedUserServiceServer
	uc user.UserUsecase
}
