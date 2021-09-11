package models

type Items struct {
	Name     string `db:"name,omitempty"`
	Price    int    `db:"price,omitempty"`
	Quantity int    `db:"quantity,omitempty"`
}

type CategoryStruct struct {
	Category string `form:"category"`
}

type NewItems struct {
	Category string `db:"category,omitempty"`
	Name     string `db:"name,omitempty"`
	Price    string `db:"price,omitempty"`
	Quantity string `db:"quantity,omitempty"`
}
