package models

type UserInformation struct {
	Model

	UserID int64 `gorm:"type:bigint;not null"`

	Fullname string `gorm:"type:character varying(50);not null"`
	Email    string `gorm:"type:character varying(50)"`
}
