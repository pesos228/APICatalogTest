package handler

import (
	"api/test/catalog/internal/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type createProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type updateProductRequest struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type productHandler struct {
	service service.ProductService
}

func (h *productHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.service.FindAll())
}

func (h *productHandler) SaveProduct(w http.ResponseWriter, r *http.Request) {
	var req createProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := h.service.Save(req.Name, req.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("new product with id:" + id))
}

func (h *productHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")
	if id == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product, err := h.service.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		http.Error(w, "Error JSON coder", http.StatusInternalServerError)
		return
	}
}

func (h *productHandler) DeleteProductById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("product with id " + id + " deleted"))
}

func (h *productHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var req updateProductRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.Update(id, req.Name, req.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("product success updated with id " + id))
}

func NewProductHandler(serv service.ProductService) *productHandler {
	return &productHandler{
		service: serv,
	}
}
