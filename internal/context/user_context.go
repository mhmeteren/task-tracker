package context

import "github.com/gofiber/fiber/v2"

const UserContextKey = "user_ctx"

type UserContext struct {
	UserID uint
	Role   string
}

func SetUserContext(c *fiber.Ctx, ctx UserContext) {
	c.Locals(UserContextKey, ctx)
}

//Get user information from locals
func GetUserContext(c *fiber.Ctx) *UserContext {
	val := c.Locals(UserContextKey)
	if userCtx, ok := val.(UserContext); ok {
		return &userCtx
	}
	return nil
}
