package analytics

import (
	"portfolio-backend/internal/modules/certificate"
	"portfolio-backend/internal/modules/experience"
	"portfolio-backend/internal/modules/project"
	"portfolio-backend/internal/modules/skill"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetSummary(c *fiber.Ctx) error {
	var projectsCount int64
	var experiencesCount int64
	var certificatesCount int64
	var skillsCount int64

	if err := h.db.Model(&project.Project{}).Count(&projectsCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to count projects",
		})
	}

	if err := h.db.Model(&experience.Experience{}).Count(&experiencesCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to count experiences",
		})
	}

	if err := h.db.Model(&certificate.Certificate{}).Count(&certificatesCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to count certificates",
		})
	}

	if err := h.db.Model(&skill.Skill{}).Count(&skillsCount).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to count skills",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Analytics summary retrieved successfully",
		"data": fiber.Map{
			"projects":     projectsCount,
			"experiences":  experiencesCount,
			"certificates": certificatesCount,
			"skills":       skillsCount,
		},
	})
}
