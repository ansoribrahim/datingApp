package main

import (
	"fmt"
	"log"

	"datingApp/config"
	"datingApp/models"
	"datingApp/repositories"
	"datingApp/routes"
	"datingApp/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	log.Println("Resetting and migrating the database...")

	// Drop all tables
	err := db.Migrator().DropTable(
		&models.User{},
		&models.Profile{},
		&models.Swipe{},
		&models.PremiumPackage{},
		&models.UserPremium{},
	)
	if err != nil {
		return fmt.Errorf("failed to drop tables: %w", err)
	}
	log.Println("Dropped all tables successfully")

	// Enable UUID extension
	err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error
	if err != nil {
		return fmt.Errorf("failed to create uuid extension: %w", err)
	}

	// Re-run migrations
	err = db.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Swipe{},
		&models.PremiumPackage{},
		&models.UserPremium{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}

func main() {
	jwtSecret := "just_for_test"
	// Load configuration from the .env file
	cfg := config.LoadConfig()

	// Database connection string using environment variables
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run database migrations
	if err := autoMigrate(db); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepo(db)
	swipeRepo := repositories.NewSwipeRepo(db)
	premiumRepo := repositories.NewPremiumRepo(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, jwtSecret)
	swipeService := services.NewSwipeService(userRepo, swipeRepo)
	premiumService := services.NewPremiumService(premiumRepo)

	// Initialize router
	router := gin.Default()

	// Register routes
	routes.RegisterAuthRoutes(router, authService)
	routes.RegisterSwipeRoutes(router, swipeService, jwtSecret)
	routes.RegisterPremiumRoutes(router, premiumService)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
