package edche

import (
	"strings"
	"testing"
)

// Tests Inverse function
func TestInverse(t *testing.T) {
	c := []string{"40", "-29", "29", "-25", "33", "-29", "32", "-28"}

	m := New(c)
	m = Inverse(m)

	if strings.Compare(m.E0.RatString(), "70168/2313425") != 0 ||
		strings.Compare(m.E1.RatString(), "55739/2313425") != 0 ||
		strings.Compare(m.E2.RatString(), "-30363/2313425") != 0 ||
		strings.Compare(m.E3.RatString(), "50303/2313425") != 0 ||
		strings.Compare(m.E12.RatString(), "-37879/2313425") != 0 ||
		strings.Compare(m.E13.RatString(), "54491/2313425") != 0 ||
		strings.Compare(m.E23.RatString(), "-34752/2313425") != 0 ||
		strings.Compare(m.E123.RatString(), "-24324/2313425") != 0 {
		t.Errorf("Wrong results for inverse.")
	}
}

// Tests auxiliary function Numerator.
func TestNumerator(t *testing.T) {
	c := []string{"40", "-29", "29", "-25", "33", "-29", "32", "-28"}

	m := New(c)
	m = Numerator(m)

	// 70168e0 + 55739e1 + -30363e2 + 50303e3 + -37879e12 + 54491e13 + -34752e23 + -24324e123
	if strings.Compare(m.E0.RatString(), "70168") != 0 ||
		strings.Compare(m.E1.RatString(), "55739") != 0 ||
		strings.Compare(m.E2.RatString(), "-30363") != 0 ||
		strings.Compare(m.E3.RatString(), "50303") != 0 ||
		strings.Compare(m.E12.RatString(), "-37879") != 0 ||
		strings.Compare(m.E13.RatString(), "54491") != 0 ||
		strings.Compare(m.E23.RatString(), "-34752") != 0 ||
		strings.Compare(m.E123.RatString(), "-24324") != 0 {
		t.Errorf("Wrong results for numerator.")
	}
}

// Tests auxiliary function Denominator.
func TestDenominator(t *testing.T) {
	c := []string{"40", "-29", "29", "-25", "33", "-29", "32", "-28"}

	m := New(c)
	e0 := Denominator(m)

	if strings.Compare(e0, "2313425") != 0 {
		t.Errorf("Wrong result for coefficient.")
	}
}
