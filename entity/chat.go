package entity

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	ID        int `gorm:"primarykey"`
	UserID    int
	Name      string
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
