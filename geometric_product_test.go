package edche

import (
	"strings"
	"testing"
)

// Tests if the geometric product between two multivectors (plaintext and key) are calculated properly.
func TestGeometricProduct(t *testing.T) {
	m0 := []string{"40", "-29", "29", "-25", "33", "-29", "32", "-28"}
	s0 := []string{"113", "-84", "89", "-65", "131", "-107", "117", "-93"}

	m1 := New(m0)
	s1 := New(s0)

	m := GeometricProduct(m1, s1)

	if strings.Compare(m.E0.RatString(), "-2612") != 0 ||
		strings.Compare(m.E1.RatString(), "-2037") != 0 ||
		strings.Compare(m.E2.RatString(), "12348") != 0 ||
		strings.Compare(m.E3.RatString(), "2524") != 0 ||
		strings.Compare(m.E12.RatString(), "12938") != 0 ||
		strings.Compare(m.E13.RatString(), "-2914") != 0 ||
		strings.Compare(m.E23.RatString(), "13417") != 0 ||
		strings.Compare(m.E123.RatString(), "-12701") != 0 {
		t.Errorf("Wrong results for the geometric product calculation.")
	}
}
