package main

import (
	"git-functionality/config"
	"git-functionality/internal/handlers"
	"git-functionality/internal/repository"
	"git-functionality/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Setup database
	db, err := config.InitDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize repositories
	userRepository := repository.NewUserRepository(db)
	versionRepository := repository.NewVersionRepository(db)

	// Initialize services
	userService := service.NewUserService(db)
	versionService := service.NewVersionService(
		db,
		userRepository,
		versionRepository,
	)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(
		userService,
		versionService,
	)

	// Setup Gin router
	r := gin.Default()

	// API Group
	v1 := r.Group("/api/v1")
	{
		// User routes
		v1.GET("/users", userHandler.GetUsers)
		v1.POST("/users", userHandler.CreateUser)
		v1.GET("/users/:id", userHandler.GetUserByID)
		v1.POST("/users/:id/versions", userHandler.CommitChanges)
		v1.GET("/users/:id/versions", userHandler.GetVersionHistory)
	}

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
