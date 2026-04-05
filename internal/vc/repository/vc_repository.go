package repository

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"github.com/r200a/vc-platform/internal/vc/model"
)

type VCRepository struct {
	db *sql.DB
}

func NewVCRepository(db *sql.DB) *VCRepository {
	return &VCRepository{db: db}
}

func (r *VCRepository) CreateVC(userID string, req model.CreateVCRequest) (string, error) {
	var vcID string
	err := r.db.QueryRow(
		`INSERT INTO vc_profiles
		 (user_id, fund_name, fund_size, ticket_size_min, ticket_size_max,
		  focus_industries, focus_stages, website_url)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		 RETURNING vc_id`,
		userID,
		req.FundName,
		req.FundSize,
		req.TicketSizeMin,
		req.TicketSizeMax,
		pq.Array(req.FocusIndustries),
		pq.Array(req.FocusStages),
		req.WebsiteURL,
	).Scan(&vcID)
	return vcID, err
}

func (r *VCRepository) GetVCByID(vcID string) (model.VCProfile, error) {
	var v model.VCProfile
	err := r.db.QueryRow(
		`SELECT vc_id, user_id, fund_name, fund_size, ticket_size_min,
		        ticket_size_max, focus_industries, focus_stages, website_url, created_at
		 FROM vc_profiles WHERE vc_id = $1`,
		vcID,
	).Scan(
		&v.VCID, &v.UserID, &v.FundName, &v.FundSize,
		&v.TicketSizeMin, &v.TicketSizeMax,
		pq.Array(&v.FocusIndustries),
		pq.Array(&v.FocusStages),
		&v.WebsiteURL, &v.CreatedAt,
	)
	return v, err
}

func (r *VCRepository) ListVCs(filter model.VCFilter) ([]model.VCProfile, error) {
	query := `SELECT vc_id, user_id, fund_name, fund_size, ticket_size_min,
	                 ticket_size_max, focus_industries, focus_stages, website_url, created_at
	          FROM vc_profiles WHERE 1=1`
	args := []interface{}{}
	i := 1

	if filter.Industry != "" {
		query += fmt.Sprintf(` AND $%d = ANY(focus_industries)`, i)
		args = append(args, filter.Industry)
		i++
	}
	if filter.MinTicket > 0 {
		query += fmt.Sprintf(` AND ticket_size_min >= $%d`, i)
		args = append(args, filter.MinTicket)
		i++
	}
	if filter.Stage != "" {
		query += fmt.Sprintf(` AND $%d = ANY(focus_stages)`, i)
		args = append(args, filter.Stage)
		i++
	}

	query += ` ORDER BY created_at DESC`

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vcs []model.VCProfile
	for rows.Next() {
		var v model.VCProfile
		if err := rows.Scan(
			&v.VCID, &v.UserID, &v.FundName, &v.FundSize,
			&v.TicketSizeMin, &v.TicketSizeMax,
			pq.Array(&v.FocusIndustries),
			pq.Array(&v.FocusStages),
			&v.WebsiteURL, &v.CreatedAt,
		); err != nil {
			return nil, err
		}
		vcs = append(vcs, v)
	}
	return vcs, nil
}

func (r *VCRepository) UpdateVC(userID string, req model.UpdateVCRequest) error {
	_, err := r.db.Exec(
		`UPDATE vc_profiles
		 SET fund_name=$1, fund_size=$2, ticket_size_min=$3, ticket_size_max=$4,
		     focus_industries=$5, focus_stages=$6, website_url=$7, updated_at=NOW()
		 WHERE user_id=$8`,
		req.FundName,
		req.FundSize,
		req.TicketSizeMin,
		req.TicketSizeMax,
		pq.Array(req.FocusIndustries),
		pq.Array(req.FocusStages),
		req.WebsiteURL,
		userID,
	)
	return err
}
