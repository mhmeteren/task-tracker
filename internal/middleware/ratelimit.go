package middleware

import (
	"fmt"
	"task-tracker/config"
	"task-tracker/internal/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimitByIP() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        config.Cfg.RateLimit.Max,
		Expiration: config.Cfg.RateLimit.Expiration,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return &util.RateLimitError{Message: "Too many requests. Please try again later."}
		},
	})
}

func RateLimitByTask() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        config.Cfg.TaskRateLimit.Max,
		Expiration: config.Cfg.TaskRateLimit.Expiration,
		KeyGenerator: func(c *fiber.Ctx) string {
			taskKey := c.Params("taskkey")
			taskSecret := c.Params("taskSecret")
			return fmt.Sprintf("%s:%s", taskKey, taskSecret)
		},
		LimitReached: func(c *fiber.Ctx) error {
			return &util.RateLimitError{Message: "Rate limit exceeded for this task. Try again later."}
		},
	})
}
