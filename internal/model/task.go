package model

import "time"

type Task struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title" `
	Done      bool      `gorm:"type:bool;not null" json:"done"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
