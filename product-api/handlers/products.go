// Package classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"log"
	"net/http"
	"strconv"

	"chaos-io/microserver/product-api/data"

	"github.com/gorilla/mux"
)

type KeyProduct struct{}

// Products is a http.Handler
type Products struct {
	l *log.Logger
	v *data.Validation
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// swagger:route GET /products products listProducts
// Returns a list of products
// Produces:
// - application/json
// responses:
// 	200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	w.Header().Add("Content-Type", "application/json")
	// fetch the products from the datastore
	pl := data.GetProducts()

	// serialize the list to JSON
	err := data.ToJSON(pl, w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products getProduct
// get a product details
//
// responses:
// 	201: noContentResponse
//	404: errorResponse
func (p *Products) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	p.l.Println("Handle GET Product")

	// fetch the product from the datastore
	pl, err := data.GetProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "product not found", http.StatusInternalServerError)
		return
	}

	// serialize the list to JSON
	err = data.ToJSON(pl, w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route POST /products products addProduct
// add a product
//
// responses:
// 	200: productsResponse
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

// swagger:route PUT /products products updateProduct
//
// Consumes:
// 	- application/json
// responses:
// 	200: productsResponse
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle PUT Product", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "product not found", http.StatusInternalServerError)
		return
	}
}

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
// 	201: noContentResponse
//	404: errorResponse
// DeleteProduct deletes a product from tht database
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle Delete Product", id)
	err := data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "product not found", http.StatusInternalServerError)
		return
	}
}
