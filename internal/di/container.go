package di

import (
	"task-tracker/internal/controller"
	"task-tracker/internal/repository"
	"task-tracker/internal/service"

	"gorm.io/gorm"
)

type Container struct {
	UserController *controller.UserController
	AuthController *controller.AuthController
	TaskController *controller.TaskController
	LogController  *controller.LogController
}

func InitContainer(db *gorm.DB) *Container {
	//User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	//Auth
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)

	//Task
	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskController := controller.NewTaskController(taskService)

	//Log
	logRepo := repository.NewLogRepository(db)
	logService := service.NewLogService(logRepo)
	logController := controller.NewLogController(logService, userService, taskService)

	return &Container{
		UserController: userController,
		AuthController: authController,
		TaskController: taskController,
		LogController:  logController,
	}
}
