package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r200a/vc-platform/internal/application/model"
	"github.com/r200a/vc-platform/internal/application/service"
)

type AppHandler struct {
	service *service.AppService
}

func NewAppHandler(s *service.AppService) *AppHandler {
	return &AppHandler{service: s}
}

func (h *AppHandler) Apply(c *gin.Context) {
	founderID := c.GetString("user_id")

	var req model.ApplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	applicationID, err := h.service.Apply(founderID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"application_id": applicationID,
		"message":        "application submitted",
	})
}

func (h *AppHandler) GetFounderApplications(c *gin.Context) {
	founderID := c.GetString("user_id")

	apps, err := h.service.GetFounderApplications(founderID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]model.ApplicationResponse, 0)
	for _, a := range apps {
		resp = append(resp, model.ToResponse(a))
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  resp,
		"count": len(resp),
	})
}

func (h *AppHandler) GetVCApplications(c *gin.Context) {
	vcUserID := c.GetString("user_id")

	apps, err := h.service.GetVCApplications(vcUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]model.ApplicationResponse, 0)
	for _, a := range apps {
		resp = append(resp, model.ToResponse(a))
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  resp,
		"count": len(resp),
	})
}
func (h *AppHandler) UpdateStatus(c *gin.Context) {
	applicationID := c.Param("id")

	var req model.UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateStatus(applicationID, req.Status, req.RejectionNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "status updated to " + req.Status,
	})
}
