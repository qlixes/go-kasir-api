package provider

import (
	"kasir-api/internal/service"
)

type ServiceProvider struct {
	CategoryService service.CategoryService
	MainService     service.MainService
}

func NewService(provider *RepositoryProvider) *ServiceProvider {
	categoryService := service.NewCategoryService(provider.categoryRepo)
	mainService := service.NewMainService()

	return &ServiceProvider{
		CategoryService: categoryService,
		MainService:     mainService,
	}
}
