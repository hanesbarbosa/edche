package edche

import (
	"strings"
	"testing"
)

// Tests if the scalar multiplication is calculated properly.
func TestAmplitudeSquared(t *testing.T) {
	c := []string{"40", "-29", "29", "-25", "33", "-29", "32", "-28"}

	m := New(c)
	m = AmplitudeSquared(m)

	if strings.Compare(m.E0.RatString(), "1463") != 0 ||
		strings.Compare(m.E1.RatString(), "0") != 0 ||
		strings.Compare(m.E2.RatString(), "0") != 0 ||
		strings.Compare(m.E3.RatString(), "0") != 0 ||
		strings.Compare(m.E12.RatString(), "0") != 0 ||
		strings.Compare(m.E13.RatString(), "0") != 0 ||
		strings.Compare(m.E23.RatString(), "0") != 0 ||
		strings.Compare(m.E123.RatString(), "-416") != 0 {
		t.Errorf("Wrong results for the amplitude squared calculation.")
	}
}
