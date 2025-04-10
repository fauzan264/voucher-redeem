package services

import (
	"github.com/fauzan264/voucher-redeem/domain/dto/request"
	"github.com/fauzan264/voucher-redeem/domain/dto/response"
	"github.com/fauzan264/voucher-redeem/repositories"
)

type userService struct {
	repository repositories.UserRepository
}

type UserService interface {
	GetUserByID(request request.GetUser) (response.UserResponse, error)
}
func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) GetUserByID(request request.GetUser) (response.UserResponse, error) {
	user, err := s.repository.GetUserByID(request.ID)
	if err != nil {
		return response.UserResponse{}, err
	}

	userResponse := response.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		TotalPoints: user.TotalPoints,
		IsAdmin: user.IsAdmin,
	}

	return userResponse, nil
}