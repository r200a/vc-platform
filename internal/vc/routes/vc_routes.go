package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r200a/vc-platform/internal/vc/handler"
	"github.com/r200a/vc-platform/pkg/middleware"
)

func RegisterVCRoutes(r *gin.Engine, h *handler.VCHandler) {
	r.GET("/vc", h.ListVCs)
	r.GET("/vc/:id", h.GetVCByID)

	vc := r.Group("/vc", middleware.JWTAuth(), middleware.RequireRole("vc"))
	{
		vc.POST("", h.CreateVC)
		vc.PUT("/me", h.UpdateVC)
	}
}
