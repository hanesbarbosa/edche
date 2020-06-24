package edche

import (
	"strings"
	"testing"
)

func TestMean(t *testing.T) {
	m1 := NumberToMultivector("5")
	m2 := NumberToMultivector("4")
	ms := []Multivector{m1, m2}
	m := Mean(ms)

	if strings.Compare(m.ToString(), "9/2e0+0e1+0e2+0e3+0e12+0e13+0e23+0e123") != 0 {
		t.Errorf("The average is not correct!")
	}
}
