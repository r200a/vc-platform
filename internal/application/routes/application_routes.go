package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r200a/vc-platform/internal/application/handler"
	"github.com/r200a/vc-platform/pkg/middleware"
)

func RegisterAppRoutes(r *gin.Engine, h *handler.AppHandler) {
	app := r.Group("/applications", middleware.JWTAuth())
	{
		app.POST("", middleware.RequireRole("founder"), h.Apply)
		app.GET("/my", middleware.RequireRole("founder"), h.GetFounderApplications)
		app.GET("/incoming", middleware.RequireRole("vc"), h.GetVCApplications)
		app.PATCH("/:id/status", middleware.RequireRole("vc"), h.UpdateStatus)
	}
}
