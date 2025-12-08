package model

import "time"

type Task struct {
	ID        int       `gorm:"primarykey" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title" `
	Status    string    `gorm:"type:varchar(255);not null" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
