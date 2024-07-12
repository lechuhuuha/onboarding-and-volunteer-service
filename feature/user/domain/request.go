package domain

import "time"

type Request struct {
	ID          int
	UserID      int
	Type        string
	Status      int
	RejectNotes string
	VerifierID  int
	CreateAt    time.Time
	UpdateAt    time.Time
}
