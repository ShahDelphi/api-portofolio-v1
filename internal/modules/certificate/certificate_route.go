package certificate

import (
	"portfolio-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	cert := router.Group("/certificates")

	// Public routes
	cert.Get("/", handler.GetAll)
	cert.Get("/:id", handler.GetByID)

	// Protected routes
	cert.Use(middleware.AuthMiddleware())
	cert.Post("/", handler.Create)
	cert.Put("/:id", handler.Update)
	cert.Delete("/:id", handler.Delete)
}
