package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:        "apple",
		Description: "123",
		SKU:         "sku-saf-kj",
	}

	err := NewValidation().Validate(p)
	if err != nil {
		t.Error(err)
	}
}
