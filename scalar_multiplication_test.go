package edche

import (
	"math/big"
	"strings"
	"testing"
)

// Tests if the scalar multiplication is calculated properly.
func TestScalarMultiplication(t *testing.T) {
	c := []string{"40", "-29", "29", "-25", "33", "-29", "32", "-28"}

	m := New(c)

	n := new(big.Rat)
	n.SetString("2/1")

	m = ScalarMultiplication(m, n)

	if strings.Compare(m.E0.RatString(), "80") != 0 ||
		strings.Compare(m.E1.RatString(), "-58") != 0 ||
		strings.Compare(m.E2.RatString(), "58") != 0 ||
		strings.Compare(m.E3.RatString(), "-50") != 0 ||
		strings.Compare(m.E12.RatString(), "66") != 0 ||
		strings.Compare(m.E13.RatString(), "-58") != 0 ||
		strings.Compare(m.E23.RatString(), "64") != 0 ||
		strings.Compare(m.E123.RatString(), "-56") != 0 {
		t.Errorf("Wrong results for the scalar multiplication calculation.")
	}
}
