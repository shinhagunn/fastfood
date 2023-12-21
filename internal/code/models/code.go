package models

import "time"

type CodeType string

var (
	CodeTypeEmail = CodeType("email")
	CodeTypePhone = CodeType("phone")
)

type Code struct {
	Model

	UserID       int64     `gorm:"type:bigint;not null"`
	Type         CodeType  `gorm:"type:character varying(10);not null"`
	AttemptCount int64     `gorm:"type:bigint;not null"`
	ValidatedAt  time.Time `gorm:"type:timestamp;not null"`
	ExpriedAt    time.Time `gorm:"type:timestamp;not null"`
}
