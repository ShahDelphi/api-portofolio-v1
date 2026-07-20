package database

import (
	"log"
	"os"

	"portfolio-backend/internal/modules/admin"
	"portfolio-backend/internal/modules/project"
	"portfolio-backend/internal/modules/profile"
	"portfolio-backend/internal/modules/skill"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedData seeds initial data into the database
func SeedData(db *gorm.DB) error {
	// 1. Seed Admin User
	var adminCount int64
	db.Model(&admin.Admin{}).Count(&adminCount)

	if adminCount == 0 {
		adminUser := os.Getenv("ADMIN_USERNAME")
		adminPass := os.Getenv("ADMIN_PASSWORD")
		adminEmail := os.Getenv("ADMIN_EMAIL")

		if adminUser == "" { adminUser = "admin" }
		if adminPass == "" { adminPass = "admin123" }
		if adminEmail == "" { adminEmail = "admin@example.com" }

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPass), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user := admin.Admin{
			ID:       uuid.New(),
			Name:     "Portfolio Admin",
			Username: adminUser,
			Email:    adminEmail,
			Password: string(hashedPassword),
		}

		if err := db.Create(&user).Error; err != nil {
			return err
		}
		log.Printf("✓ Admin user seeded: %s / %s", adminUser, adminPass)
	} else {
		log.Println("✓ Admin user already exists, skipping seed")
	}

	// 2. Seed Projects
	var projectCount int64
	db.Model(&project.Project{}).Count(&projectCount)

	if projectCount == 0 {
		projects := []project.Project{
			{
				ID:          uuid.New(),
				Title:       "YOLOv8 Segmentation for Palm Oil Disease Detection",
				Description: "Applied deep learning segmentation (YOLOv8) to identify and classify diseases in oil palm leaves from high-resolution drone imagery. Improved leaf disease detection accuracy by 15% and enabled automated estate health monitoring.",
				Thumbnail:   "/uploads/projects/sample-yolo.webp",
				GithubURL:   "https://github.com/username/yolov8-palm-oil",
				DemoURL:     "",
				TechStack:   "Python,YOLOv8,PyTorch,OpenCV,Flask",
				Order:       1,
			},
			{
				ID:          uuid.New(),
				Title:       "AEON Mall Forecasting & Ordering System",
				Description: "Developed an inventory forecasting engine integrated with a smart ordering workflow for AEON Mall fresh food sections. Utilized seasonal ARIMA models to reduce food waste by 22% while maintaining optimal stock availability.",
				Thumbnail:   "/uploads/projects/sample-forecasting.webp",
				GithubURL:   "https://github.com/username/aeon-forecasting",
				DemoURL:     "",
				TechStack:   "Golang,React,PostgreSQL,Python,ARIMA",
				Order:       2,
			},
			{
				ID:          uuid.New(),
				Title:       "Oil Market Forecasting for Supply Chain Decision Making",
				Description: "Built a predictive analytics dashboard utilizing historical prices, geopolitical events, and freight capacity data to model global crude oil pricing. Provided actionable supply chain recommendations to minimize procurement costs.",
				Thumbnail:   "/uploads/projects/sample-supplychain.webp",
				GithubURL:   "https://github.com/username/oil-supply-chain",
				DemoURL:     "",
				TechStack:   "Python,Pandas,XGBoost,FastAPI,React,Docker",
				Order:       3,
			},
			{
				ID:          uuid.New(),
				Title:       "Fullstack Printing Web App",
				Description: "A modern web portal for handling online printing requests. Users can upload files, choose paper options, view instant quotes, and pay online. Admins can manage incoming jobs, update printing status, and handle billing.",
				Thumbnail:   "/uploads/projects/sample-printing.webp",
				GithubURL:   "https://github.com/username/printing-web-app",
				DemoURL:     "https://printing-demo.example.com",
				TechStack:   "React,Tailwind CSS,Node.js,Express,PostgreSQL",
				Order:       4,
			},
			{
				ID:          uuid.New(),
				Title:       "Backend & API Projects Collection",
				Description: "A compilation of clean-architecture microservices and RESTful API endpoints built using Golang (Fiber/Gin) and Node.js. Features rate limiting, structured logging, JWT authentication, and comprehensive unit testing.",
				Thumbnail:   "/uploads/projects/sample-backend.webp",
				GithubURL:   "https://github.com/username/backend-api-collection",
				DemoURL:     "",
				TechStack:   "Golang,Fiber,Redis,PostgreSQL,JWT,Docker",
				Order:       5,
			},
		}

		for _, p := range projects {
			if err := db.Create(&p).Error; err != nil {
				return err
			}
		}
		log.Println("✓ Projects seeded successfully")
	} else {
		log.Println("✓ Projects already exist, skipping seed")
	}

	// 3. Seed Skills
	var skillCount int64
	db.Model(&skill.Skill{}).Count(&skillCount)

	if skillCount == 0 {
		skills := []skill.Skill{
			// Frontend
			{ID: uuid.New(), Name: "React", Category: "Frontend", Order: 1},
			{ID: uuid.New(), Name: "Tailwind CSS", Category: "Frontend", Order: 2},
			{ID: uuid.New(), Name: "Vite", Category: "Frontend", Order: 3},
			{ID: uuid.New(), Name: "HTML5/CSS3", Category: "Frontend", Order: 4},

			// Backend
			{ID: uuid.New(), Name: "Golang (Fiber/Gin)", Category: "Backend", Order: 1},
			{ID: uuid.New(), Name: "Node.js (Express)", Category: "Backend", Order: 2},
			{ID: uuid.New(), Name: "RESTful API", Category: "Backend", Order: 3},
			{ID: uuid.New(), Name: "GORM / Hibernate", Category: "Backend", Order: 4},

			// Data Science & AI
			{ID: uuid.New(), Name: "Python (Pandas/NumPy)", Category: "Data Science & AI", Order: 1},
			{ID: uuid.New(), Name: "Machine Learning (XGBoost/Scikit-Learn)", Category: "Data Science & AI", Order: 2},
			{ID: uuid.New(), Name: "Deep Learning (YOLOv8/PyTorch)", Category: "Data Science & AI", Order: 3},
			{ID: uuid.New(), Name: "Forecasting (ARIMA/SARIMAX)", Category: "Data Science & AI", Order: 4},

			// Databases & Tools
			{ID: uuid.New(), Name: "PostgreSQL", Category: "Databases & Tools", Order: 1},
			{ID: uuid.New(), Name: "MySQL", Category: "Databases & Tools", Order: 2},
			{ID: uuid.New(), Name: "Git & GitHub", Category: "Databases & Tools", Order: 3},
			{ID: uuid.New(), Name: "Docker", Category: "Databases & Tools", Order: 4},
		}

		for _, s := range skills {
			if err := db.Create(&s).Error; err != nil {
				return err
			}
		}
		log.Println("✓ Skills seeded successfully")
	} else {
		log.Println("✓ Skills already exist, skipping seed")
	}

	// 4. Seed Profile
	var profileCount int64
	db.Model(&profile.Profile{}).Count(&profileCount)

	if profileCount == 0 {
		prof := profile.Profile{
			ID:        uuid.New(),
			Name:      "Brittany Chiang",
			Title:     "Brittany Chiang.",
			Subtitle:  "I build things for the web.",
			Intro:     "I’m a software engineer specializing in building (and occasionally designing) exceptional digital experiences. Currently, I’m focused on building accessible, human-centered products at Upstatement.",
			Bio:       "<p>Hello! My name is Brittany and I enjoy creating things that live on the internet. My interest in web development started back in 2012 when I decided to try editing custom Tumblr themes — turns out hacking together a custom reblog button taught me a lot about HTML & CSS!</p>\n\n<p>Fast-forward to today, and I’ve had the privilege of working at <a href=\"https://us.mullenlowe.com/\">an advertising agency</a>, <a href=\"https://starry.com/\">a start-up</a>, <a href=\"https://www.apple.com/\">a huge corporation</a>, and <a href=\"https://scout.camd.northeastern.edu/\">a student-led design studio</a>. My main focus these days is building accessible, inclusive products and digital experiences at <a href=\"https://upstatement.com/\">Upstatement</a> for a variety of clients.</p>\n\n<p>I also recently <a href=\"https://www.newline.co/courses/build-a-spotify-connected-app\">launched a course</a> that covers everything you need to build a web app with the Spotify API using Node & React.</p>\n\n<p>Here are a few technologies I’ve been working with recently:</p>",
			Avatar:    "",
			ResumeURL: "",
			Github:    "https://github.com/bchiang7",
			Instagram: "https://www.instagram.com/bchiang7",
			Linkedin:  "https://www.linkedin.com/in/bchiang7",
			Email:     "brittany.chiang@gmail.com",
		}

		if err := db.Create(&prof).Error; err != nil {
			return err
		}
		log.Println("✓ Profile seeded successfully")
	} else {
		log.Println("✓ Profile already exists, skipping seed")
	}

	return nil
}
