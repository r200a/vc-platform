package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/r200a/vc-platform/internal/vc/model"
	"github.com/r200a/vc-platform/internal/vc/service"
)

type VCHandler struct {
	service *service.VCService
}

func NewVCHandler(s *service.VCService) *VCHandler {
	return &VCHandler{service: s}
}

func (h *VCHandler) CreateVC(c *gin.Context) {
	userID := c.GetString("user_id")

	var req model.CreateVCRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vcID, err := h.service.CreateVC(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create VC profile"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"vc_id": vcID})
}

func (h *VCHandler) GetVCByID(c *gin.Context) {
	vcID := c.Param("id")

	vc, err := h.service.GetVCByID(vcID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "VC not found"})
		return
	}

	c.JSON(http.StatusOK, vc)
}

func (h *VCHandler) ListVCs(c *gin.Context) {
	var filter model.VCFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vcs, err := h.service.ListVCs(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch VCs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  vcs,
		"count": len(vcs),
	})
}

func (h *VCHandler) UpdateVC(c *gin.Context) {
	userID := c.GetString("user_id")

	var req model.UpdateVCRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateVC(userID, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update VC profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "VC profile updated"})
}
