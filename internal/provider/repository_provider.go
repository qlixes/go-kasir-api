package provider

import (
	"database/sql"

	"kasir-api/internal/repository"
)

type RepositoryProvider struct {
	categoryRepo repository.CategoryRepository
}

func NewRepository(db *sql.DB) *RepositoryProvider {
	categoryRepo := repository.NewCategoryRepository(db)

	return &RepositoryProvider{
		categoryRepo: categoryRepo,
	}
}
