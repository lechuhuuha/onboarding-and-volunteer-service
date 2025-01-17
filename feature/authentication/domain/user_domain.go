package domain

import (
	"time"
)

type User struct {
	ID                 int       `gorm:"primaryKey;autoIncrement"`
	RoleID             int       `gorm:"index;default:1"`
	DepartmentID       *int      `gorm:"index"`
	Email              string    `gorm:"unique;not null"`
	Password           string    `gorm:"not null"`
	Name               string    `gorm:"not null"`
	Surname            string    `gorm:"not null"`
	Gender             string    `gorm:"not null"`
	Dob                time.Time `gorm:"not null"`
	Mobile             string    `gorm:"not null"`
	CountryID          int       `gorm:"index;default:1"`
	ResidentCountryID  int       `gorm:"index;default:1"`
	Avatar             *string
	VerificationStatus int       `gorm:"default:0"`
	Status             int       `gorm:"not null"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}
