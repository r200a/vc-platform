package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/r200a/vc-platform/pkg/config"
	"github.com/r200a/vc-platform/storage/db"

	authHdlr "github.com/r200a/vc-platform/internal/auth/handler"
	authRepo "github.com/r200a/vc-platform/internal/auth/repository"
	authRoutes "github.com/r200a/vc-platform/internal/auth/routes"
	authSvc "github.com/r200a/vc-platform/internal/auth/service"

	userHdlr "github.com/r200a/vc-platform/internal/user/handler"
	userRepo "github.com/r200a/vc-platform/internal/user/repository"
	userRoutes "github.com/r200a/vc-platform/internal/user/routes"
	userSvc "github.com/r200a/vc-platform/internal/user/service"

	vcHdlr "github.com/r200a/vc-platform/internal/vc/handler"
	vcRepo "github.com/r200a/vc-platform/internal/vc/repository"
	vcRoutes "github.com/r200a/vc-platform/internal/vc/routes"
	vcSvc "github.com/r200a/vc-platform/internal/vc/service"

	startupHdlr "github.com/r200a/vc-platform/internal/startup/handler"
	startupRepo "github.com/r200a/vc-platform/internal/startup/repository"
	startupRoutes "github.com/r200a/vc-platform/internal/startup/routes"
	startupSvc "github.com/r200a/vc-platform/internal/startup/service"

	appHdlr "github.com/r200a/vc-platform/internal/application/handler"
	appRepo "github.com/r200a/vc-platform/internal/application/repository"
	appRoutes "github.com/r200a/vc-platform/internal/application/routes"
	appSvc "github.com/r200a/vc-platform/internal/application/service"
)

func main() {
	// Load Config,
	cfg := config.Load()
	// making git main
	// Database Setup
	database := db.Connect(cfg.DBURL)

	// Setup Router
	// auth
	aRepo := authRepo.NewAuthRepository(database)
	aSvc := authSvc.NewAuthService(aRepo)
	aHdlr := authHdlr.NewAuthHandler(aSvc)

	// user
	uRepo := userRepo.NewUserRepository(database)
	uSvc := userSvc.NewUserService(uRepo)
	uHdlr := userHdlr.NewUserHandler(uSvc)

	// vc
	vRepo := vcRepo.NewVCRepository(database)
	vSvc := vcSvc.NewVCService(vRepo)
	vHdlr := vcHdlr.NewVCHandler(vSvc)

	// startup
	sRepo := startupRepo.NewStartupRepository(database)
	sSvc := startupSvc.NewStartupService(sRepo)
	sHdlr := startupHdlr.NewStartupHandler(sSvc)

	// application
	appRepository := appRepo.NewAppRepository(database)
	appService := appSvc.NewAppService(appRepository)
	appHandler := appHdlr.NewAppHandler(appService)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"Status": "OK", "service": "VC"}) })
	r.GET("/", func(c *gin.Context) { c.JSONP(200, gin.H{"VC": "test"}) })

	authRoutes.RegisterAuthRoutes(r, aHdlr)
	userRoutes.RegisterUserRoutes(r, uHdlr)
	vcRoutes.RegisterVCRoutes(r, vHdlr)
	startupRoutes.RegisterStartupRoutes(r, sHdlr)
	appRoutes.RegisterAppRoutes(r, appHandler)

	// Setup Server
	fmt.Println("Server starting on:", cfg.Port)
	r.Run(":" + cfg.Port)
}
