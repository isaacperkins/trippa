package users

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Email     string    `gorm:"not null"`
	Password  string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
