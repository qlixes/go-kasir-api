package handler

import (
	"encoding/json"
	"net/http"

	"kasir-api/internal/model"
	"kasir-api/internal/service"
	"kasir-api/internal/util"
)

type CategoryHandler struct {
	category service.CategoryService
}

func NewCategoryHandler(category service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		category: category,
	}
}

func (s *CategoryHandler) GetCategoryIndex(w http.ResponseWriter, r *http.Request) {
	var items []model.Category
	items, err := s.category.ShowCategory()
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}
	util.ResponseSuccess(w, 200, "Successfully", items)
}

func (s *CategoryHandler) PostCategoryStore(w http.ResponseWriter, r *http.Request) {
	var item model.Category
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	items, err := s.category.StoreCategory(item)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}
	util.ResponseSuccess(w, 201, "Successfully created", items)
}

func (s *CategoryHandler) PutCategoryId(w http.ResponseWriter, r *http.Request) {
	var item model.Category
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	data, err := s.category.EditCategory(r.PathValue("id"), item)
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	util.ResponseSuccess(w, 200, "Successfully updated", data)
}

func (s *CategoryHandler) GetCategoryId(w http.ResponseWriter, r *http.Request) {
	var item *model.Category
	item, err := s.category.FindCategoryId(r.PathValue("id"))
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	util.ResponseSuccess(w, 200, "Successfully", item)
}

func (s *CategoryHandler) DeleteCategoryId(w http.ResponseWriter, r *http.Request) {
	err := s.category.EraseCategory(r.PathValue("id"))
	if err != nil {
		util.ResponseFail(w, 500, err.Error())
		return
	}

	util.ResponseSuccess(w, 200, "Successfully deleted", nil)
}
