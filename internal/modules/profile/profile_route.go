package profile

import (
	"portfolio-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router, handler *Handler) {
	// Root profile endpoint
	profile := router.Group("/profile")

	// Public routes
	profile.Get("/", handler.Get)

	// Protected routes (Admin only)
	profile.Put("/", middleware.AuthMiddleware(), handler.Update)
}
