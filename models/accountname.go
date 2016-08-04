package models

import "time"

type AccountName struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	UserID    uint      `json:"user_id"`
	Name      string    `json:"name"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
