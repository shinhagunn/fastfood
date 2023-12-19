package models

import "github.com/shopspring/decimal"

type Product struct {
	Model

	CategoryID int64 `gorm:"type:bigint;not null"`

	Name  string          `gorm:"type:character varying(50);not null;unique"`
	Price decimal.Decimal `gorm:"type:numeric(10,2);not null"`
}
