package provider

import "kasir-api/internal/handler"

type HandlerProvider struct {
	CategoryHandler *handler.CategoryHandler
	MainHandler     *handler.MainHandler
}

func NewHandler(service *ServiceProvider) *HandlerProvider {
	categoryHandler := handler.NewCategoryHandler(service.CategoryService)
	mainHandler := handler.NewMainHandler(service.MainService)

	return &HandlerProvider{
		CategoryHandler: categoryHandler,
		MainHandler:     mainHandler,
	}
}
