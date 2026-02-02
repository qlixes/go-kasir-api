package handler

import (
	"encoding/json"
	"net/http"

	"kasir-api/internal/model"
	"kasir-api/internal/service"
	"kasir-api/internal/util"
)

type ProductHandler struct {
	service service.ProductService
}

var (
	items []model.Product
	item  model.Product
)

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (s *ProductHandler) GetProductIndex(w http.ResponseWriter, r *http.Request) {
	items, err := s.service.ShowProduct()
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}
	util.ResponseSuccess(w, 200, "Successfully", items)
}

func (s *ProductHandler) PostProductStore(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	items, err := s.service.StoreProduct(item)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}
	util.ResponseSuccess(w, 201, "Successfully created", items)
}

func (s *ProductHandler) PutProductId(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	data, err := s.service.EditProduct(r.PathValue("id"), item)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	util.ResponseSuccess(w, 200, "Successfully updated", data)
}

func (s *ProductHandler) GetProductId(w http.ResponseWriter, r *http.Request) {
	item, err := s.service.FindProductId(r.PathValue("id"))
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	util.ResponseSuccess(w, 200, "Successfully", item)
}

func (s *ProductHandler) DeleteProductId(w http.ResponseWriter, r *http.Request) {
	err := s.service.EraseProduct(r.PathValue("id"))
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	util.ResponseSuccess(w, 200, "Successfully deleted", nil)
}
