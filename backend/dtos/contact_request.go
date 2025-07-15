package dtos

import "time"

type ContactRequest struct {
	ID       int       `json:"id"`
	UserID   int       `json:"user_id"`
	UpdateAt time.Time `json:"update_at"`
	FullName string    `json:"full_name"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	Address  string    `json:"address"`
}
