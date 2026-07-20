package upload

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

var allowedTypes = map[string]bool{
	"image/jpeg":      true,
	"image/jpg":       true,
	"image/png":       true,
	"image/webp":      true,
	"image/gif":       true,
	"application/pdf": true,
}

var uploadPaths = map[string]string{
	"project":     "./public/uploads/projects",
	"certificate": "./public/uploads/certificates",
	"profile":     "./public/uploads/profile",
}

// UploadImage handles image upload
func (h *Handler) UploadImage(c *fiber.Ctx) error {
	// Get upload type from params
	uploadType := c.Params("type")

	// Validate upload type
	uploadPath, exists := uploadPaths[uploadType]
	if !exists {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid upload type. Allowed types: project, certificate, profile",
		})
	}

	// Ensure upload directory exists
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create upload directory",
			"error":   err.Error(),
		})
	}

	// Get file from form
	file, err := c.FormFile("file")
	if err != nil {
		file, err = c.FormFile("image") // fallback
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "No file uploaded",
				"error":   err.Error(),
			})
		}
	}

	// Validate file size (max 2MB)
	maxSize := int64(2 * 1024 * 1024) // 2MB
	if file.Size > maxSize {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "File size too large. Maximum 2MB",
		})
	}

	// Validate file type
	contentType := file.Header.Get("Content-Type")
	if contentType == "application/pdf" && uploadType != "profile" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "PDF files are only allowed for profile upload",
		})
	}
	if !allowedTypes[contentType] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid file type. Allowed types: jpg, jpeg, png, webp, gif, pdf",
		})
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s-%d%s", uuid.New().String(), time.Now().Unix(), ext)

	// Save file
	filePath := filepath.Join(uploadPath, filename)
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to save file",
			"error":   err.Error(),
		})
	}

	// Generate URL
	// Remove ./public from path and replace backslashes with forward slashes
	urlPath := strings.ReplaceAll(uploadPath, "./public", "")
	urlPath = strings.ReplaceAll(urlPath, "\\", "/")
	fileURL := fmt.Sprintf("%s/%s", urlPath, filename)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "File uploaded successfully",
		"data": fiber.Map{
			"filename": filename,
			"url":      fileURL,
			"size":     file.Size,
			"type":     contentType,
		},
	})
}
