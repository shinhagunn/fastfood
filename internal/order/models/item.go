package models

import "github.com/shopspring/decimal"

type OrderItem struct {
	Model

	OrderID   int64           `gorm:"type:bigint;not null"`
	ProductID int64           `gorm:"type:bigint;not null"`
	Name      string          `gorm:"type:character varying(50);not null"`
	Price     decimal.Decimal `gorm:"type:numeric(10,2);not null"`
	Amount    int64           `gorm:"type:bigint;not null"`
	Total     decimal.Decimal `gorm:"type:numeric(10,2);not null"`
}
