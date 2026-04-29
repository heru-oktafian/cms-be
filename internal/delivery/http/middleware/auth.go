package middleware

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtAuth "github.com/heru-oktafian/cms-be/internal/infrastructure/auth"
)

func AdminJWT(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := strings.TrimSpace(c.Get("Authorization"))
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "missing authorization header",
				"data":    nil,
			})
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid authorization header",
				"data":    nil,
			})
		}

		claims, err := jwtAuth.ParseAdminJWT(parts[1], secret)
		if err != nil {
			log.Printf("[jwt-debug] parse error: %v", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "invalid token",
				"data":    nil,
			})
		}

		c.Locals("auth_user", fiber.Map{
			"sub":   claims.Subject,
			"email": claims.Email,
			"role":  claims.Role,
			"exp":   claims.ExpiresAt.Unix(),
		})
		return c.Next()
	}
}
