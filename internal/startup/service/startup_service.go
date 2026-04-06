package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/r200a/vc-platform/internal/startup/model"
	"github.com/r200a/vc-platform/internal/startup/repository"
)

type StartupService struct {
	repo *repository.StartupRepository
}

func NewStartupService(repo *repository.StartupRepository) *StartupService {
	return &StartupService{repo: repo}
}

func (s *StartupService) CreateStartup(founderID string, req model.CreateStartupRequest) (string, error) {
	return s.repo.CreateStartup(founderID, req)
}

func (s *StartupService) GetStartup(startupID string) (model.StartupProfile, error) {
	return s.repo.GetStartupByID(startupID)
}

func (s *StartupService) ListStartups(filter model.StartupFilter) ([]model.StartupProfile, error) {
	return s.repo.ListStartups(filter)
}

func (s *StartupService) UpdateStartup(founderID string, req model.UpdateStartupRequest) error {
	return s.repo.UpdateStartup(founderID, req)
}

func (s *StartupService) GeneratePitchDeckUploadURL(startupID string) (string, error) {
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(cfg)
	presigner := s3.NewPresignClient(client)
	key := fmt.Sprintf("pitch-decks/%s.pdf", startupID)

	req, err := presigner.PresignPutObject(context.TODO(),
		&s3.PutObjectInput{
			Bucket: aws.String(os.Getenv("S3_BUCKET")),
			Key:    aws.String(key),
		},
		s3.WithPresignExpires(15*time.Minute),
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	publicURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s",
		os.Getenv("S3_BUCKET"), key)
	s.repo.UpdatePitchDeckURL(startupID, publicURL)

	return req.URL, nil
}

// This is for local setup
/*type StartupService struct {
	repo *repository.StartupRepository
}

func NewStartupService(repo *repository.StartupRepository) *StartupService {
	return &StartupService{repo: repo}
}

func (s *StartupService) CreateStartup(founderID string, req model.CreateStartupRequest) (string, error) {
	return s.repo.CreateStartup(founderID, req)
}

func (s *StartupService) GetStartup(startupID string) (model.StartupProfile, error) {
	return s.repo.GetStartupByID(startupID)
}

func (s *StartupService) ListStartups(filter model.StartupFilter) ([]model.StartupProfile, error) {
	return s.repo.ListStartups(filter)
}

func (s *StartupService) UpdateStartup(founderID string, req model.UpdateStartupRequest) error {
	return s.repo.UpdateStartup(founderID, req)
}

// S3 not configured yet — returns placeholder
func (s *StartupService) GeneratePitchDeckUploadURL(startupID string) (string, error) {
	return "s3-not-configured-yet", nil
}
*/
