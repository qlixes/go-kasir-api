package models

type Category struct {
	BaseModel
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
}
