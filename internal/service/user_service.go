package service

import (
	"task-tracker/internal/model"
	"task-tracker/internal/repository"
	"task-tracker/internal/util"
)

type UserService interface {
	GetAll() ([]model.User, error)
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetByEmailWithRole(email string) (*model.User, error)
	GetProfile(userID uint) (*model.User, error)
	Create(user *model.User, plainPassword string) error
	Update(user *model.User) error
	Delete(id uint) error
	GetUserByIdCheckAndExists(userID uint) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetAll() ([]model.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetByID(id uint) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetByEmail(email string) (*model.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *userService) GetByEmailWithRole(email string) (*model.User, error) {
	return s.repo.FindByEmailWithRole(email)
}

func (s *userService) GetProfile(userID uint) (*model.User, error) {
	user, err := s.GetUserByIdCheckAndExists(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Create(user *model.User, plainPassword string) error {
	hashedPassword, err := util.HashPassword(plainPassword)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	user.RoleID = 2 // user

	if err := s.repo.Create(user); err != nil {
		return err
	}
	return nil
}

func (s *userService) Update(user *model.User) error {
	return s.repo.Update(user)
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *userService) GetUserByIdCheckAndExists(userID uint) (*model.User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, &util.NotFoundError{Message: "user not found"}
	}
	return user, nil
}
