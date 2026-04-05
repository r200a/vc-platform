package model

import "time"

type UserProfile struct {
	ProfileID    string    `db:"profile_id"`
	UserID       string    `db:"user_id"`
	Name         string    `db:"name"`
	Bio          string    `db:"bio"`
	Location     string    `db:"location"`
	ProfileImage string    `db:"profile_image"`
	LinkedinURL  string    `db:"linkedin_url"`
	CreatedAt    time.Time `db:"created_at"`
}

type CreateProfileRequest struct {
	Name        string `json:"name"         binding:"required"`
	Bio         string `json:"bio"`
	Location    string `json:"location"`
	LinkedinURL string `json:"linkedin_url"`
}

type UpdateProfileRequest struct {
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	Location    string `json:"location"`
	LinkedinURL string `json:"linkedin_url"`
}
