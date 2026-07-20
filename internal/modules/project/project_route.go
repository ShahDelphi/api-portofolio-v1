package project

import (
	"portfolio-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	proj := router.Group("/projects")

	// Public routes
	proj.Get("/", handler.GetAll)
	proj.Get("/:id", handler.GetByID)

	// Protected routes
	proj.Use(middleware.AuthMiddleware())
	proj.Post("/", handler.Create)
	proj.Put("/:id", handler.Update)
	proj.Delete("/:id", handler.Delete)
}
