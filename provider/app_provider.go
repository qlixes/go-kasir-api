package provider

import (
	"kasir-api/controllers"
)

type Container struct {
	CategoryController *controllers.CategoryController
	MainController     *controllers.MainController
}

func NewContainer() *Container {

}
