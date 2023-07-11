package handler

import (
	"athun/pkg/domain"
	"athun/pkg/pb"
	interfaces "athun/pkg/usecase/interface"
	"context"
	"net/http"
)

type UserHandler struct {
	UseCase interfaces.UserUseCase
	pb.AuthServiceServer
}

func NewUserHandler(useCase interfaces.UserUseCase) *UserHandler {
	return &UserHandler{
		UseCase: useCase,
	}
}

func (h *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := domain.User{
		Username: req.Username,
		Email:    req.Email,
	}
	err := h.UseCase.Register(user)
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error",
		}, err
	}
	return &pb.RegisterResponse{
		Status: http.StatusOK,
		Error:  "nil",
	}, nil
}
