package handlers

import (
	"chaos-io/microserver/product-api/data"
	"log"
	"net/http"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and satisfies the http.Handler interface
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	// catch all
	// if no method satisfied return an error
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	pl := data.GetProducts()

	// serialize the list to JSON
	err := pl.ToJSON(w)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}
}
