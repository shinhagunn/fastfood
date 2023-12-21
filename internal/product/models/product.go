package models

import "github.com/shopspring/decimal"

type ProductStatus string

var (
	ProductStatusEnabled  ProductStatus = "enabled"
	ProductStatusDisabled ProductStatus = "disabled"
)

type Product struct {
	Model

	CategoryID int64 `gorm:"type:bigint;not null"`

	Name        string          `gorm:"type:character varying(50);not null"`
	Description string          `gorm:"type:character varying(255);not null"`
	Price       decimal.Decimal `gorm:"type:numeric(10,2);not null"`
	ImageUrl    string          `gorm:"type:character varying(255)"`
	Status      ProductStatus   `gorm:"type:character varying(10);not null"`
}
