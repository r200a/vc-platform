package service

import (
	"errors"
	"fmt"

	"github.com/r200a/vc-platform/internal/application/model"
	"github.com/r200a/vc-platform/internal/application/repository"
)

type AppService struct {
	repo *repository.AppRepository
}

func NewAppService(repo *repository.AppRepository) *AppService {
	return &AppService{repo: repo}
}

var validTransitions = map[string][]string{
	"applied":     {"shortlisted", "rejected"},
	"shortlisted": {"pitching", "rejected"},
	"pitching":    {"funded", "rejected"},
	"funded":      {},
	"rejected":    {},
}

func (s *AppService) Apply(founderID string, req model.ApplyRequest) (string, error) {
	startup, err := s.repo.GetStartupByFounderID(founderID)
	if err != nil {
		return "", errors.New("founder has no startup profile — create one first")
	}
	return s.repo.Create(startup.StartupID, req.VCID, req.CoverNote)
}

func (s *AppService) GetFounderApplications(founderID string) ([]model.Application, error) {
	return s.repo.GetByFounderID(founderID)
}

func (s *AppService) GetVCApplications(vcUserID string) ([]model.Application, error) {
	return s.repo.GetByVCID(vcUserID)
}

/*func (s *AppService) UpdateStatus(applicationID, newStatus, rejectionNote string) error {
	app, err := s.repo.GetByID(applicationID)
	if err != nil {
		return errors.New("application not found")
	}

	allowed := validTransitions[app.Status]
	for _, s := range allowed {
		if s == newStatus {
			return s.repo.UpdateStatus(applicationID, newStatus, rejectionNote)
		}
	}

	return fmt.Errorf("invalid transition: cannot move from '%s' to '%s'", app.Status, newStatus)
}*/

func (s *AppService) UpdateStatus(applicationID, newStatus, rejectionNote string) error {
	app, err := s.repo.GetByID(applicationID)
	if err != nil {
		return errors.New("application not found")
	}

	allowed := validTransitions[app.Status]
	for _, next := range allowed { // ← renamed s to next
		if next == newStatus {
			return s.repo.UpdateStatus(applicationID, newStatus, rejectionNote) // ← s is AppService again
		}
	}

	return fmt.Errorf("invalid transition: cannot move from '%s' to '%s'", app.Status, newStatus)
}
