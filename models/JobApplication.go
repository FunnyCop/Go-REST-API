package models

import "database/sql"

// Job Application Struct
type JobApplication struct {
	ID             sql.NullInt64  `json:"id"`
	CompanyName    sql.NullString `json:"company_name"`
	CompanyWebsite sql.NullString `json:"company_website"`
	CreatedAt      sql.NullString `json:"created_at"`
	Description    sql.NullString `json:"description"`
	Interviews     sql.NullInt32  `json:"interviews"`
	JobTitle       sql.NullString `json:"job_title"`
	Link           sql.NullString `json:"link"`
	Offer          sql.NullString `json:"offer"`
	Status         sql.NullString `json:"status"`
	UpdatedAt      sql.NullString `json:"updated_at"`
}
