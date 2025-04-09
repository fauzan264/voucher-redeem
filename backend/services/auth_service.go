package services

import (
	"time"

	"github.com/fauzan264/voucher-redeem/backend/constants"
	"github.com/fauzan264/voucher-redeem/backend/domain/dto/request"
	"github.com/fauzan264/voucher-redeem/backend/domain/dto/response"
	"github.com/fauzan264/voucher-redeem/backend/domain/models"
	"github.com/fauzan264/voucher-redeem/backend/repositories"
	"github.com/fauzan264/voucher-redeem/backend/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(request request.RegisterRequest) (response.UserResponse, error)
	LoginUser(request request.LoginRequest) (response.UserResponse, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) *authService {
	return &authService{userRepository}
}

func (s *authService) RegisterUser(request request.RegisterRequest) (response.UserResponse, error) {
	userData := models.User{
		ID: uuid.New(),
		Name: request.Name,
		Email: request.Email,
		IsAdmin: request.IsAdmin,
		CreatedAt: time.Now(),
	}
	
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		return response.UserResponse{}, err
	}
	
	userData.PasswordHash = string(passwordHash)
	user, err := s.userRepository.CreateUser(userData)
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

func (s *authService) LoginUser(request request.LoginRequest) (response.UserResponse, error) {
	email := request.Email
	password := request.Password

	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return response.UserResponse{}, constants.ErrWrongUserOrPassword
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return response.UserResponse{}, constants.ErrWrongUserOrPassword
	}
	
	token, err := utils.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		return response.UserResponse{}, err
	}

	loginResponse := response.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		TotalPoints: user.TotalPoints,
		IsAdmin: user.IsAdmin,
		Token: token,
	}

	return loginResponse, err
}