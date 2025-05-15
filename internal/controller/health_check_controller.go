package controller

import (
	"log"
	"task-tracker/internal/database"
	"task-tracker/internal/dto"

	"github.com/gofiber/fiber/v2"
)

type HealthCheckController struct{}

func NewHealthCheckController() *HealthCheckController {
	return &HealthCheckController{}
}

// HealthCheck godoc
// @Summary Check system health
// @Description Returns application health status and database connectivity check
// @Tags System
// @Produce json
// @Success 200 {object} dto.HealthCheckResponse
// @Failure 503 {object} dto.HealthCheckResponse
// @Router /health [get]
func (h *HealthCheckController) Health(c *fiber.Ctx) error {

	response := dto.HealthCheckResponse{
		Status: "Healthy",
		Entries: map[string]dto.HealthEntry{
			"SQL Server Check": {
				Status: "Healthy",
				Tags:   []string{"sql", "sql-server", "pgsql"},
			},
		},
	}

	sqlDB, err := database.DB.DB()

	if err != nil || sqlDB.Ping() != nil {
		log.Println("Health check failed:", err)
		response.Status = "Unhealthy"

		response.Entries["SQL Server Check"] = dto.HealthEntry{
			Status: "Unhealthy",
			Tags:   response.Entries["SQL Server Check"].Tags,
		}

	}

	var status int = fiber.StatusOK
	if response.Status == "Unhealthy" {
		status = fiber.StatusServiceUnavailable
	}

	return c.Status(status).JSON(response)
}
