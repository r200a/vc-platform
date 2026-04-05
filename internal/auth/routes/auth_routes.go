package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r200a/vc-platform/internal/auth/handler"
)

func RegisterAuthRoutes(r *gin.Engine, h *handler.AuthHandler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/refresh-token", h.RefreshToken)
	}
}
