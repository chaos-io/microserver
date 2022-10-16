package main

import (
	"fmt"
	"testing"

	"chaos-io/microserver/product-api/sdk/client"
	"chaos-io/microserver/product-api/sdk/client/products"
)

func TestClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prods, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(prods)
}
