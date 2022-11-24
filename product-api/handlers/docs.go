// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "github.com/chaos-io/microserver/product-api/data"

// A list of products return in the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContentResponse
type noContentResponse struct {
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// swagger:parameters deleteProduct getProduct
type productIDParametersWrapper struct {
	// The id of the product ot delete from the database
	// in:path
	// require: true
	ID int `json:"id"`
}
