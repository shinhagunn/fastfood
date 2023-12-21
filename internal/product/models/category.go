package models

type Category struct {
	Model

	Name string `gorm:"type:character varying(50);not null;unique"`
}
