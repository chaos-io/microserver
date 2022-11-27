package data

import (
	"fmt"
	"testing"

	"github.com/hashicorp/go-hclog"
)

func TestNewRates(t *testing.T) {
	rates, err := NewRates(hclog.Default())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("rates %#v\n", rates.rates)
}
