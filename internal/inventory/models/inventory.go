package models

type InventoryState string

const (
	InventoryStateAvailable  = InventoryState("available")
	InventoryStateOutOfStock = InventoryState("out_of_stock")
)

type Inventory struct {
	Model

	ProductID int64          `gorm:"type:bigint;not null"`
	Quantity  int64          `gorm:"type:bigint;not null"`
	State     InventoryState `gorm:"type:character varying(20);not null"`
}
