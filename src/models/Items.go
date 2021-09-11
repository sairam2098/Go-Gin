package models

type Items struct {
	Name     string `db:"name,omitempty"`
	Price    string `db:"price,omitempty"`
	Quantity int    `db:"quantity,omitempty"`
}

type CategoryStruct struct {
	Category string `form:"category"`
}
