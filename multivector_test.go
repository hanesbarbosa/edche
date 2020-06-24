package edche

import (
	"strings"
	"testing"
)

// Tests if the function is returning a multivector with rational coefficients.
func TestNew(t *testing.T) {
	n := []string{"12", "23", "34", "45", "56", "67", "78", "89"}
	m := New(n)
	// m.e0, m.e1, m.e2, m.e3, m.e12, m.e13, m.e23, m.e123
	if strings.Compare(m.E0.RatString(), n[0]) != 0 ||
		strings.Compare(m.E1.RatString(), n[1]) != 0 ||
		strings.Compare(m.E2.RatString(), n[2]) != 0 ||
		strings.Compare(m.E3.RatString(), n[3]) != 0 ||
		strings.Compare(m.E12.RatString(), n[4]) != 0 ||
		strings.Compare(m.E13.RatString(), n[5]) != 0 ||
		strings.Compare(m.E23.RatString(), n[6]) != 0 ||
		strings.Compare(m.E123.RatString(), n[7]) != 0 {
		t.Errorf("Coefficient numerator is different from given string.")
	}
}

func TestToString(t *testing.T) {
	m := New([]string{"12", "23", "34", "45", "56", "67", "78", "89"})

	if strings.Compare(m.ToString(), "12e0+23e1+34e2+45e3+56e12+67e13+78e23+89e123") != 0 {
		t.Errorf("Wrong formatting for multivector.")
	}

	m = New([]string{"12", "-23", "34", "45", "56", "-67", "78", "89"})

	if strings.Compare(m.ToString(), "12e0+-23e1+34e2+45e3+56e12+-67e13+78e23+89e123") != 0 {
		t.Errorf("Wrong formatting for multivector.")
	}
}

func TestStringToMultivector(t *testing.T) {
	s := "230e0+-179e1+156e2+-110e3+238e12+-192e13+246e23+-200e123"
	m := StringToMultivector(s)

	if strings.Compare(m.E0.RatString(), "230") != 0 ||
		strings.Compare(m.E1.RatString(), "-179") != 0 ||
		strings.Compare(m.E2.RatString(), "156") != 0 ||
		strings.Compare(m.E3.RatString(), "-110") != 0 ||
		strings.Compare(m.E12.RatString(), "238") != 0 ||
		strings.Compare(m.E13.RatString(), "-192") != 0 ||
		strings.Compare(m.E23.RatString(), "246") != 0 ||
		strings.Compare(m.E123.RatString(), "-200") != 0 {
		t.Errorf("Coefficients from multivector are different from given string.")
	}
}
