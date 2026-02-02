package service

import (
	"kasir-api/internal/model"
	"kasir-api/internal/repository"
)

type ProductService interface {
	ShowProduct() ([]model.Product, error)
	FindProductId(id string) (*model.Product, error)
	EraseProduct(id string) error
	EditProduct(id string, Product model.Product) (*model.Product, error)
	StoreProduct(Product model.Product) (*model.Product, error)
}

type ProductService struct {
	ProductRepo repository.ProductRepository
}

func NewProductService(ProductRepo repository.ProductRepository) ProductService {
	return &ProductService{
		ProductRepo: ProductRepo,
	}
}

func (s *ProductService) ShowProduct() ([]model.Product, error) {
	return s.ProductRepo.FindAll()
}

func (s *ProductService) FindProductId(id string) (*model.Product, error) {
	return s.ProductRepo.FindId(id)
}

func (s *ProductService) EraseProduct(id string) error {
	return s.ProductRepo.Erase(id)
}

func (s *ProductService) EditProduct(id string, Product model.Product) (*model.Product, error) {
	_, err := s.ProductRepo.FindId(id)
	if err != nil {
		return nil, err
	}

	return s.ProductRepo.Edit(id, &Product)
}

func (s *ProductService) StoreProduct(Product model.Product) (*model.Product, error) {
	_, err := s.ProductRepo.FindName(Product.Name)
	if err != nil {
		return nil, err
	}

	return s.ProductRepo.Store(&Product)
}
