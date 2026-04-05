package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/r200a/vc-platform/pkg/config"
	"github.com/r200a/vc-platform/storage/db"

	authHandler "github.com/r200a/vc-platform/internal/auth/handler"
	authRepo "github.com/r200a/vc-platform/internal/auth/repository"
	authRoutes "github.com/r200a/vc-platform/internal/auth/routes"
	authService "github.com/r200a/vc-platform/internal/auth/service"
)

func main() {
	// Load Config,
	cfg := config.Load()
	// making git main
	// Database Setup
	database := db.Connect(cfg.DBURL)

	// Setup Router
	// auth wiring
	aRepo := authRepo.NewAuthRepository(database)
	aSvc := authService.NewAuthService(aRepo)
	aHdlr := authHandler.NewAuthHandler(aSvc)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"Status": "OK", "service": "VC"}) })
	r.GET("/", func(c *gin.Context) { c.JSONP(200, gin.H{"VC": "test"}) })
	authRoutes.RegisterAuthRoutes(r, aHdlr)

	// Setup Server
	fmt.Println("Server starting on:", cfg.Port)
	r.Run(":" + cfg.Port)
}
