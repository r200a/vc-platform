package repository

import (
	"database/sql"

	"github.com/r200a/vc-platform/internal/user/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateProfile(userID string, req model.CreateProfileRequest) (string, error) {
	var profileID string
	err := r.db.QueryRow(
		`INSERT INTO user_profiles (user_id, name, bio, location, linkedin_url)
         VALUES ($1, $2, $3, $4, $5)
         RETURNING profile_id`,
		userID, req.Name, req.Bio, req.Location, req.LinkedinURL,
	).Scan(&profileID)
	return profileID, err
}

func (r *UserRepository) GetProfileByUserID(userID string) (model.UserProfile, error) {
	var p model.UserProfile
	err := r.db.QueryRow(
		`SELECT profile_id, user_id, name, bio, location, profile_image, linkedin_url
         FROM user_profiles WHERE user_id = $1`,
		userID,
	).Scan(&p.ProfileID, &p.UserID, &p.Name, &p.Bio, &p.Location, &p.ProfileImage, &p.LinkedinURL)
	return p, err
}

func (r *UserRepository) UpdateProfile(userID string, req model.UpdateProfileRequest) error {
	_, err := r.db.Exec(
		`UPDATE user_profiles
         SET name=$1, bio=$2, location=$3, linkedin_url=$4, updated_at=NOW()
         WHERE user_id=$5`,
		req.Name, req.Bio, req.Location, req.LinkedinURL, userID,
	)
	return err
}
