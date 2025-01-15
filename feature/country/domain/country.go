package domain

import (
	"time"
)

// Country struct that interacts with databases (GORM)
type Country struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Status    uint      `gorm:"not null" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Request struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Status    uint      `gorm:"not null" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
