package domain

import (
	"time"
	"gorm.io/gorm"
)

type Volunteer struct {
	gorm.Model
	Id           uint `gorm:"primaryKey"`
	UserId       uint `gorm:"notnull"`
	DepartmentId uint `gorm:"notnull"`
	Status       uint `gorm:"notnull"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
