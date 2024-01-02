package model

import (
	"time"
)

type Project struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"not null"`
	Desc      string
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
