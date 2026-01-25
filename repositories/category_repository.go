package repositories

import (
	"kasir-api/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]models.Category, error)
	FindId(id string) (*models.Category, error)
	FindName(name string) (*models.Category, error)
	Erase(id string) error
	Edit(id string, payload *models.Category) (models.Category, error)
	Store(payload models.Category) (*models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) FindAll() ([]models.Category, error) {
	var categories []models.Category

	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) FindId(id string) (*models.Category, error) {
	var category models.Category

	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *categoryRepository) Erase(id string) error {
	var category models.Category

	err := r.db.Find(&category, id).Error
	return err
}

func (r *categoryRepository) Edit(id string, payload *models.Category) (models.Category, error) {
	var category models.Category

	err := r.db.Find(&category, id).Updates(payload).Error
	return category, err
}

func (r *categoryRepository) Store(payload models.Category) (*models.Category, error) {
	result := r.db.Create(&payload)
	if result.Error != nil {
		return nil, result.Error
	}

	return &payload, nil
}

func (r *categoryRepository) FindName(name string) (*models.Category, error) {
	var category models.Category

	err := r.db.Where("name = ?", name).First(&category).Error
	return &category, err
}
