package cmd

import (
	"auth-service/internal/handler"
	repository "auth-service/internal/repository/db"
	routes "auth-service/internal/router"
	"auth-service/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

type Dependency struct {
	Config *Config
}

func InitDependencies(r *gin.Engine) *Dependency {
	cfg, err := Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// External dependencies
	dbConn, err := NewClientDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	// User module
	userRepository := repository.NewUserRepository(dbConn)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	// Setup Gin router
	routes.SetupRoutes(r, userHandler)

	return &Dependency{Config: cfg}
}
