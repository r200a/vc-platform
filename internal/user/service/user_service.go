package service

import (
	"github.com/r200a/vc-platform/internal/user/model"
	"github.com/r200a/vc-platform/internal/user/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateProfile(userID string, req model.CreateProfileRequest) (string, error) {
	return s.repo.CreateProfile(userID, req)
}

func (s *UserService) GetProfile(userID string) (model.UserProfile, error) {
	return s.repo.GetProfileByUserID(userID)
}

func (s *UserService) UpdateProfile(userID string, req model.UpdateProfileRequest) error {
	return s.repo.UpdateProfile(userID, req)
}
