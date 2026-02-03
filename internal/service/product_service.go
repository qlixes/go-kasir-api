package service

import (
	"database/sql"
	"errors"
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

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (s *productService) ShowProduct() ([]model.Product, error) {
	return s.productRepo.FindAll()
}

func (s *productService) FindProductId(id string) (*model.Product, error) {
	return s.productRepo.FindId(id)
}

func (s *productService) EraseProduct(id string) error {
	return s.productRepo.Erase(id)
}

func (s *productService) EditProduct(id string, Product model.Product) (*model.Product, error) {
	_, err := s.productRepo.FindId(id)
	if err != nil {
		return nil, err
	}

	return s.productRepo.Edit(id, &Product)
}

func (s *productService) StoreProduct(Product model.Product) (*model.Product, error) {
	_, err := s.productRepo.FindName(Product.Name)
	if err == sql.ErrNoRows {
		return s.productRepo.Store(&Product)
	}

	return nil, errors.New("Duplicate record")
}
