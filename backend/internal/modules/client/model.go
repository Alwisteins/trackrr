package client

import "time"

type Client struct {
	ID           int64     `json:"id" db:"id"`
	UUID         string    `json:"uuid" db:"uuid"`
	CompanyName  string    `json:"company_name" db:"company_name"`
	ContactName  string    `json:"contact_name" db:"contact_name"`
	ContactEmail string    `json:"contact_email" db:"contact_email"`
	ContactPhone string    `json:"contact_phone" db:"contact_phone"`
	Notes        string    `json:"notes" db:"notes"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type ClientRequest struct {
	CompanyName  string `json:"company_name" binding:"required,min=2"`
	ContactName  string `json:"contact_name" binding:"required"`
	ContactEmail string `json:"contact_email" binding:"omitempty,email"`
	ContactPhone string `json:"contact_phone" binding:"required"`
	Notes        string `json:"notes"`
}
