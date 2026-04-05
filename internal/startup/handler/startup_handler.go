package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r200a/vc-platform/internal/startup/model"
	"github.com/r200a/vc-platform/internal/startup/service"
)

type StartupHandler struct {
	service *service.StartupService
}

func NewStartupHandler(s *service.StartupService) *StartupHandler {
	return &StartupHandler{service: s}
}

func (h *StartupHandler) CreateStartup(c *gin.Context) {
	founderID := c.GetString("user_id")

	var req model.CreateStartupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startupID, err := h.service.CreateStartup(founderID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create startup"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"startup_id": startupID})
}

func (h *StartupHandler) GetStartup(c *gin.Context) {
	startupID := c.Param("id")

	startup, err := h.service.GetStartup(startupID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "startup not found"})
		return
	}

	c.JSON(http.StatusOK, startup)
}

func (h *StartupHandler) ListStartups(c *gin.Context) {
	var filter model.StartupFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startups, err := h.service.ListStartups(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch startups"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  startups,
		"count": len(startups),
	})
}

func (h *StartupHandler) UpdateStartup(c *gin.Context) {
	founderID := c.GetString("user_id")

	var req model.UpdateStartupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateStartup(founderID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update startup"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "startup updated"})
}

func (h *StartupHandler) GetPitchDeckUploadURL(c *gin.Context) {
	startupID := c.Param("id")

	url, err := h.service.GeneratePitchDeckUploadURL(startupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate upload URL"})
		return
	}

	c.JSON(http.StatusOK, model.PitchDeckURLResponse{
		UploadURL: url,
		StartupID: startupID,
	})
}
