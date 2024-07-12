package domain

import "time"

type Admin struct {
	ID                 int       `gorm:"primaryKey"`
	RoleID             int       `gorm:"index"`
	DepartmentID       *int      `gorm:"index"`
	Email              string    `gorm:"unique;not null"`
	Password           string    `gorm:"not null"`
	Name               string    `gorm:"not null"`
	Surname            string    `gorm:"not null"`
	Gender             string    `gorm:"not null"`
	Dob                time.Time `gorm:"not null"`
	Mobile             string    `gorm:"not null"`
	CountryID          int       `gorm:"index"`
	ResidentCountryID  int       `gorm:"index"`
	Avatar             *string
	VerificationStatus int       `gorm:"default:0"`
	Status             int       `gorm:"not null"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}

type Request struct {
	ID          int    `gorm:"primaryKey"`
	UserID      int    `gorm:"index"`
	Type        string `gorm:"not null"`
	Status      int    `gorm:"not null"`
	RejectNotes string
	VerifierID  int       `gorm:"index"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
