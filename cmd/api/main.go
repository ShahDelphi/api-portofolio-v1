package main

import (
	"fmt"
	"log"
	"os"

	"portfolio-backend/configs"
	"portfolio-backend/internal/database"
	"portfolio-backend/internal/modules/admin"
	"portfolio-backend/internal/modules/analytics"
	"portfolio-backend/internal/modules/auth"
	"portfolio-backend/internal/modules/certificate"
	"portfolio-backend/internal/modules/experience"
	"portfolio-backend/internal/modules/profile"
	"portfolio-backend/internal/modules/project"
	"portfolio-backend/internal/modules/skill"
	"portfolio-backend/internal/modules/upload"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// 1. Load configuration
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// 2. Connect to database
	db, err := database.Connect(config.GetDSN())
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 3. Auto migrate schemas
	if err := database.AutoMigrate(db,
		&admin.Admin{},
		&project.Project{},
		&experience.Experience{},
		&certificate.Certificate{},
		&skill.Skill{},
		&profile.Profile{},
	); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// 4. Check for --seed flag to manually seed
	if len(os.Args) > 1 && os.Args[1] == "--seed" {
		if err := database.SeedData(db); err != nil {
			log.Fatal("Failed to seed database:", err)
		}
		log.Println("✓ Database seeded successfully")
		return
	}

	// Always run database seed for default admin and sample data if database is empty
	if err := database.SeedData(db); err != nil {
		log.Println("Warning: Failed to seed database:", err)
	}

	// 5. Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: config.AppName,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})

	// 6. Global Middlewares
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-Requested-With",
	}))

	// 7. Serve static files (uploaded images)
	app.Static("/uploads", "./public/uploads")

	// 8. Health check & Root endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Portfolio API is running",
			"version": "1.0.0",
		})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "Server is healthy",
		})
	})

	// 9. API v1 routes group
	api := app.Group("/api/v1")

	// Auth routes (public login)
	authRepo := auth.NewRepository(db)
	authHandler := auth.NewHandler(authRepo)
	auth.RegisterRoutes(api, authHandler)

	// Projects routes (public GET, protected POST/PUT/DELETE)
	projectRepo := project.NewRepository(db)
	projectHandler := project.NewHandler(projectRepo)
	project.RegisterRoutes(api, projectHandler)

	// Experiences routes
	experienceRepo := experience.NewRepository(db)
	experienceHandler := experience.NewHandler(experienceRepo)
	experience.RegisterRoutes(api, experienceHandler)

	// Certificates routes
	certificateRepo := certificate.NewRepository(db)
	certificateHandler := certificate.NewHandler(certificateRepo)
	certificate.RegisterRoutes(api, certificateHandler)

	// Skills routes
	skillRepo := skill.NewRepository(db)
	skillHandler := skill.NewHandler(skillRepo)
	skill.RegisterRoutes(api, skillHandler)

	// Profile routes
	profileRepo := profile.NewRepository(db)
	profileHandler := profile.NewHandler(profileRepo)
	profile.SetupRoutes(api, profileHandler)

	// Upload routes (protected)
	uploadHandler := upload.NewHandler()
	upload.RegisterRoutes(api, uploadHandler)

	// Analytics routes (protected)
	analyticsHandler := analytics.NewHandler(db)
	analytics.RegisterRoutes(api, analyticsHandler)

	// 10. Start server
	port := config.AppPort
	log.Printf("🚀 Server starting on http://%s:%s\n", config.AppHost, port)
	log.Printf("📚 Environment: %s\n", config.AppEnv)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
