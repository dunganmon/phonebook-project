package models

import "time"

type Contact struct {
	ID       uint      `json:"id" gorm:"primaryKey"`
	FullName string    `json:"full_name"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	UserId   int       `json:"user_id"`
	UpdateAt time.Time `json:"update_at"`
	Address  string    `json:"address"`
}
