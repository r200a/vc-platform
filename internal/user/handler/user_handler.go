package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r200a/vc-platform/internal/user/model"
	"github.com/r200a/vc-platform/internal/user/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) CreateProfile(c *gin.Context) {
	userID := c.GetString("user_id") // set by JWT middleware
	var req model.CreateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profileID, err := h.service.CreateProfile(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create profile"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"profile_id": profileID})
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.Param("id")

	profile, err := h.service.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "profile not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	var req model.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateProfile(userID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "profile updated"})
}
