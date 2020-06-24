package edche

import (
	"strings"
	"testing"
)

// Transforms a number in string format into a multivector.
// This is the simple transformation without any dynamic packing or improvements.
func TestNumberToMultivector(t *testing.T) {
	n := "23"
	m := NumberToMultivector(n)

	if strings.Compare(m.E0.RatString(), "9") != 0 ||
		strings.Compare(m.E1.RatString(), "2") != 0 ||
		strings.Compare(m.E2.RatString(), "2") != 0 ||
		strings.Compare(m.E3.RatString(), "2") != 0 ||
		strings.Compare(m.E12.RatString(), "2") != 0 ||
		strings.Compare(m.E13.RatString(), "2") != 0 ||
		strings.Compare(m.E23.RatString(), "2") != 0 ||
		strings.Compare(m.E123.RatString(), "2") != 0 {
		t.Errorf("Wrong conversion from number to multivector.")
	}
}
