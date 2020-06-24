package edche

import (
	"strings"
	"testing"
)

// Tests the addition of two multivectors.
func TestAddition(t *testing.T) {
	m1 := NumberToMultivector("5")
	m2 := NumberToMultivector("2")

	ms := Addition(m1, m2)
	ns := MultivectorToNumber(ms)

	if strings.Compare(ns.String(), "7") != 0 {
		t.Errorf("The addition is not correct!")
	}
}
