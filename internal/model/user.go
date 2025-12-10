package model

import "time"

type User struct {
	Id        uint      `gorm:"primarykey"`
	Name      string    `gorm:"not null;type:varchar(100) required:true min:3 max:100"`
	Email     string    `gorm:"not null;unique;type:varchar(100) required:true min:3 max:100"`
	Password  string    `gorm:"not null;type:varchar(100) required:true min:6 max:100"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime omitempty"`
}
