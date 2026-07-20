package experience

import (
	"portfolio-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	exp := router.Group("/experiences")

	// Public routes
	exp.Get("/", handler.GetAll)
	exp.Get("/:id", handler.GetByID)

	// Protected routes
	exp.Use(middleware.AuthMiddleware())
	exp.Post("/", handler.Create)
	exp.Put("/:id", handler.Update)
	exp.Delete("/:id", handler.Delete)
}
