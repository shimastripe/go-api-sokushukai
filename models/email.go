package models

import "time"

type Email struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID    uint      `json:"user_id"`
	Email     string    `json:"email"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
