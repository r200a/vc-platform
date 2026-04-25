package service

import (
	"errors"
	"fmt"

	"github.com/r200a/vc-platform/internal/application/model"
	"github.com/r200a/vc-platform/internal/application/repository"
	"github.com/r200a/vc-platform/pkg/events"
)

type AppService struct {
	repo     *repository.AppRepository
	producer *events.Producer
}

func NewAppService(repo *repository.AppRepository, producer *events.Producer) *AppService {
	return &AppService{
		repo:     repo,
		producer: producer,
	}
}

var validTransitions = map[string][]string{
	"applied":     {"shortlisted", "rejected"},
	"shortlisted": {"pitching", "rejected"},
	"pitching":    {"funded", "rejected"},
	"funded":      {},
	"rejected":    {},
}

///Before Kafka
/*func (s *AppService) Apply(founderID string, req model.ApplyRequest) (string, error) {
	startup, err := s.repo.GetStartupByFounderID(founderID)
	if err != nil {
		return "", errors.New("founder has no startup profile — create one first")
	}
	return s.repo.Create(startup.StartupID, req.VCID, req.CoverNote)
}*/

func (s *AppService) Apply(founderID string, req model.ApplyRequest) (string, error) {
	startup, err := s.repo.GetStartupByFounderID(founderID)
	if err != nil {
		return "", errors.New("founder has no startup profile — create one first")
	}

	applicationID, err := s.repo.Create(startup.StartupID, req.VCID, req.CoverNote)
	if err != nil {
		return "", err
	}

	// publish event to Kafka
	s.producer.Publish(events.ApplicationEvent{
		EventType:     "application.submitted",
		ApplicationID: applicationID,
		StartupID:     startup.StartupID,
		Status:        "applied",
	})

	return applicationID, nil
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

///Before Kafka
/*func (s *AppService) UpdateStatus(applicationID, newStatus, rejectionNote string) error {
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
}*/

func (s *AppService) UpdateStatus(applicationID, newStatus, rejectionNote string) error {
	app, err := s.repo.GetByID(applicationID)
	if err != nil {
		return errors.New("application not found")
	}

	allowed := validTransitions[app.Status]
	for _, next := range allowed {
		if next == newStatus {
			// update DB first
			if err := s.repo.UpdateStatus(applicationID, newStatus, rejectionNote); err != nil {
				return err
			}

			// then publish event
			s.producer.Publish(events.ApplicationEvent{
				EventType:     "application." + newStatus,
				ApplicationID: applicationID,
				StartupID:     app.StartupID,
				VCID:          app.VCID,
				Status:        newStatus,
				RejectionNote: rejectionNote,
			})

			return nil
		}
	}

	return fmt.Errorf("invalid transition: cannot move from '%s' to '%s'", app.Status, newStatus)
}
