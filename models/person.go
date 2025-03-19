package models

type PersonResponse struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	City        string `json:"city"`
	State       string `json:"state"`
	Street1     string `json:"street1"`
	Street2     string `json:"street2"`
	ZipCode     string `json:"zip_code"`
}

type PersonRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=100"`
	PhoneNumber string `json:"phone_number" binding:"required,len=12,regexp=^\\d{3}-\\d{3}-\\d{4}$"`
	City        string `json:"city" binding:"required"`
	State       string `json:"state" binding:"required,len=2"`
	Street1     string `json:"street1" binding:"required"`
	Street2     string `json:"street2"`
	ZipCode     string `json:"zip_code" binding:"required,len=5,numeric"`
}
