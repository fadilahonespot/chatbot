package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primarykey"`
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Chat      []Chat
}
