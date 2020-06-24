package edche

import (
	"strings"
	"testing"
)

// Tests if a multivector can be transformed back into a number.
func TestMultivectorToNumber(t *testing.T) {
	c := []string{"31", "29", "29", "29", "29", "29", "29", "29"}

	m := New(c)
	n := MultivectorToNumber(m)

	if strings.Compare(n.String(), "234") != 0 {
		t.Errorf("Wrong conversion from multivector to number.")
	}
}
