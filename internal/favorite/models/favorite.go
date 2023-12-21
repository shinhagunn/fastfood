package models

type Favorite struct {
	Model

	UserID    int64 `gorm:"type:bigint;not null"`
	ProductID int64 `gorm:"type:bigint;not null"`
}
