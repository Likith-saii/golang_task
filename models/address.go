package models

type Address struct {
	ID      int    `json:"id" db:"id"`
	City    string `json:"city" db:"city"`
	State   string `json:"state" db:"state"`
	Street1 string `json:"street1" db:"street1"`
	Street2 string `json:"street2,omitempty" db:"street2"`
	ZipCode string `json:"zip_code" db:"zip_code"`
}
