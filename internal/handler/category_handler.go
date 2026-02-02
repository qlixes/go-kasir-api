package handler

import (
	"encoding/json"
	"net/http"

	"kasir-api/internal/model"
	"kasir-api/internal/service"
	"kasir-api/internal/util"
)

type CategoryHandler struct {
	service service.CategoryService
}

var (
	items []model.Category
	item  model.Category
)

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (s *CategoryHandler) GetCategoryIndex(w http.ResponseWriter, r *http.Request) {
	items, err := s.service.ShowCategory()
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}
	util.ResponseSuccess(w, 200, "Successfully", items)
}

func (s *CategoryHandler) PostCategoryStore(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	items, err := s.service.StoreCategory(item)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}
	util.ResponseSuccess(w, 201, "Successfully created", items)
}

func (s *CategoryHandler) PutCategoryId(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	data, err := s.service.EditCategory(r.PathValue("id"), item)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	util.ResponseSuccess(w, 200, "Successfully updated", data)
}

func (s *CategoryHandler) GetCategoryId(w http.ResponseWriter, r *http.Request) {
	item, err := s.service.FindCategoryId(r.PathValue("id"))
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	util.ResponseSuccess(w, 200, "Successfully", item)
}

func (s *CategoryHandler) DeleteCategoryId(w http.ResponseWriter, r *http.Request) {
	err := s.service.EraseCategory(r.PathValue("id"))
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	util.ResponseSuccess(w, 200, "Successfully deleted", nil)
}
