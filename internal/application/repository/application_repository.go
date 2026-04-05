package repository

import (
	"database/sql"

	"github.com/r200a/vc-platform/internal/application/model"
)

type AppRepository struct {
	db *sql.DB
}

func NewAppRepository(db *sql.DB) *AppRepository {
	return &AppRepository{db: db}
}

func (r *AppRepository) Create(startupID, vcID, coverNote string) (string, error) {
	var applicationID string
	err := r.db.QueryRow(
		`INSERT INTO applications (startup_id, vc_id, cover_note)
		 VALUES ($1, $2, $3)
		 RETURNING application_id`,
		startupID, vcID, coverNote,
	).Scan(&applicationID)
	return applicationID, err
}

func (r *AppRepository) GetByID(applicationID string) (model.Application, error) {
	var a model.Application
	err := r.db.QueryRow(
		`SELECT application_id, startup_id, vc_id, status,
		        cover_note, rejection_note, applied_at, updated_at
		 FROM applications WHERE application_id = $1`,
		applicationID,
	).Scan(
		&a.ApplicationID, &a.StartupID, &a.VCID, &a.Status,
		&a.CoverNote, &a.RejectionNote, &a.AppliedAt, &a.UpdatedAt,
	)
	return a, err
}

func (r *AppRepository) GetByFounderID(founderID string) ([]model.Application, error) {
	rows, err := r.db.Query(
		`SELECT a.application_id, a.startup_id, a.vc_id, a.status,
		        a.cover_note, a.rejection_note, a.applied_at, a.updated_at
		 FROM applications a
		 JOIN startup_profiles s ON a.startup_id = s.startup_id
		 WHERE s.founder_id = $1
		 ORDER BY a.applied_at DESC`,
		founderID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanApplications(rows)
}

func (r *AppRepository) GetByVCID(vcID string) ([]model.Application, error) {
	rows, err := r.db.Query(
		`SELECT a.application_id, a.startup_id, a.vc_id, a.status,
		        a.cover_note, a.rejection_note, a.applied_at, a.updated_at
		 FROM applications a
		 JOIN vc_profiles v ON a.vc_id = v.vc_id
		 WHERE v.user_id = $1
		 ORDER BY a.applied_at DESC`,
		vcID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanApplications(rows)
}

func (r *AppRepository) UpdateStatus(applicationID, status, rejectionNote string) error {
	_, err := r.db.Exec(
		`UPDATE applications
		 SET status=$1, rejection_note=$2, updated_at=NOW()
		 WHERE application_id=$3`,
		status, rejectionNote, applicationID,
	)
	return err
}

func scanApplications(rows *sql.Rows) ([]model.Application, error) {
	var apps []model.Application
	for rows.Next() {
		var a model.Application
		if err := rows.Scan(
			&a.ApplicationID, &a.StartupID, &a.VCID, &a.Status,
			&a.CoverNote, &a.RejectionNote, &a.AppliedAt, &a.UpdatedAt,
		); err != nil {
			return nil, err
		}
		apps = append(apps, a)
	}
	return apps, nil
}

// add to application/repository/application_repository.go
func (r *AppRepository) GetStartupByFounderID(founderID string) (model.Application, error) {
	var startupID string
	err := r.db.QueryRow(
		`SELECT startup_id FROM startup_profiles WHERE founder_id = $1`,
		founderID,
	).Scan(&startupID)
	return model.Application{StartupID: startupID}, err
}
