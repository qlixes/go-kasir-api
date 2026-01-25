package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoryService interface {
	ShowCategory() ([]models.Category, error)
	FindCategoryId(id string) (*models.Category, error)
	EraseCategory(id string) error
	EditCategory(id string, category *models.Category) (models.Category, error)
	StoreCategory(category models.Category) (*models.Category, error)
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) ShowCategory() ([]models.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) FindCategoryId(id string) (*models.Category, error) {
	return s.repo.FindId(id)
}

func (s *categoryService) EraseCategory(id string) error {
	return s.repo.Erase(id)
}

func (s *categoryService) EditCategory(id string, category *models.Category) (models.Category, error) {
	_, err := s.repo.FindId(id)

	if err != nil {
		return models.Category{}, err
	}

	return s.repo.Edit(id, category)
}

func (s *categoryService) StoreCategory(category models.Category) (*models.Category, error) {
	_, err := s.repo.FindName(category.Name)

	if err != nil {
		return nil, err
	}

	return s.repo.Store(category)
}
