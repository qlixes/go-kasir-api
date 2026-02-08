package provider

import (
	"database/sql"

	"kasir-api/internal/repository"
)

type RepositoryProvider struct {
	categoryRepo repository.CategoryRepository
	productRepo  repository.ProductRepository
}

func NewRepository(db *sql.DB) *RepositoryProvider {
	categoryRepo := repository.NewCategoryRepository(db)
	productRepo := repository.NewProductRepository(db)

	return &RepositoryProvider{
		categoryRepo: categoryRepo,
		productRepo:  productRepo,
	}
}
