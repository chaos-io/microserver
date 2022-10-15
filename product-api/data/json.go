package data

import (
	"encoding/json"
	"io"
)

// ToJSON serializers the given interface into a string based JSON fromat
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON deserializers the object from JSON string
// in an io.Reader to the given interface
func FromJSON(i interface{}, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(i)
}
