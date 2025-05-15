package controller_test

import (
	"io"
	"net/http/httptest"
	"strings"
	"task-tracker/internal/context"
	"task-tracker/internal/controller"
	"task-tracker/internal/model"
	"task-tracker/internal/service/mock"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserController_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mock.NewMockUserService(ctrl)

	userController := controller.NewUserController(mockUserService)

	app := fiber.New()

	app.Post("/users", func(c *fiber.Ctx) error {
		return userController.CreateUser(c)
	})

	// Test case: email already exists
	mockUserService.EXPECT().
		GetByEmail("test@example.com").
		Return(&model.User{Email: "test@example.com"}, nil).
		Times(1)

	reqBody := `{"name":"Test User", "email":"test@example.com", "password":"123456"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestUserController_GetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mock.NewMockUserService(ctrl)
	userController := controller.NewUserController(mockUserService)

	app := fiber.New()

	app.Get("/profile", func(c *fiber.Ctx) error {
		c.Locals(context.UserContextKey, context.UserContext{
			UserID: 1,
			Role:   "user",
		})
		return userController.GetProfile(c)
	})

	expectedUser := &model.User{
		Name:  "John Doe",
		Email: "john@example.com",
	}
	expectedUser.ID = 1

	mockUserService.
		EXPECT().
		GetProfile(uint(1)).
		Return(expectedUser, nil).
		Times(1)

	req := httptest.NewRequest("GET", "/profile", nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	assert.Contains(t, string(body), `"email":"john@example.com"`)
}
