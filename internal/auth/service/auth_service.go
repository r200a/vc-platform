package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/r200a/vc-platform/internal/auth/model"
	"github.com/r200a/vc-platform/internal/auth/repository"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(req model.RegisterRequest) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return "", errors.New("failed to process password")
	}
	return s.repo.CreateUser(req.Email, string(hash), req.Role)
}

func (s *AuthService) Login(req model.LoginRequest) (model.AuthResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return model.AuthResponse{}, errors.New("invalid credentials")
	}

	if !user.IsActive {
		return model.AuthResponse{}, errors.New("account suspended")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash), []byte(req.Password),
	); err != nil {
		return model.AuthResponse{}, errors.New("invalid credentials")
	}

	accessToken, err := generateJWT(user.UserID, user.Role, 15*time.Minute)
	if err != nil {
		return model.AuthResponse{}, err
	}

	refreshToken, err := generateJWT(user.UserID, user.Role, 7*24*time.Hour)
	if err != nil {
		return model.AuthResponse{}, err
	}

	if err := s.repo.SaveRefreshToken(user.UserID, refreshToken); err != nil {
		return model.AuthResponse{}, err
	}

	return model.AuthResponse{
		UserID:       user.UserID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Role:         user.Role,
	}, nil
}

func (s *AuthService) RefreshToken(refreshToken string) (string, error) {
	user, err := s.repo.GetUserByRefreshToken(refreshToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	return generateJWT(user.UserID, user.Role, 15*time.Minute)
}

func generateJWT(userID, role string, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(expiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
