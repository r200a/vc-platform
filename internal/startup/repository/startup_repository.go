package repository

import (
	"database/sql"
	"fmt"

	"github.com/r200a/vc-platform/internal/startup/model"
)

type StartupRepository struct {
	db *sql.DB
}

func NewStartupRepository(db *sql.DB) *StartupRepository {
	return &StartupRepository{db: db}
}

func (r *StartupRepository) CreateStartup(founderID string, req model.CreateStartupRequest) (string, error) {
	var startupID string
	err := r.db.QueryRow(
		`INSERT INTO startup_profiles
		 (founder_id, name, tagline, industry, stage, revenue_monthly, website_url, team_size)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		 RETURNING startup_id`,
		founderID,
		req.Name,
		req.Tagline,
		req.Industry,
		req.Stage,
		req.RevenueMonthly,
		req.WebsiteURL,
		req.TeamSize,
	).Scan(&startupID)
	return startupID, err
}

func (r *StartupRepository) GetStartupByID(startupID string) (model.StartupProfile, error) {
	var s model.StartupProfile
	err := r.db.QueryRow(
		`SELECT startup_id, founder_id, name, tagline, industry, stage,
		        revenue_monthly, pitch_deck_url, website_url, team_size, is_active, created_at
		 FROM startup_profiles WHERE startup_id = $1`,
		startupID,
	).Scan(
		&s.StartupID, &s.FounderID, &s.Name, &s.Tagline,
		&s.Industry, &s.Stage, &s.RevenueMonthly, &s.PitchDeckURL,
		&s.WebsiteURL, &s.TeamSize, &s.IsActive, &s.CreatedAt,
	)
	return s, err
}

func (r *StartupRepository) ListStartups(filter model.StartupFilter) ([]model.StartupProfile, error) {
	query := `SELECT startup_id, founder_id, name, tagline, industry, stage,
	                 revenue_monthly, pitch_deck_url, website_url, team_size, is_active, created_at
	          FROM startup_profiles WHERE is_active = true`
	args := []interface{}{}
	i := 1

	if filter.Industry != "" {
		query += fmt.Sprintf(` AND industry = $%d`, i)
		args = append(args, filter.Industry)
		i++
	}
	if filter.Stage != "" {
		query += fmt.Sprintf(` AND stage = $%d`, i)
		args = append(args, filter.Stage)
		i++
	}

	query += ` ORDER BY created_at DESC`

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var startups []model.StartupProfile
	for rows.Next() {
		var s model.StartupProfile
		if err := rows.Scan(
			&s.StartupID, &s.FounderID, &s.Name, &s.Tagline,
			&s.Industry, &s.Stage, &s.RevenueMonthly, &s.PitchDeckURL,
			&s.WebsiteURL, &s.TeamSize, &s.IsActive, &s.CreatedAt,
		); err != nil {
			return nil, err
		}
		startups = append(startups, s)
	}
	return startups, nil
}

func (r *StartupRepository) UpdateStartup(founderID string, req model.UpdateStartupRequest) error {
	_, err := r.db.Exec(
		`UPDATE startup_profiles
		 SET name=$1, tagline=$2, industry=$3, stage=$4,
		     revenue_monthly=$5, website_url=$6, team_size=$7, updated_at=NOW()
		 WHERE founder_id=$8`,
		req.Name, req.Tagline, req.Industry, req.Stage,
		req.RevenueMonthly, req.WebsiteURL, req.TeamSize,
		founderID,
	)
	return err
}

func (r *StartupRepository) UpdatePitchDeckURL(startupID, url string) error {
	_, err := r.db.Exec(
		`UPDATE startup_profiles SET pitch_deck_url=$1, updated_at=NOW()
		 WHERE startup_id=$2`,
		url, startupID,
	)
	return err
}

func (r *StartupRepository) GetStartupByFounderID(founderID string) (model.StartupProfile, error) {
	var s model.StartupProfile
	err := r.db.QueryRow(
		`SELECT startup_id, founder_id, name, tagline, industry, stage,
		        revenue_monthly, pitch_deck_url, website_url, team_size, is_active, created_at
		 FROM startup_profiles WHERE founder_id = $1`,
		founderID,
	).Scan(
		&s.StartupID, &s.FounderID, &s.Name, &s.Tagline,
		&s.Industry, &s.Stage, &s.RevenueMonthly, &s.PitchDeckURL,
		&s.WebsiteURL, &s.TeamSize, &s.IsActive, &s.CreatedAt,
	)
	return s, err
}
