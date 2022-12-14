package main

import (
	"fmt"
	"testing"

	"github.com/chaos-io/microserver/product-api/sdk/client"
	"github.com/chaos-io/microserver/product-api/sdk/client/products"
)

func TestClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prods, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", prods.GetPayload()[0])
	fmt.Println(prods)
}
