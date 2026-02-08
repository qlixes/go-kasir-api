package provider

import (
	"kasir-api/internal/service"
)

type ServiceProvider struct {
	CategoryService service.CategoryService
	MainService     service.MainService
	ProductService  service.ProductService
}

func NewService(provider *RepositoryProvider) *ServiceProvider {
	categoryService := service.NewCategoryService(provider.categoryRepo)
	mainService := service.NewMainService()
	productService := service.NewProductService(provider.productRepo)

	return &ServiceProvider{
		CategoryService: categoryService,
		MainService:     mainService,
		ProductService:  productService,
	}
}
