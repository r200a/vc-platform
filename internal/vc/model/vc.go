package model

import "time"

type VCProfile struct {
	VCID            string    `json:"vc_id"             db:"vc_id"`
	UserID          string    `json:"user_id"           db:"user_id"`
	FundName        string    `json:"fund_name"         db:"fund_name"`
	FundSize        int64     `json:"fund_size"         db:"fund_size"`
	TicketSizeMin   int64     `json:"ticket_size_min"   db:"ticket_size_min"`
	TicketSizeMax   int64     `json:"ticket_size_max"   db:"ticket_size_max"`
	FocusIndustries []string  `json:"focus_industries"  db:"focus_industries"`
	FocusStages     []string  `json:"focus_stages"      db:"focus_stages"`
	WebsiteURL      string    `json:"website_url"       db:"website_url"`
	CreatedAt       time.Time `json:"created_at"        db:"created_at"`
}

type CreateVCRequest struct {
	FundName        string   `json:"fund_name"       binding:"required"`
	FundSize        int64    `json:"fund_size"`
	TicketSizeMin   int64    `json:"ticket_size_min"`
	TicketSizeMax   int64    `json:"ticket_size_max"`
	FocusIndustries []string `json:"focus_industries"`
	FocusStages     []string `json:"focus_stages"`
	WebsiteURL      string   `json:"website_url"`
}

type UpdateVCRequest struct {
	FundName        string   `json:"fund_name"`
	FundSize        int64    `json:"fund_size"`
	TicketSizeMin   int64    `json:"ticket_size_min"`
	TicketSizeMax   int64    `json:"ticket_size_max"`
	FocusIndustries []string `json:"focus_industries"`
	FocusStages     []string `json:"focus_stages"`
	WebsiteURL      string   `json:"website_url"`
}

type VCFilter struct {
	Industry  string `form:"industry"`
	MinTicket int64  `form:"min_ticket"`
	Stage     string `form:"stage"`
}
