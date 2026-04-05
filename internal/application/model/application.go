package model

import "time"

type Application struct {
	ApplicationID string    `json:"application_id"  db:"application_id"`
	StartupID     string    `json:"startup_id"      db:"startup_id"`
	VCID          string    `json:"vc_id"           db:"vc_id"`
	Status        string    `json:"status"          db:"status"`
	CoverNote     string    `json:"cover_note"      db:"cover_note"`
	RejectionNote string    `json:"rejection_note"  db:"rejection_note"`
	AppliedAt     time.Time `json:"applied_at"      db:"applied_at"`
	UpdatedAt     time.Time `json:"updated_at"      db:"updated_at"`
}

type ApplyRequest struct {
	VCID      string `json:"vc_id"      binding:"required"`
	CoverNote string `json:"cover_note"`
}

type UpdateStatusRequest struct {
	Status        string `json:"status"         binding:"required,oneof=shortlisted pitching funded rejected"`
	RejectionNote string `json:"rejection_note"`
}
