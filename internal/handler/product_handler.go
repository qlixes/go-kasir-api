package handler

import (
	"encoding/json"
	"net/http"

	"kasir-api/internal/model"
	"kasir-api/internal/service"
	"kasir-api/internal/util"

	"github.com/google/uuid"
)

type ProductHandler struct {
	product  service.ProductService
	category service.CategoryService
}

func NewProductHandler(product service.ProductService, category service.CategoryService) *ProductHandler {
	return &ProductHandler{
		product:  product,
		category: category,
	}
}

func (s *ProductHandler) GetProductIndex(w http.ResponseWriter, r *http.Request) {
	var items []model.Product
	name := r.URL.Query().Get("name")
	active := r.URL.Query().Get("active")
	items, err := s.product.ShowProduct(name, &active)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}
	util.ResponseSuccess(w, 200, "Successfully", items)
}

func (s *ProductHandler) PostProductStore(w http.ResponseWriter, r *http.Request) {
	var body model.ProductDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	// check category_id
	categories, err := s.category.FindCategoryId(body.CategoryID.String())
	if err != nil {
		util.ResponseFail(w, 500, "Invalid Category ID")
		return
	}

	// parser
	item := model.Product{
		Name:     body.Name,
		Price:    body.Price,
		Quantity: body.Quantity,
		Status:   body.Status,
		Category: *categories,
	}

	items, err := s.product.StoreProduct(item)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	util.ResponseSuccess(w, 201, "Successfully created", items)
}

func (s *ProductHandler) PutProductId(w http.ResponseWriter, r *http.Request) {
	var body model.ProductDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	item, err := s.product.FindProductId(r.PathValue("id"))
	if err != nil {
		util.ResponseFail(w, 500, "Invalid Product ID")
		return
	}

	if body.Name != "" {
		item.Name = body.Name
	}

	if body.Price != 0 {
		item.Price = body.Price
	}

	if body.Quantity != 0 {
		item.Quantity = body.Quantity
	}

	if body.Status != false {
		item.Status = body.Status
	}

	if body.CategoryID != uuid.Nil {
		item.Category.ID = body.CategoryID
	}

	categories, err := s.category.FindCategoryId(item.Category.ID.String())
	if err != nil {
		util.ResponseFail(w, 500, "Invalid Category ID")
		return
	}

	item.Category = *categories

	data, err := s.product.EditProduct(r.PathValue("id"), *item)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	util.ResponseSuccess(w, 200, "Successfully updated", data)
}

func (s *ProductHandler) GetProductId(w http.ResponseWriter, r *http.Request) {
	var item *model.Product
	item, err := s.product.FindProductId(r.PathValue("id"))
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	util.ResponseSuccess(w, 200, "Successfully", item)
}

func (s *ProductHandler) DeleteProductId(w http.ResponseWriter, r *http.Request) {
	err := s.product.EraseProduct(r.PathValue("id"))
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	util.ResponseSuccess(w, 200, "Successfully deleted", nil)
}
