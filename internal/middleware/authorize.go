package middleware

import (
	"slices"
	"strings"
	"task-tracker/internal/context"
	"task-tracker/internal/util"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Authorize(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			return &util.AuthError{Message: "Missing token"}
		}

		tokenString := strings.TrimPrefix(auth, "Bearer ")

		token, err := util.ParseJWT(tokenString)
		if err != nil || !token.Valid {
			return &util.AuthError{Message: "Invalid token"}

		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return &util.AuthError{Message: "Failed to parse token"}
		}

		if len(allowedRoles) > 0 && !slices.Contains(allowedRoles, claims["role"].(string)) {
			return &util.AuthError{Message: "You do not have permission to perform this action"}
		}

		context.SetUserContext(c, context.UserContext{
			UserID: uint(claims["user_id"].(float64)),
			Role:   claims["role"].(string),
		})

		return c.Next()
	}
}
