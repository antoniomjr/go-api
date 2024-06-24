package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/antoniomjr/go/9-apis/internal/dto"
	"github.com/antoniomjr/go/9-apis/internal/entity"
	"github.com/antoniomjr/go/9-apis/internal/infra/database"
	entityPkg "github.com/antoniomjr/go/9-apis/pkg/entity"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}
// Create Product godoc
// @Summary 		Create product
// @Description 	Create products
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			request body dto.CreateProductInput	true "product request"
// @Success 		201 
// @Failure 		500	{object}	Error
// @Router 			/products [post]
// @Security 		ApiKeyAuth
func (ph *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = ph.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// ListAccounts godoc
// @Summary 		List product
// @Description 	get all products
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			page query string false "page number"
// @Param 			limit query string false "limit"
// @Param 			sort query string false "sort"
// @Success 		200 {array} entity.Product 
// @Failure 		404	{object}	Error
// @Failure 		500	{object}	Error
// @Router 			/products [get]
// @Security 		ApiKeyAuth
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// GetProduct godoc
// @Summary 		Get a product
// @Description 	get a products
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "product ID" Format(uuid)
// @Success 		200 {object} entity.Product 
// @Failure 		404	
// @Failure 		500	{object}	Error
// @Router 			/products/{id} [get]
// @Security 		ApiKeyAuth
func (h *ProductHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct godoc
// @Summary 		Update a product
// @Description 	Update a product
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "product ID" Format(uuid)
// @Param 			request body dto.CreateProductInput true "product request"
// @Success 		200 
// @Failure 		404	
// @Failure 		500	{object}	Error
// @Router 			/products/{id} [put]
// @Security 		ApiKeyAuth
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteProduct godoc
// @Summary 		Delete a product
// @Description 	Delete a products
// @Tags 			products
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "product ID" Format(uuid)
// @Success 		200  
// @Failure 		404	
// @Failure 		500	{object}	Error
// @Router 			/products/{id} [delete]
// @Security 		ApiKeyAuth
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err := h.ProductDB.FindById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}