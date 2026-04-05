package model

import "time"

type StartupProfile struct {
	StartupID      string    `json:"startup_id"       db:"startup_id"`
	FounderID      string    `json:"founder_id"       db:"founder_id"`
	Name           string    `json:"name"             db:"name"`
	Tagline        string    `json:"tagline"          db:"tagline"`
	Industry       string    `json:"industry"         db:"industry"`
	Stage          string    `json:"stage"            db:"stage"`
	RevenueMonthly int64     `json:"revenue_monthly"  db:"revenue_monthly"`
	PitchDeckURL   string    `json:"pitch_deck_url"   db:"pitch_deck_url"`
	WebsiteURL     string    `json:"website_url"      db:"website_url"`
	TeamSize       int       `json:"team_size"        db:"team_size"`
	IsActive       bool      `json:"is_active"        db:"is_active"`
	CreatedAt      time.Time `json:"created_at"       db:"created_at"`
}

type CreateStartupRequest struct {
	Name           string `json:"name"            binding:"required"`
	Tagline        string `json:"tagline"`
	Industry       string `json:"industry"        binding:"required"`
	Stage          string `json:"stage"           binding:"required,oneof=idea pre-seed seed series-a series-b"`
	RevenueMonthly int64  `json:"revenue_monthly"`
	WebsiteURL     string `json:"website_url"`
	TeamSize       int    `json:"team_size"`
}

type UpdateStartupRequest struct {
	Name           string `json:"name"`
	Tagline        string `json:"tagline"`
	Industry       string `json:"industry"`
	Stage          string `json:"stage"`
	RevenueMonthly int64  `json:"revenue_monthly"`
	WebsiteURL     string `json:"website_url"`
	TeamSize       int    `json:"team_size"`
}

type StartupFilter struct {
	Industry string `form:"industry"`
	Stage    string `form:"stage"`
}

type PitchDeckURLResponse struct {
	UploadURL string `json:"upload_url"`
	StartupID string `json:"startup_id"`
}
