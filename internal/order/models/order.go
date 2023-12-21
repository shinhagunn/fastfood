package models

import "github.com/shopspring/decimal"

type OrderStatus string

var (
	OrderStatusPending = OrderStatus("pending")
	OrderStatusSucceed = OrderStatus("Succeed")
)

type Order struct {
	Model

	UserID int64           `gorm:"type:bigint;not null"`
	Status OrderStatus     `gorm:"type:character varying(10);not null"`
	Total  decimal.Decimal `gorm:"type:numeric(10,2);not null"`
}
