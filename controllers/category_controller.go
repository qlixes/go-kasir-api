package controllers

import (
	"encoding/json"
	"kasir-api/services"
	"net/http"
)

type CategoryController struct {
	svc services.CategoryService
}

func NewCategoryController(svc services.CategoryService) {
	return &CategoryController{svc: svc}
}

func (c *CategoryController) CategoryIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := c.categoryService.ShowCategory()

	if err != nil {
		format := map[string]string{
			"status":  "Failed",
			"message": "Failed load data",
		}
		json.NewEncoder(w).Encode(format)
	} else {
		format := map[string]any{
			"status":  "Success",
			"message": "Success load data",
			"data":    data,
		}
		json.NewEncoder(w).Encode(format)
	}
}

func CategoryStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := controller.categoryService.ShowCategory()

	if err != nil {
		format := map[string]string{
			"status":  "Failed",
			"message": "Failed load data",
		}
		json.NewEncoder(w).Encode(format)
	} else {
		format := map[string]any{
			"status":  "Success",
			"message": "Success load data",
			"data":    data,
		}
		json.NewEncoder(w).Encode(format)
	}
}

func CategoryEdit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := controller.categoryService.ShowCategory()

	if err != nil {
		format := map[string]string{
			"status":  "Failed",
			"message": "Failed load data",
		}
		json.NewEncoder(w).Encode(format)
	} else {
		format := map[string]any{
			"status":  "Success",
			"message": "Success load data",
			"data":    data,
		}
		json.NewEncoder(w).Encode(format)
	}
}

func CategoryFind(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := controller.categoryService.ShowCategory()

	if err != nil {
		format := map[string]string{
			"status":  "Failed",
			"message": "Failed load data",
		}
		json.NewEncoder(w).Encode(format)
	} else {
		format := map[string]any{
			"status":  "Success",
			"message": "Success load data",
			"data":    data,
		}
		json.NewEncoder(w).Encode(format)
	}
}

func CategoryRemove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := controller.categoryService.ShowCategory()

	if err != nil {
		format := map[string]string{
			"status":  "Failed",
			"message": "Failed load data",
		}
		json.NewEncoder(w).Encode(format)
	} else {
		format := map[string]any{
			"status":  "Success",
			"message": "Success load data",
			"data":    data,
		}
		json.NewEncoder(w).Encode(format)
	}
}
