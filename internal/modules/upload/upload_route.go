package upload

import (
	"portfolio-backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, handler *Handler) {
	upload := router.Group("/upload")

	// Protected routes (auth required)
	upload.Use(middleware.AuthMiddleware())
	upload.Post("/:type", handler.UploadImage)
}
