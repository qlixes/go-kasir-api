package models

type Product struct {
	BaseModel
	Name     string `json:"name" gorm:"unique"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}
