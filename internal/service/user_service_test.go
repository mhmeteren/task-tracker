package service_test

import (
	"errors"
	"testing"

	"task-tracker/internal/model"
	"task-tracker/internal/repository/mock"
	"task-tracker/internal/service"

	"github.com/golang/mock/gomock"
)

func TestUserService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		Create(gomock.Any()).
		DoAndReturn(func(user *model.User) error {
			user.ID = 1
			return nil
		})

	userService := service.NewUserService(mockRepo)

	user := &model.User{
		Name:  "Test User",
		Email: "test@example.com",
	}

	err := userService.Create(user, "123456")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
func TestUserService_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUser := model.User{
		Name:  "Test User",
		Email: "test@example.com",
	}
	mockUser.ID = 1

	mockRepo := mock.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		FindByID(uint(1)).
		Return(&mockUser, nil).
		Times(1)

	mockRepo.EXPECT().
		FindByID(uint(2)).
		Return(nil, errors.New("user not found")).
		Times(1)

	userService := service.NewUserService(mockRepo)

	user, err := userService.GetByID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user == nil || user.ID != 1 {
		t.Errorf("expected user with ID 1, got %+v", user)
	}

	user, err = userService.GetByID(2)
	if err == nil {
		t.Errorf("expected error for non-existent user, got nil")
	}
	if user != nil {
		t.Errorf("expected nil user for non-existent user, got %+v", user)
	}
}
