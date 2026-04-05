package service

import (
	"github.com/r200a/vc-platform/internal/vc/model"
	"github.com/r200a/vc-platform/internal/vc/repository"
)

type VCService struct {
	repo *repository.VCRepository
}

func NewVCService(repo *repository.VCRepository) *VCService {
	return &VCService{repo: repo}
}

func (s *VCService) CreateVC(userID string, req model.CreateVCRequest) (string, error) {
	return s.repo.CreateVC(userID, req)
}

func (s *VCService) GetVCByID(vcID string) (model.VCProfile, error) {
	return s.repo.GetVCByID(vcID)
}

func (s *VCService) ListVCs(filter model.VCFilter) ([]model.VCProfile, error) {
	return s.repo.ListVCs(filter)
}

func (s *VCService) UpdateVC(userID string, req model.UpdateVCRequest) error {
	return s.repo.UpdateVC(userID, req)
}
