package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	FirstName string
	MiddleName string
	LastName string
	StudentID string
	Age uint8
	DOB time.Time
	Gender string
	IsCurrentStudent bool
}