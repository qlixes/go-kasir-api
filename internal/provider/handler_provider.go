package provider

import "kasir-api/internal/handler"

type HandlerProvider struct {
	CategoryHandler *handler.CategoryHandler
	MainHandler     *handler.MainHandler
	ProductHandler  *handler.ProductHandler
}

func NewHandler(service *ServiceProvider) *HandlerProvider {
	categoryHandler := handler.NewCategoryHandler(service.CategoryService)
	mainHandler := handler.NewMainHandler(service.MainService)
	productHandler := handler.NewProductHandler(service.ProductService, service.CategoryService)

	return &HandlerProvider{
		CategoryHandler: categoryHandler,
		MainHandler:     mainHandler,
		ProductHandler:  productHandler,
	}
}
