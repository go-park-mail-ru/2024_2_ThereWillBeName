package http

import (
	"2024_2_ThereWillBeName/internal/models"
	"2024_2_ThereWillBeName/internal/pkg/categories"
	httpresponse "2024_2_ThereWillBeName/internal/pkg/httpresponses"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

type CategoryHandler struct {
	ch categories.CategoryUsecase
}

func NewCategoryHandler(ch categories.CategoryUsecase) *CategoryHandler { return &CategoryHandler{ch} }

func (h *CategoryHandler) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	categories, err := h.ch.GetCategories(r.Context(), requestData.Limit, requestData.Offset)
	if err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf("Couldn't get list of categories: %v", err)
		return
	}
	logger.Printf("Got list of categories: %v", categories)
	httpresponse.SendJSONResponse(w, categories, http.StatusOK)
}

func (h *CategoryHandler) GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	category, err := h.ch.GetCategory(r.Context(), requestData.ID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			httpresponse.SendJSONResponse(w, nil, http.StatusNotFound)
			logger.Println(err.Error())
			return
		}
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf("Couldn't get category: %v", err)
		return
	}
	httpresponse.SendJSONResponse(w, category, http.StatusOK)
}

func (h *CategoryHandler) PostCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	logger.Println(category)
	if err := h.ch.CreateCategory(r.Context(), category); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf("Couldn't create category: %v", err)
		return
	}
	httpresponse.SendJSONResponse(w, "category successfully created", http.StatusCreated)
}

func (h *CategoryHandler) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		ID int `json:"id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	if err := h.ch.DeleteCategory(r.Context(), requestData.ID); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			httpresponse.SendJSONResponse(w, nil, http.StatusNotFound)
			logger.Println(err.Error())
			return
		}
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf("Couldn't delete category: %v", err)
		return
	}
	httpresponse.SendJSONResponse(w, "category successfully deleted", http.StatusNoContent)
}

func (h *CategoryHandler) PutCategoryHandler(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		httpresponse.SendJSONResponse(w, nil, http.StatusBadRequest)
		logger.Println(err.Error())
		return
	}
	if err := h.ch.UpdateCategory(r.Context(), category); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			httpresponse.SendJSONResponse(w, nil, http.StatusNotFound)
			logger.Println(err.Error())
			return
		}
		httpresponse.SendJSONResponse(w, nil, http.StatusInternalServerError)
		logger.Printf("Couldn't update category: %v", err)
		return
	}
	httpresponse.SendJSONResponse(w, "category successfully updated", http.StatusOK)
}
