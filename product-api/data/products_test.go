package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:        "apple",
		Description: "123",
		SKU:         "sku-saf-kj",
	}

	err := p.Validate()
	if err != nil {
		t.Error(err)
	}
}
