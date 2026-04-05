package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/r200a/vc-platform/internal/user/handler"
	"github.com/r200a/vc-platform/pkg/middleware"
)

func RegisterUserRoutes(r *gin.Engine, h *handler.UserHandler) {
	users := r.Group("/users", middleware.JWTAuth())
	{
		users.POST("", h.CreateProfile)
		users.GET("/:id", h.GetProfile)
		users.PUT("/me", h.UpdateProfile)
	}
}
