package domain

import (
	"time"
)

type User struct {
	ID                 int
	RoleID             int
	DepartmentID       *int
	Email              string
	Password           string
	Name               string
	Surname            string
	Gender             string
	DOB                time.Time
	Mobile             string
	CountryID          int
	ResidentCountryID  int
	Avatar             *string
	VerificationStatus int
	Status             int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
