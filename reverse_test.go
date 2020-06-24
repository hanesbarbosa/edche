package edche

import (
	"strings"
	"testing"
)

// Tests if the reverse is calculated properly.
func TestReverse(t *testing.T) {
	c := []string{"40", "-29", "29", "-25", "33", "-29", "32", "-28"}

	m := New(c)
	m = Reverse(m)

	// 40e0 + 29e1 + -29e2 + 25e3 + 33e12 + -29e13 + 32e23 + 28e123
	if strings.Compare(m.E0.RatString(), "40") != 0 ||
		strings.Compare(m.E1.RatString(), "29") != 0 ||
		strings.Compare(m.E2.RatString(), "-29") != 0 ||
		strings.Compare(m.E3.RatString(), "25") != 0 ||
		strings.Compare(m.E12.RatString(), "33") != 0 ||
		strings.Compare(m.E13.RatString(), "-29") != 0 ||
		strings.Compare(m.E23.RatString(), "32") != 0 ||
		strings.Compare(m.E123.RatString(), "28") != 0 {
		t.Errorf("Wrong results for the reverse calculation.")
	}
}
