package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r200a/vc-platform/internal/startup/handler"
	"github.com/r200a/vc-platform/pkg/middleware"
)

func RegisterStartupRoutes(r *gin.Engine, h *handler.StartupHandler) {
	r.GET("/startup", h.ListStartups)
	r.GET("/startup/:id", h.GetStartup)

	startup := r.Group("/startup", middleware.JWTAuth(), middleware.RequireRole("founder"))
	{
		startup.POST("", h.CreateStartup)
		startup.PUT("/me", h.UpdateStartup)
		startup.POST("/:id/pitch-deck", h.GetPitchDeckUploadURL)
	}
}
