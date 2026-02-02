package repository

import "kasir-api/internal/model"

type ProductRepository interface {
	FindAll() ([]model.Category, error)
	FindId(id string) (*model.Category, error)
	FindName(name string) (*model.Category, error)
	Erase(id string) error
	Edit(id string, payload *model.Category) (model.Category, error)
	Store(payload model.Category) (*model.Category, error)
}
