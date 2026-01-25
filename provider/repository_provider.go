package provider

import (
	"kasir-api/configs"
	"kasir-api/models"
	"kasir-api/repositories"
)

type RepositoryProvider struct {
	CategoryRepo *repositories.CategoryRepository
}

func NewRepository() *RepositoryProvider {
	db := configs.LoadDB()
	db.AutoMigrate(&models.Category, &models.Product)

	return &RepositoryProvider{
		CategoryRepo: repositories.NewCategoryRepository(db),
	}
}
