package models

type Review struct {
	Model

	UserID    int64  `gorm:"type:bigint;not null"`
	ProductID int64  `gorm:"type:bigint;not null"`
	Rate      int64  `gorm:"type:bigint;not null"`
	Message   string `gorm:"type:character varying(255);not null"`
}
