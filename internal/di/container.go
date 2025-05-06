package di

import (
	"task-tracker/internal/controller"
	"task-tracker/internal/repository"
	"task-tracker/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	UserController *controller.UserController
}

func InitContainer(db *gorm.DB) *Container {
	//User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	return &Container{
		UserController: userController,
	}
}
