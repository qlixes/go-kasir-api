package service

import (
	"database/sql"
	"errors"
	"kasir-api/internal/model"
	"kasir-api/internal/repository"
)

type CategoryService interface {
	ShowCategory() ([]model.Category, error)
	FindCategoryId(id string) (*model.Category, error)
	EraseCategory(id string) error
	EditCategory(id string, category model.Category) (*model.Category, error)
	StoreCategory(category model.Category) (*model.Category, error)
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (s *categoryService) ShowCategory() ([]model.Category, error) {
	return s.categoryRepo.FindAll()
}

func (s *categoryService) FindCategoryId(id string) (*model.Category, error) {
	return s.categoryRepo.FindId(id)
}

func (s *categoryService) EraseCategory(id string) error {
	return s.categoryRepo.Erase(id)
}

func (s *categoryService) EditCategory(id string, category model.Category) (*model.Category, error) {
	_, err := s.categoryRepo.FindId(id)
	if err != nil {
		return nil, err
	}

	return s.categoryRepo.Edit(id, &category)
}

func (s *categoryService) StoreCategory(category model.Category) (*model.Category, error) {
	_, err := s.categoryRepo.FindName(category.Name)
	if err == sql.ErrNoRows {
		return s.categoryRepo.Store(&category)
	}

	return nil, errors.New("Duplicate record")
}
