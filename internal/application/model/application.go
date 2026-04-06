package model

import (
	"database/sql"
	"time"
)

type Application struct {
	ApplicationID string         `json:"application_id"`
	StartupID     string         `json:"startup_id"`
	VCID          string         `json:"vc_id"`
	Status        string         `json:"status"`
	CoverNote     string         `json:"cover_note"`
	RejectionNote sql.NullString `json:"-"` // internal use only
	AppliedAt     time.Time      `json:"applied_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

// clean JSON response — rejection_note shows as null or "string"
type ApplicationResponse struct {
	ApplicationID string    `json:"application_id"`
	StartupID     string    `json:"startup_id"`
	VCID          string    `json:"vc_id"`
	Status        string    `json:"status"`
	CoverNote     string    `json:"cover_note"`
	RejectionNote *string   `json:"rejection_note"`
	AppliedAt     time.Time `json:"applied_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ApplyRequest struct {
	VCID      string `json:"vc_id"      binding:"required"`
	CoverNote string `json:"cover_note"`
}

type UpdateStatusRequest struct {
	Status        string `json:"status"         binding:"required,oneof=shortlisted pitching funded rejected"`
	RejectionNote string `json:"rejection_note"`
}

func ToResponse(a Application) ApplicationResponse {
	var note *string
	if a.RejectionNote.Valid {
		note = &a.RejectionNote.String
	}
	return ApplicationResponse{
		ApplicationID: a.ApplicationID,
		StartupID:     a.StartupID,
		VCID:          a.VCID,
		Status:        a.Status,
		CoverNote:     a.CoverNote,
		RejectionNote: note,
		AppliedAt:     a.AppliedAt,
		UpdatedAt:     a.UpdatedAt,
	}
}
