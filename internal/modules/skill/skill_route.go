package skill

import (
	"portfolio-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	sk := router.Group("/skills")

	// Public routes
	sk.Get("/", handler.GetAll)
	sk.Get("/:id", handler.GetByID)

	// Protected routes
	sk.Use(middleware.AuthMiddleware())
	sk.Post("/", handler.Create)
	sk.Put("/:id", handler.Update)
	sk.Delete("/:id", handler.Delete)
}
