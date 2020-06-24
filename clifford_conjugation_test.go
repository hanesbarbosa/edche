package edche

import (
	"strings"
	"testing"
)

// Tests if the clifford conjugation is calculated properly.
func TestCliffordConjugation(t *testing.T) {
	c := []string{"40", "-29", "29", "-25", "33", "-29", "32", "-28"}

	m := New(c)
	m = CliffordConjugation(m)

	if strings.Compare(m.E0.RatString(), "40") != 0 ||
		strings.Compare(m.E1.RatString(), "29") != 0 ||
		strings.Compare(m.E2.RatString(), "-29") != 0 ||
		strings.Compare(m.E3.RatString(), "25") != 0 ||
		strings.Compare(m.E12.RatString(), "-33") != 0 ||
		strings.Compare(m.E13.RatString(), "29") != 0 ||
		strings.Compare(m.E23.RatString(), "-32") != 0 ||
		strings.Compare(m.E123.RatString(), "-28") != 0 {
		t.Errorf("Wrong results for the clifford conjugation calculation.")
	}
}
