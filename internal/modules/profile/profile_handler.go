package profile

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repo Repository
}

func NewHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

// Get handles fetching the profile
func (h *Handler) Get(c *fiber.Ctx) error {
	prof, err := h.repo.Get()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Profile details not found",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Profile retrieved successfully",
		"data":    prof,
	})
}

// Update handles updating the profile
func (h *Handler) Update(c *fiber.Ctx) error {
	prof, err := h.repo.Get()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Profile not found to update",
		})
	}

	var updateData Profile
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validation
	if updateData.Name == "" || updateData.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Name and title are required fields",
		})
	}

	// Update fields
	prof.Name = updateData.Name
	prof.Title = updateData.Title
	prof.Subtitle = updateData.Subtitle
	prof.Intro = updateData.Intro
	prof.Bio = updateData.Bio
	prof.Avatar = updateData.Avatar
	prof.ResumeURL = updateData.ResumeURL
	prof.Github = updateData.Github
	prof.Instagram = updateData.Instagram
	prof.Linkedin = updateData.Linkedin
	prof.Email = updateData.Email

	if err := h.repo.Update(prof); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to update profile details",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Profile updated successfully",
		"data":    prof,
	})
}
