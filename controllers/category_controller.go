package controllers

import (
	"encoding/json"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
)

type CategoryController struct {
	svc services.CategoryService
}

func NewCategoryController(svc services.CategoryService) *CategoryController {
	return &CategoryController{svc: svc}
}

func (c *CategoryController) CategoryIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := c.svc.ShowCategory()

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

func (c *CategoryController) CategoryStore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body models.Category
	err := json.NewDecoder(r.Body).Decode(&body)

	data, err := c.svc.StoreCategory(body)

	if err != nil {
		format := map[string]string{
			"status":  "Failed",
			"message": "Failed load data",
		}
		json.NewEncoder(w).Encode(format)
	} else {
		format := map[string]any{
			"status":  "Success",
			"message": "Success store data",
			"data":    data,
		}
		json.NewEncoder(w).Encode(format)
	}
}

func (c *CategoryController) CategoryEdit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categoryId := r.PathValue("id")

	var body models.Category
	err := json.NewDecoder(r.Body).Decode(&body)

	data, err := c.svc.EditCategory(categoryId, &body)

	if err != nil {
		format := map[string]string{
			"status":  "Failed",
			"message": "Failed load data",
		}
		json.NewEncoder(w).Encode(format)
	} else {
		format := map[string]any{
			"status":  "Success",
			"message": "Success update data",
			"data":    data,
		}
		json.NewEncoder(w).Encode(format)
	}
}

func (c *CategoryController) CategoryFindId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categoryId := r.PathValue("id")

	data, err := c.svc.FindCategoryId(categoryId)

	if err != nil {
		format := map[string]string{
			"status":  "Failed",
			"message": "Failed load data",
		}
		json.NewEncoder(w).Encode(format)
	} else {
		format := map[string]any{
			"status":  "Success",
			"message": "Success show category",
			"data":    data,
		}
		json.NewEncoder(w).Encode(format)
	}
}

func (c *CategoryController) CategoryRemove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	categoryId := r.PathValue("id")

	err := c.svc.EraseCategory(categoryId)

	if err != nil {
		format := map[string]string{
			"status":  "Failed",
			"message": "Failed load data",
		}
		json.NewEncoder(w).Encode(format)
	} else {
		format := map[string]any{
			"status":  "Success",
			"message": "Success remove category",
		}
		json.NewEncoder(w).Encode(format)
	}
}
