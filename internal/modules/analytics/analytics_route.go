package analytics

import (
	"portfolio-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	analytics := router.Group("/analytics")

	// Protected routes (auth required)
	analytics.Use(middleware.AuthMiddleware())
	analytics.Get("/summary", handler.GetSummary)
}
