package edche

import (
	"strings"
	"testing"
)

// Tests if the scalar division is calculated properly.
func TestScalarDivision(t *testing.T) {
	c := []string{"40", "-29", "29", "-25", "33", "-29", "32", "-28"}

	m := New(c)
	m = ScalarDivision(m, "2")

	if strings.Compare(m.E0.RatString(), "20") != 0 ||
		strings.Compare(m.E1.RatString(), "-29/2") != 0 ||
		strings.Compare(m.E2.RatString(), "29/2") != 0 ||
		strings.Compare(m.E3.RatString(), "-25/2") != 0 ||
		strings.Compare(m.E12.RatString(), "33/2") != 0 ||
		strings.Compare(m.E13.RatString(), "-29/2") != 0 ||
		strings.Compare(m.E23.RatString(), "16") != 0 ||
		strings.Compare(m.E123.RatString(), "-14") != 0 {
		t.Errorf("Wrong results for the rational form of coefficients.")
	}
}
