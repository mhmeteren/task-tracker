package service

import (
	"errors"
	"task-tracker/internal/dto"
	"task-tracker/internal/model"
	"task-tracker/internal/repository"
	"task-tracker/internal/util"
	"time"
)

type AuthService interface {
	Login(dto *dto.LoginRequest) (result *dto.AuthResponse, err error)
	RefreshToken(dto *dto.RefreshTokenRequest) (*dto.AuthResponse, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (service *authService) Login(dto *dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := service.userRepo.FindByEmailWithRole(dto.Email)
	if err != nil {
		return nil, errors.New("email or password is wrong")
	}

	if !util.CheckPasswordHash(dto.Password, user.Password) {
		return nil, errors.New("email or password is wrong")
	}

	result, err := service.createAndSaveTokens(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *authService) RefreshToken(dto *dto.RefreshTokenRequest) (*dto.AuthResponse, error) {

	user, err := service.userRepo.FindByTokenWithRole(util.HashToken(dto.RefreshToken))
	if err != nil || user.RefreshTokenExpiresAt.Before(time.Now()) {
		return nil, errors.New("invalid or expired refresh token")
	}

	result, err := service.createAndSaveTokens(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *authService) createAndSaveTokens(user *model.User) (*dto.AuthResponse, error) {
	accessToken, err := util.GenerateJWT(user.ID, user.Role.Name)
	if err != nil {
		return nil, err
	}

	refreshToken, expiry, err := util.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	hashToken := util.HashToken(refreshToken)
	user.RefreshToken = &hashToken
	user.RefreshTokenExpiresAt = expiry
	service.userRepo.Update(user)

	return &dto.AuthResponse{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
